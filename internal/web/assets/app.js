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
  const fullscreenButton = document.getElementById("fullscreen-button");
  const themeButton = document.getElementById("theme-button");
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

  // Mirrors internal/tui/reader/book_view.go's ViewMode, with an additional
  // browser-only vertical strip that displays every physical page in order.
  const VIEW_KEY = "comicread:view";
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

    updatePageStatus(slots);

    preloadSpread(state.page + 1);
    preloadSpread(state.page - 1);
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
    if (!state) return;
    const clamped = Math.min(Math.max(spread, 0), totalSpreads(state.view, state.totalPages) - 1);
    if (clamped === state.page) return;
    state.page = clamped;
    saveSession();
    renderPage();
  }

  function setView(view) {
    if (!state || !VIEWS.includes(view) || view === state.view) return;
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
    img.addEventListener("load", () => stage.removeAttribute("aria-busy"));
    img.addEventListener("error", () => stage.removeAttribute("aria-busy"));
  }

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
