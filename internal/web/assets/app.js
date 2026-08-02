// comicread web reader — vanilla JS only, no framework.
//
// Two screens live in the same page and are toggled by hiding/showing their
// root elements: #picker (native file input) and #reader (page view). State
// (current book, page, view mode) is kept in sessionStorage only, so it
// resumes across a reload but does not outlive the browser tab's session.
(() => {
  "use strict";

  const SESSION_KEY = "comicread:session";

  const picker = document.getElementById("picker");
  const pickerStatus = document.getElementById("picker-status");
  const fileInput = document.getElementById("file-input");
  const pickButton = document.getElementById("pick-button");

  const reader = document.getElementById("reader");
  const stage = document.getElementById("stage");
  const pageImageLeft = document.getElementById("page-image-left");
  const pageImageRight = document.getElementById("page-image-right");
  const verticalPages = document.getElementById("vertical-pages");
  const pageCountButton = document.getElementById("page-count-button");
  const pageGotoForm = document.getElementById("page-goto-form");
  const pageGotoInput = document.getElementById("page-goto-input");
  const prevButton = document.getElementById("prev-button");
  const nextButton = document.getElementById("next-button");
  const viewSelect = document.getElementById("view-select");
  const pageGapLabel = document.getElementById("page-gap-label");
  const pageGapToggle = document.getElementById("page-gap-toggle");
  const fullscreenButton = document.getElementById("fullscreen-button");
  const themeButton = document.getElementById("theme-button");
  const animationSelect = document.getElementById("animation-select");
  const bgColorInput = document.getElementById("bg-color-input");
  const controls = document.getElementById("controls");
  const hideControlsButton = document.getElementById("hide-controls-button");
  const showControlsButton = document.getElementById("show-controls-button");
  const openFileButton = document.getElementById("open-file-button");
  const prefetchProgress = document.getElementById("prefetch-progress");
  const readingProgress = document.getElementById("reading-progress");

  const THEME_KEY = "comicread:theme";
  const BG_COLOR_KEY = "comicread:bg-color";
  const DEFAULT_BG_COLOR = "#000000";
  const ANIMATION_KEY = "comicread:animation";
  const ANIMATIONS = ["none", "slide", "fade", "turn", "curl"];
  const DEFAULT_ANIMATION = "fade";

  // Mirrors internal/tui/reader/book_view.go's ViewMode, with an additional
  // browser-only vertical strip that displays every physical page in order.
  const VIEW_KEY = "comicread:view";
  const PAGE_GAP_KEY = "comicread:pages-touching";
  const VIEWS = ["single-page", "vertical-scroll", "book-view", "right-view", "circle-view", "right-circle-view"];
  const DEFAULT_VIEW = "single-page";

  // state.page is a zero-based index: a physical page in single-page and
  // vertical-scroll views, or a spread index in the paired views (see
  // pageSlots below). It is only shown as a one-based physical page number.
  let state = null;

  // prefetchAbort stops the background prefetch loop (see prefetchAll)
  // started for the book previously open, once a different one replaces it.
  let prefetchAbort = null;

  // prefetchActive tracks whether prefetchAll is still running for the
  // current book, so #prefetch-progress can be hidden once every page has
  // been cached (or shown again if the controls panel is toggled back on).
  let prefetchActive = false;
  let verticalScrollFrame = 0;
  let transitionInProgress = false;
  let activeTransitionCleanup = null;

  function bookUrl(token) {
    return `/api/books/${encodeURIComponent(token)}`;
  }

  function pageUrl(token, index) {
    return `${bookUrl(token)}/pages/${index}`;
  }

  function currentViewPreference() {
    const saved = localStorage.getItem(VIEW_KEY);
    return VIEWS.includes(saved) ? saved : DEFAULT_VIEW;
  }

  function currentAnimationPreference() {
    const saved = localStorage.getItem(ANIMATION_KEY);
    return ANIMATIONS.includes(saved) ? saved : DEFAULT_ANIMATION;
  }

  function currentPageGapPreference() {
    return localStorage.getItem(PAGE_GAP_KEY) === "true";
  }

  // --- view-mode / page-pairing math, ported from
  // internal/tui/reader/book_view.go so the web reader lays out spreads
  // exactly like the terminal UI does ---

  function isBookView(view) {
    return view !== "single-page" && view !== "vertical-scroll";
  }

  function isVerticalScroll(view) {
    return view === "vertical-scroll";
  }

  function isRightToLeft(view) {
    return view === "right-view" || view === "right-circle-view";
  }

  function isCircleView(view) {
    return view === "circle-view" || view === "right-circle-view";
  }

  function circleRightPageForSpread(spread, totalPages) {
    if (spread === 0) return Math.min(1, totalPages - 1);
    return Math.min(spread * 2 + 1, totalPages - 1);
  }

  // pageSlots returns the [left, right] zero-based physical page indices to
  // display for a spread; -1 means that slot is empty. Mirrors
  // Model.pageSlots().
  function pageSlots(view, spread, totalPages) {
    if (totalPages < 1) return [-1, -1];
    if (!isBookView(view)) return [spread, -1];

    let pages;
    if (isCircleView(view)) {
      const right = circleRightPageForSpread(spread, totalPages);
      const left = spread > 0 ? circleRightPageForSpread(spread - 1, totalPages) : 0;
      pages = [left, right];
    } else {
      const left = spread * 2;
      const right = left + 1 < totalPages ? left + 1 : -1;
      pages = [left, right];
    }
    return isRightToLeft(view) ? [pages[1], pages[0]] : pages;
  }

  // canGoToNextSpread mirrors Model.canNextPage().
  function canGoToNextSpread(view, spread, totalPages) {
    if (!isBookView(view)) return spread + 1 < totalPages;
    if (!isCircleView(view)) return (spread + 1) * 2 < totalPages;
    return circleRightPageForSpread(spread, totalPages) < totalPages - 1;
  }

  function totalSpreads(view, totalPages) {
    if (totalPages < 1) return 0;
    if (!isBookView(view)) return totalPages;
    let spread = 0;
    while (canGoToNextSpread(view, spread, totalPages)) spread++;
    return spread + 1;
  }

  // lowestPageForSpread mirrors Model.currentPageNumber() (zero-based here).
  function lowestPageForSpread(view, spread, totalPages) {
    const slots = pageSlots(view, spread, totalPages);
    return Math.min(...slots.filter((slot) => slot >= 0));
  }

  // spreadForPage finds the spread that displays pageIndex, for jumping to
  // a physical page number (goto input) or preserving the reading position
  // when the view mode changes. Circle views can skip a page entirely
  // (mirroring the terminal UI), in which case the nearest preceding spread
  // is used.
  function spreadForPage(view, pageIndex, totalPages) {
    if (!isBookView(view)) return pageIndex;
    const spreads = totalSpreads(view, totalPages);
    let nearest = 0;
    for (let spread = 0; spread < spreads; spread++) {
      const slots = pageSlots(view, spread, totalPages);
      if (slots.includes(pageIndex)) return spread;
      if (lowestPageForSpread(view, spread, totalPages) <= pageIndex) nearest = spread;
    }
    return nearest;
  }

  function pageLabel(slots, totalPages) {
    const valid = slots.filter((slot) => slot >= 0);
    const low = Math.min(...valid) + 1;
    const high = Math.max(...valid) + 1;
    return low === high ? `${low} / ${totalPages}` : `${low}-${high} / ${totalPages}`;
  }

  function saveSession() {
    if (!state) return;
    sessionStorage.setItem(SESSION_KEY, JSON.stringify(state));
  }

  function loadSession() {
    const raw = sessionStorage.getItem(SESSION_KEY);
    if (!raw) return null;
    try {
      const parsed = JSON.parse(raw);
      if (parsed && typeof parsed.token === "string") return parsed;
    } catch {
      // fall through
    }
    return null;
  }

  function setStatus(message, kind) {
    pickerStatus.textContent = message || "";
    if (kind) pickerStatus.dataset.kind = kind;
    else delete pickerStatus.dataset.kind;
  }

  function showPicker() {
    cancelPageTransition();
    if (prefetchAbort) prefetchAbort.abort();
    prefetchActive = false;
    sessionStorage.removeItem(SESSION_KEY);
    state = null;
    reader.hidden = true;
    picker.hidden = false;
  }

  function showReader() {
    picker.hidden = true;
    reader.hidden = false;
    setControlsVisible(true);
  }

  function setControlsVisible(visible) {
    controls.hidden = !visible;
    showControlsButton.hidden = visible;
    updatePrefetchProgressVisibility();
  }

  // The progress strip only ever shows while prefetchAll is still running,
  // and only alongside the controls panel — hiding the panel for immersive
  // reading hides it too.
  function updatePrefetchProgressVisibility() {
    prefetchProgress.hidden = !prefetchActive || controls.hidden;
  }

  async function openFile(file) {
    const ext = (file.name.split(".").pop() || "").toLowerCase();
    if (!["cbz", "pdf", "epub"].includes(ext)) {
      reportOpenError("Unsupported file type. Choose a CBZ, PDF, or EPUB file.");
      return;
    }

    // The same file input is reused for the initial picker screen and for
    // the reader's "open a different file" button; whichever triggered this
    // open gets the busy indicator.
    const triggerButton = reader.hidden ? pickButton : openFileButton;
    setStatus("Opening…");
    triggerButton.setAttribute("aria-busy", "true");
    triggerButton.disabled = true;

    const form = new FormData();
    form.append("file", file, file.name);

    try {
      const response = await fetch("/api/open", { method: "POST", body: form });
      const body = await response.json().catch(() => ({}));
      if (!response.ok) {
        throw new Error(body.error || `server returned ${response.status}`);
      }
      startReading(body, 0, currentViewPreference());
      setStatus("");
    } catch (err) {
      reportOpenError(`Could not open file: ${err.message}`);
    } finally {
      triggerButton.removeAttribute("aria-busy");
      triggerButton.disabled = false;
    }
  }

  // reportOpenError always updates #picker-status (harmless when the picker
  // screen is hidden), and additionally alerts when the reader is what's
  // currently visible, since #picker-status can't be seen there.
  function reportOpenError(message) {
    setStatus(message, "error");
    if (!reader.hidden) alert(message);
  }

  function startReading(info, page, view) {
    cancelPageTransition();
    state = {
      token: info.token,
      title: info.title,
      totalPages: info.totalPages,
      view: VIEWS.includes(view) ? view : DEFAULT_VIEW,
      page: 0,
    };
    state.page = Math.min(Math.max(page || 0, 0), totalSpreads(state.view, state.totalPages) - 1);
    viewSelect.value = state.view;
    saveSession();
    showReader();
    renderPage();
    prefetchAll(state.token, state.totalPages);
  }

  // prefetchAll walks every physical page of the book in order, one request
  // at a time, independent of which page (or spread) is currently being
  // viewed or which view mode is active. It shares the server's per-page
  // PNG endpoint and the browser's HTTP cache with ordinary navigation, so
  // pages requested by the reader UI while this is still running are served
  // immediately and simply let this loop catch up around them later;
  // already-cached pages resolve instantly and add no extra load.
  //
  // Once every page has been cached by the browser, it tells the server to
  // release the chapter from memory (DELETE /api/books/{token}). From that
  // point on this token can no longer be resumed (e.g. after a reload) —
  // that trade-off is intentional, in exchange for not holding the chapter
  // in server memory for longer than needed.
  function prefetchAll(token, totalPages) {
    if (prefetchAbort) prefetchAbort.abort();
    const controller = new AbortController();
    prefetchAbort = controller;

    prefetchActive = true;
    prefetchProgress.max = totalPages;
    prefetchProgress.value = 0;
    updatePrefetchProgressVisibility();

    (async () => {
      for (let index = 0; index < totalPages; index++) {
        if (controller.signal.aborted) return;
        await new Promise((resolve) => {
          const img = new Image();
          img.onload = resolve;
          img.onerror = resolve;
          img.src = pageUrl(token, index);
        });
        prefetchProgress.value = index + 1;
      }
      if (controller.signal.aborted) return;
      prefetchActive = false;
      updatePrefetchProgressVisibility();
      fetch(bookUrl(token), { method: "DELETE", keepalive: true }).catch(() => {});
    })();
  }

  function renderPage() {
    if (!state) return;

    if (isVerticalScroll(state.view)) {
      updatePageGapControl([-1, -1]);
      renderVerticalPages();
      updatePageStatus([state.page, -1]);
      scrollToVerticalPage(state.page);
      return;
    }

    stage.classList.remove("vertical-scroll-view");
    verticalPages.hidden = true;
    pageImageLeft.hidden = false;
    stage.setAttribute("aria-busy", "true");

    const slots = pageSlots(state.view, state.page, state.totalPages);
    setSlotImage(pageImageLeft, slots[0]);
    setSlotImage(pageImageRight, slots[1]);
    updatePageGapControl(slots);

    updatePageStatus(slots);

    preloadSpread(state.page + 1);
    preloadSpread(state.page - 1);
  }

  function updatePageGapControl(slots) {
    const pairedView = state && isBookView(state.view);
    pageGapToggle.disabled = !pairedView;
    pageGapLabel.setAttribute("aria-disabled", String(!pairedView));
    const touching = Boolean(pairedView && slots[1] >= 0 && pageGapToggle.checked);
    stage.classList.toggle("pages-touching", touching);
    clearTouchingPageSizes();
    if (!touching) {
      return;
    }
    updateTouchingPageSizes();
  }

  function clearTouchingPageSizes() {
    for (const image of [pageImageLeft, pageImageRight]) {
      image.style.width = "";
      image.style.height = "";
    }
  }

  function updateTouchingPageSizes() {
    if (!stage.classList.contains("pages-touching") ||
        pageImageLeft.hidden || pageImageRight.hidden ||
        !pageImageLeft.complete || !pageImageRight.complete ||
        !pageImageLeft.naturalWidth || !pageImageRight.naturalWidth) return;

    const leftRatio = pageImageLeft.naturalWidth / pageImageLeft.naturalHeight;
    const rightRatio = pageImageRight.naturalWidth / pageImageRight.naturalHeight;
    const height = Math.min(stage.clientHeight, stage.clientWidth / (leftRatio + rightRatio));
    pageImageLeft.style.width = `${height * leftRatio}px`;
    pageImageLeft.style.height = `${height}px`;
    pageImageRight.style.width = `${height * rightRatio}px`;
    pageImageRight.style.height = `${height}px`;
  }

  function updatePageStatus(slots) {
    if (!state) return;

    const label = pageLabel(slots, state.totalPages);
    pageCountButton.textContent = label;
    document.title = `${label} — ${state.title} — comicread`;

    readingProgress.max = state.totalPages;
    readingProgress.value = Math.max(...slots.filter((slot) => slot >= 0)) + 1;

    prevButton.disabled = state.page <= 0;
    nextButton.disabled = !canGoToNextSpread(state.view, state.page, state.totalPages);
  }

  function renderVerticalPages() {
    stage.removeAttribute("aria-busy");
    stage.classList.add("vertical-scroll-view");
    pageImageLeft.hidden = true;
    pageImageRight.hidden = true;
    verticalPages.hidden = false;

    const currentToken = verticalPages.dataset.token;
    if (currentToken === state.token && verticalPages.children.length === state.totalPages) return;

    verticalPages.replaceChildren();
    verticalPages.dataset.token = state.token;
    const fragment = document.createDocumentFragment();
    for (let index = 0; index < state.totalPages; index++) {
      const img = document.createElement("img");
      img.src = pageUrl(state.token, index);
      img.alt = `Page ${index + 1} of ${state.totalPages}`;
      img.dataset.page = String(index);
      fragment.appendChild(img);
    }
    verticalPages.appendChild(fragment);
  }

  function scrollToVerticalPage(index) {
    const page = verticalPages.children[index];
    if (page) page.scrollIntoView({ block: "start" });
  }

  function syncPageFromVerticalScroll() {
    verticalScrollFrame = 0;
    if (!state || !isVerticalScroll(state.view) || verticalPages.hidden) return;

    const stageRect = stage.getBoundingClientRect();
    const readingLine = stageRect.top + Math.min(stageRect.height * 0.25, 160);
    let current = 0;
    for (const page of verticalPages.children) {
      if (page.getBoundingClientRect().top <= readingLine) current = Number(page.dataset.page);
      else break;
    }
    if (current === state.page) return;

    state.page = current;
    saveSession();
    updatePageStatus([current, -1]);
  }

  function setSlotImage(img, index) {
    if (index < 0) {
      img.hidden = true;
      img.removeAttribute("src");
      return;
    }
    img.hidden = false;
    img.src = pageUrl(state.token, index);
    img.alt = `Page ${index + 1} of ${state.totalPages}`;
  }

  function preloadSpread(spread) {
    if (!state || spread < 0 || spread >= totalSpreads(state.view, state.totalPages)) return;
    const slots = pageSlots(state.view, spread, state.totalPages);
    preload(slots[0]);
    preload(slots[1]);
  }

  function preload(index) {
    if (!state || index < 0 || index >= state.totalPages) return;
    const img = new Image();
    img.src = pageUrl(state.token, index);
  }

  function goTo(spread) {
    if (!state || transitionInProgress) return;
    const clamped = Math.min(Math.max(spread, 0), totalSpreads(state.view, state.totalPages) - 1);
    if (clamped === state.page) return;
    const previous = state.page;
    const transition = capturePageTransition();
    state.page = clamped;
    saveSession();
    renderPage();
    playPageTransition(transition, clamped > previous);
  }

  // Capture the currently displayed page or spread before renderPage swaps
  // its image sources. The copy sits above the new spread and animates away,
  // avoiding duplicate page-layout logic for single and paired views.
  function capturePageTransition() {
    const animation = currentAnimationPreference();
    if (animation === "none" || isVerticalScroll(state.view) ||
        window.matchMedia("(prefers-reduced-motion: reduce)").matches) return null;

    const oldLayer = capturePageLayer("old");
    if (!oldLayer) return null;
    const backdrop = document.createElement("div");
    backdrop.className = "page-transition-backdrop";
    backdrop.style.backgroundColor = getComputedStyle(stage).backgroundColor;
    stage.appendChild(backdrop);
    return { oldLayer, backdrop, animation };
  }

  function capturePageLayer(kind) {
    const layer = document.createElement("div");
    layer.className = `page-transition-layer page-transition-${kind}`;
    layer.style.backgroundColor = getComputedStyle(stage).backgroundColor;
    const stageRect = stage.getBoundingClientRect();
    for (const image of [pageImageLeft, pageImageRight]) {
      if (image.hidden || !image.getAttribute("src")) continue;
      const copy = image.cloneNode(false);
      copy.removeAttribute("id");
      copy.hidden = false;
      const rect = image.getBoundingClientRect();
      // Transition layers must preserve the rendered page geometry exactly.
      // Re-running flex layout here would reintroduce a gap for No gap mode.
      copy.style.position = "absolute";
      copy.style.left = `${rect.left - stageRect.left}px`;
      copy.style.top = `${rect.top - stageRect.top}px`;
      copy.style.width = `${rect.width}px`;
      copy.style.height = `${rect.height}px`;
      layer.appendChild(copy);
    }
    if (!layer.children.length) return null;
    stage.appendChild(layer);
    return layer;
  }

  function playPageTransition(transition, forward) {
    if (!transition) return;

    // Forward navigation moves toward the reading direction. Manga layouts
    // therefore mirror every directional transition automatically.
    const movesRight = forward === isRightToLeft(state.view);
    transitionInProgress = true;

    let animations = [];
    let finished = false;
    const finish = () => {
      if (finished) return;
      finished = true;
      for (const animation of animations) animation.cancel();
      transition.oldLayer.remove();
      if (transition.newLayer) transition.newLayer.remove();
      if (transition.shadow) transition.shadow.remove();
      if (transition.leaf) transition.leaf.remove();
      transition.backdrop.remove();
      transitionInProgress = false;
      if (activeTransitionCleanup === finish) activeTransitionCleanup = null;
    };
    activeTransitionCleanup = finish;

    waitForDisplayedPages().then(() => {
      if (finished) return;
      transition.newLayer = capturePageLayer("new");
      if (!transition.newLayer || typeof transition.oldLayer.animate !== "function") {
        finish();
        return;
      }
      animations = animatePageTransition(transition, movesRight);
      Promise.allSettled(animations.map((animation) => animation.finished)).then(finish);
    });
  }

  // Animate old and new snapshots together, matching the model used by the
  // browser View Transition API. This keeps motion short and preserves visual
  // continuity instead of throwing the old spread completely off screen.
  function animatePageTransition(transition, movesRight) {
    const oldLayer = transition.oldLayer;
    const newLayer = transition.newLayer;
    const direction = movesRight ? 1 : -1;

    switch (transition.animation) {
      case "fade":
        return [
          oldLayer.animate([{ opacity: 1 }, { opacity: 0 }], motionOptions(150)),
          newLayer.animate([{ opacity: 0 }, { opacity: 1 }], motionOptions(150)),
        ];
      case "turn": {
        if (isBookView(state.view) && oldLayer.children.length > 1 && newLayer.children.length > 1) {
          return animateBookTurn(transition, movesRight, false);
        }
        const angle = direction * 78;
        const origin = movesRight ? "right center" : "left center";
        oldLayer.style.transformOrigin = origin;
        newLayer.style.transformOrigin = origin;
        return [
          oldLayer.animate([
            { opacity: 1, transform: "perspective(1800px) rotateY(0deg)", filter: "brightness(1)" },
            { opacity: 0, transform: `perspective(1800px) rotateY(${angle}deg)`, filter: "brightness(.72)" },
          ], motionOptions(300)),
          newLayer.animate([
            { opacity: 0, transform: `perspective(1800px) rotateY(${-angle}deg)`, filter: "brightness(.72)" },
            { opacity: 1, transform: "perspective(1800px) rotateY(0deg)", filter: "brightness(1)" },
          ], motionOptions(300)),
        ];
      }
      case "curl":
        if (isBookView(state.view) && oldLayer.children.length > 1 && newLayer.children.length > 1) {
          return animateBookTurn(transition, movesRight, true);
        }
        return animatePageCurl(transition, movesRight);
      case "slide":
      default:
        return [
          oldLayer.animate([
            { transform: "translate3d(0, 0, 0)" },
            { transform: `translate3d(${direction * 100}%, 0, 0)` },
          ], motionOptions(220)),
          newLayer.animate([
            { transform: `translate3d(${-direction * 100}%, 0, 0)` },
            { transform: "translate3d(0, 0, 0)" },
          ], motionOptions(220)),
        ];
    }
  }

  function motionOptions(duration) {
    return { duration, easing: "cubic-bezier(.4, 0, .2, 1)", fill: "both" };
  }

  // In paired layouts, a real book turns one leaf around the centre spine.
  // Its front is the outgoing outer page and its back is the incoming page
  // on the opposite side. The remaining half swaps near the edge-on point.
  function animateBookTurn(transition, movesRight, soft) {
    const oldImages = [...transition.oldLayer.children];
    const newImages = [...transition.newLayer.children];
    const frontSource = movesRight ? oldImages[0] : oldImages[oldImages.length - 1];
    const backSource = movesRight ? newImages[newImages.length - 1] : newImages[0];
    const stageRect = stage.getBoundingClientRect();
    const pageRect = frontSource.getBoundingClientRect();
    const direction = movesRight ? 1 : -1;
    const duration = soft ? 420 : 350;

    const leaf = document.createElement("div");
    leaf.className = "page-turn-leaf";
    leaf.style.left = `${pageRect.left - stageRect.left}px`;
    leaf.style.top = `${pageRect.top - stageRect.top}px`;
    leaf.style.width = `${pageRect.width}px`;
    leaf.style.height = `${pageRect.height}px`;
    leaf.style.transformOrigin = movesRight ? "right center" : "left center";
    leaf.style.backgroundColor = getComputedStyle(stage).backgroundColor;

    const front = frontSource.cloneNode(false);
    front.className = "page-turn-face page-turn-front";
    const back = backSource.cloneNode(false);
    back.className = "page-turn-face page-turn-back";
    for (const face of [front, back]) {
      face.style.left = "";
      face.style.top = "";
      face.style.width = "100%";
      face.style.height = "100%";
    }
    const highlight = document.createElement("div");
    highlight.className = "page-turn-highlight";
    highlight.style.background = movesRight
      ? "linear-gradient(to left, rgb(255 255 255 / .24), transparent 32%, rgb(0 0 0 / .28))"
      : "linear-gradient(to right, rgb(255 255 255 / .24), transparent 32%, rgb(0 0 0 / .28))";
    leaf.append(front, back, highlight);
    stage.appendChild(leaf);
    transition.leaf = leaf;
    transition.oldLayer.style.backgroundColor = "transparent";
    frontSource.style.visibility = "hidden";

    const angle = direction * 180;
    const leafFrames = soft
      ? [
          { transform: "rotateY(0deg) skewY(0deg)", filter: "drop-shadow(0 0 0 rgb(0 0 0 / 0))" },
          { transform: `rotateY(${angle * .48}deg) skewY(${direction * 3.5}deg) scaleX(.94)`, filter: `drop-shadow(${-direction * 1.1}rem .15rem .7rem rgb(0 0 0 / .42))`, offset: .48 },
          { transform: `rotateY(${angle}deg) skewY(0deg)`, filter: "drop-shadow(0 0 0 rgb(0 0 0 / 0))" },
        ]
      : [
          { opacity: 1, transform: "rotateY(0deg)", filter: "brightness(1) drop-shadow(0 0 0 rgb(0 0 0 / 0))" },
          { opacity: .5, transform: `rotateY(${angle * .5}deg)`, filter: `brightness(.76) drop-shadow(${-direction * .8}rem 0 .55rem rgb(0 0 0 / .38))`, offset: .5 },
          { opacity: 0, transform: `rotateY(${angle}deg)`, filter: "brightness(1) drop-shadow(0 0 0 rgb(0 0 0 / 0))" },
        ];

    const animations = [
      leaf.animate(leafFrames, { ...motionOptions(duration), easing: soft ? "cubic-bezier(.37, .02, .2, 1)" : "cubic-bezier(.45, .05, .2, 1)" }),
      transition.oldLayer.animate([
        { opacity: 1 },
        ...(soft ? [{ opacity: 1, offset: .46 }, { opacity: 0, offset: .54 }] : []),
        { opacity: 0 },
      ], motionOptions(duration)),
      highlight.animate([
        { opacity: 0 },
        { opacity: soft ? .9 : .55, offset: .5 },
        { opacity: 0 },
      ], motionOptions(duration)),
    ];
    if (!soft) {
      animations.push(transition.newLayer.animate([{ opacity: 0 }, { opacity: 1 }], motionOptions(duration)));
    }
    return animations;
  }

  // A soft page curl is a moving fold rather than a rigid 3D rotation. The
  // visible page is clipped by a changing polygon while a narrow two-sided
  // gradient follows the fold to provide highlight and cast shadow.
  function animatePageCurl(transition, movesRight) {
    const frames = [];
    const shadowFrames = [];
    const steps = 24;
    const stageWidth = stage.clientWidth;
    const shadowWidth = 64;

    transition.shadow = document.createElement("div");
    transition.shadow.className = "page-curl-shadow";
    transition.shadow.style.background = movesRight
      ? "linear-gradient(to left, transparent, rgb(255 255 255 / .32) 42%, rgb(0 0 0 / .38) 58%, transparent)"
      : "linear-gradient(to right, transparent, rgb(255 255 255 / .32) 42%, rgb(0 0 0 / .38) 58%, transparent)";
    stage.appendChild(transition.shadow);

    for (let index = 0; index <= steps; index++) {
      const progress = index / steps;
      const edge = movesRight ? progress * 100 : (1 - progress) * 100;
      const bow = Math.sin(Math.PI * progress) * 6;
      const top = Math.min(100, Math.max(0, edge + (movesRight ? bow : -bow)));
      const bottom = Math.min(100, Math.max(0, edge + (movesRight ? -bow : bow)));
      const clipPath = movesRight
        ? `polygon(${top}% 0, 100% 0, 100% 100%, ${bottom}% 100%)`
        : `polygon(0 0, ${top}% 0, ${bottom}% 100%, 0 100%)`;
      const shadowX = edge / 100 * stageWidth - shadowWidth / 2;
      frames.push({ clipPath, offset: progress });
      shadowFrames.push({
        opacity: Math.sin(Math.PI * progress) * .9,
        transform: `translate3d(${shadowX}px, 0, 0) skewY(${movesRight ? -bow : bow}deg)`,
        offset: progress,
      });
    }

    return [
      transition.oldLayer.animate(frames, motionOptions(360)),
      transition.shadow.animate(shadowFrames, motionOptions(360)),
    ];
  }

  function waitForDisplayedPages() {
    const pending = [pageImageLeft, pageImageRight]
      .filter((image) => !image.hidden && image.getAttribute("src") && !image.complete)
      .map((image) => new Promise((resolve) => {
        const done = () => {
          image.removeEventListener("load", done);
          image.removeEventListener("error", done);
          resolve();
        };
        image.addEventListener("load", done);
        image.addEventListener("error", done);
      }));
    if (!pending.length) return Promise.resolve();
    return Promise.race([
      Promise.all(pending),
      new Promise((resolve) => setTimeout(resolve, 3000)),
    ]);
  }

  function cancelPageTransition() {
    if (activeTransitionCleanup) activeTransitionCleanup();
  }

  function setView(view) {
    if (!state || !VIEWS.includes(view) || view === state.view) return;
    cancelPageTransition();
    localStorage.setItem(VIEW_KEY, view);
    const anchorPage = lowestPageForSpread(state.view, state.page, state.totalPages);
    state.view = view;
    state.page = spreadForPage(view, anchorPage, state.totalPages);
    saveSession();
    renderPage();
  }

  function toggleFullscreen() {
    if (document.fullscreenElement) {
      document.exitFullscreen();
    } else {
      reader.requestFullscreen().catch(() => {
        // Fullscreen may be unavailable (e.g. iOS Safari); nothing to do.
      });
    }
  }

  function openGoto() {
    if (!state) return;
    pageGotoInput.max = String(state.totalPages);
    pageGotoInput.value = String(lowestPageForSpread(state.view, state.page, state.totalPages) + 1);
    pageCountButton.hidden = true;
    pageGotoForm.hidden = false;
    pageGotoInput.focus();
    pageGotoInput.select();
  }

  function closeGoto() {
    pageGotoForm.hidden = true;
    pageCountButton.hidden = false;
  }

  function applyTheme(theme) {
    document.documentElement.setAttribute("data-theme", theme);
    document.documentElement.style.colorScheme = theme;
    themeButton.textContent = theme === "dark" ? "☀️" : "🌙";
    themeButton.setAttribute("aria-label", theme === "dark" ? "Switch to light theme" : "Switch to dark theme");
  }

  function applyBgColor(color) {
    stage.style.backgroundColor = color;
    bgColorInput.value = color;
  }

  // --- wiring ---

  pickButton.addEventListener("click", () => fileInput.click());
  openFileButton.addEventListener("click", () => fileInput.click());
  fileInput.addEventListener("change", () => {
    const file = fileInput.files && fileInput.files[0];
    fileInput.value = "";
    if (file) openFile(file);
  });

  prevButton.addEventListener("click", () => goTo(state.page - 1));
  nextButton.addEventListener("click", () => goTo(state.page + 1));
  viewSelect.addEventListener("change", () => setView(viewSelect.value));
  pageGapToggle.addEventListener("change", () => {
    localStorage.setItem(PAGE_GAP_KEY, String(pageGapToggle.checked));
    if (state) updatePageGapControl(pageSlots(state.view, state.page, state.totalPages));
  });
  animationSelect.addEventListener("change", () => {
    localStorage.setItem(ANIMATION_KEY, animationSelect.value);
  });
  fullscreenButton.addEventListener("click", toggleFullscreen);
  hideControlsButton.addEventListener("click", () => setControlsVisible(false));
  showControlsButton.addEventListener("click", () => setControlsVisible(true));
  stage.addEventListener("scroll", () => {
    if (!isVerticalScroll(state && state.view) || verticalScrollFrame) return;
    verticalScrollFrame = requestAnimationFrame(syncPageFromVerticalScroll);
  }, { passive: true });

  document.addEventListener("fullscreenchange", () => {
    fullscreenButton.setAttribute("aria-pressed", String(Boolean(document.fullscreenElement)));
  });

  pageCountButton.addEventListener("click", openGoto);
  pageGotoForm.addEventListener("submit", (event) => {
    event.preventDefault();
    const value = parseInt(pageGotoInput.value, 10);
    if (!Number.isNaN(value)) {
      const targetPage = Math.min(Math.max(value - 1, 0), state.totalPages - 1);
      goTo(spreadForPage(state.view, targetPage, state.totalPages));
    }
    closeGoto();
  });
  pageGotoInput.addEventListener("blur", closeGoto);
  pageGotoInput.addEventListener("keydown", (event) => {
    if (event.key === "Escape") {
      event.preventDefault();
      closeGoto();
    }
    // Keep digit entry and cursor movement from being intercepted by the
    // reader's own keyboard shortcuts below (e.g. Home/End, arrow keys).
    event.stopPropagation();
  });

  themeButton.addEventListener("click", () => {
    const next = document.documentElement.getAttribute("data-theme") === "dark" ? "light" : "dark";
    localStorage.setItem(THEME_KEY, next);
    applyTheme(next);
  });

  bgColorInput.addEventListener("input", () => {
    localStorage.setItem(BG_COLOR_KEY, bgColorInput.value);
    applyBgColor(bgColorInput.value);
  });

  for (const img of [pageImageLeft, pageImageRight]) {
    img.addEventListener("load", () => {
      stage.removeAttribute("aria-busy");
      updateTouchingPageSizes();
    });
    img.addEventListener("error", () => stage.removeAttribute("aria-busy"));
  }

  window.addEventListener("resize", updateTouchingPageSizes);

  document.addEventListener("keydown", (event) => {
    if (reader.hidden || !pageGotoForm.hidden) return;
    if (event.altKey || event.ctrlKey || event.metaKey) return;

    switch (event.key) {
      case "ArrowLeft":
      case "a":
        goTo(state.page - 1);
        break;
      case "ArrowRight":
      case "d":
      case " ":
        goTo(state.page + 1);
        break;
      case "Home":
        goTo(0);
        break;
      case "End":
        goTo(totalSpreads(state.view, state.totalPages) - 1);
        break;
      case "f":
        toggleFullscreen();
        break;
      default:
        return;
    }
    event.preventDefault();
  });

  // --- startup: apply the saved theme, page background and view mode, then
  // resume an in-session book or show the picker ---

  applyTheme(localStorage.getItem(THEME_KEY) === "light" ? "light" : "dark");
  applyBgColor(localStorage.getItem(BG_COLOR_KEY) || DEFAULT_BG_COLOR);
  viewSelect.value = currentViewPreference();
  pageGapToggle.checked = currentPageGapPreference();
  animationSelect.value = currentAnimationPreference();

  (async function init() {
    const saved = loadSession();
    if (!saved) {
      showPicker();
      return;
    }
    try {
      const response = await fetch(bookUrl(saved.token));
      if (!response.ok) throw new Error("session expired");
      const info = await response.json();
      startReading(info, saved.page, saved.view);
    } catch {
      showPicker();
    }
  })();
})();
