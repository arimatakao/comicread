// comicread web reader — vanilla JS only, no framework.
//
// Two screens live in the same page and are toggled by hiding/showing their
// root elements: #picker (native file input) and #reader (page view). State
// (current book, page) is kept in sessionStorage only, so it resumes across
// a reload but does not outlive the browser tab's session.
(() => {
  "use strict";

  const SESSION_KEY = "comicread:session";

  const picker = document.getElementById("picker");
  const pickerStatus = document.getElementById("picker-status");
  const fileInput = document.getElementById("file-input");
  const pickButton = document.getElementById("pick-button");

  const reader = document.getElementById("reader");
  const stage = document.getElementById("stage");
  const pageImage = document.getElementById("page-image");
  const pageCountButton = document.getElementById("page-count-button");
  const pageGotoForm = document.getElementById("page-goto-form");
  const pageGotoInput = document.getElementById("page-goto-input");
  const prevButton = document.getElementById("prev-button");
  const nextButton = document.getElementById("next-button");
  const fullscreenButton = document.getElementById("fullscreen-button");
  const themeButton = document.getElementById("theme-button");

  const THEME_KEY = "comicread:theme";

  // state.page is a zero-based page index, matching the server's
  // /api/books/{token}/pages/{index} route. It is only ever shown to the
  // user as a one-based number.
  let state = null;

  // prefetchAbort stops the background prefetch loop (see prefetchAll)
  // started for the book previously open, once a different one replaces it.
  let prefetchAbort = null;

  function bookUrl(token) {
    return `/api/books/${encodeURIComponent(token)}`;
  }

  function pageUrl(token, index) {
    return `${bookUrl(token)}/pages/${index}`;
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
    sessionStorage.removeItem(SESSION_KEY);
    state = null;
    reader.hidden = true;
    picker.hidden = false;
  }

  function showReader() {
    picker.hidden = true;
    reader.hidden = false;
  }

  async function openFile(file) {
    const ext = (file.name.split(".").pop() || "").toLowerCase();
    if (!["cbz", "pdf", "epub"].includes(ext)) {
      setStatus("Unsupported file type. Choose a CBZ, PDF, or EPUB file.", "error");
      return;
    }

    setStatus("Opening…");
    pickButton.setAttribute("aria-busy", "true");
    pickButton.disabled = true;

    const form = new FormData();
    form.append("file", file, file.name);

    try {
      const response = await fetch("/api/open", { method: "POST", body: form });
      const body = await response.json().catch(() => ({}));
      if (!response.ok) {
        throw new Error(body.error || `server returned ${response.status}`);
      }
      startReading(body, 0);
      setStatus("");
    } catch (err) {
      setStatus(`Could not open file: ${err.message}`, "error");
    } finally {
      pickButton.removeAttribute("aria-busy");
      pickButton.disabled = false;
    }
  }

  function startReading(info, page) {
    state = {
      token: info.token,
      title: info.title,
      totalPages: info.totalPages,
      page: Math.min(Math.max(page || 0, 0), info.totalPages - 1),
    };
    saveSession();
    showReader();
    renderPage();
    prefetchAll(state.token, state.totalPages);
  }

  // prefetchAll walks every page of the book in order, one request at a
  // time, independent of which page is currently being viewed. It shares
  // the server's per-page PNG endpoint and the browser's HTTP cache with
  // ordinary navigation, so pages requested by the reader UI while this is
  // still running are served immediately and simply let this loop catch up
  // around them later; already-cached pages resolve instantly and add no
  // extra load.
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

    (async () => {
      for (let index = 0; index < totalPages; index++) {
        if (controller.signal.aborted) return;
        await new Promise((resolve) => {
          const img = new Image();
          img.onload = resolve;
          img.onerror = resolve;
          img.src = pageUrl(token, index);
        });
      }
      if (controller.signal.aborted) return;
      fetch(bookUrl(token), { method: "DELETE", keepalive: true }).catch(() => {});
    })();
  }

  function renderPage() {
    if (!state) return;
    stage.setAttribute("aria-busy", "true");
    pageImage.src = pageUrl(state.token, state.page);
    pageImage.alt = `Page ${state.page + 1} of ${state.totalPages}`;
    pageCountButton.textContent = `${state.page + 1} / ${state.totalPages}`;

    prevButton.disabled = state.page <= 0;
    nextButton.disabled = state.page >= state.totalPages - 1;

    document.title = `${state.page + 1}/${state.totalPages} — ${state.title} — comicread`;

    preload(state.page + 1);
    preload(state.page - 1);
  }

  function preload(index) {
    if (!state || index < 0 || index >= state.totalPages) return;
    const img = new Image();
    img.src = pageUrl(state.token, index);
  }

  function goTo(index) {
    if (!state) return;
    const clamped = Math.min(Math.max(index, 0), state.totalPages - 1);
    if (clamped === state.page) return;
    state.page = clamped;
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
    pageGotoInput.value = String(state.page + 1);
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

  // --- wiring ---

  pickButton.addEventListener("click", () => fileInput.click());
  fileInput.addEventListener("change", () => {
    const file = fileInput.files && fileInput.files[0];
    fileInput.value = "";
    if (file) openFile(file);
  });

  prevButton.addEventListener("click", () => goTo(state.page - 1));
  nextButton.addEventListener("click", () => goTo(state.page + 1));
  fullscreenButton.addEventListener("click", toggleFullscreen);

  document.addEventListener("fullscreenchange", () => {
    fullscreenButton.setAttribute("aria-pressed", String(Boolean(document.fullscreenElement)));
  });

  pageCountButton.addEventListener("click", openGoto);
  pageGotoForm.addEventListener("submit", (event) => {
    event.preventDefault();
    const value = parseInt(pageGotoInput.value, 10);
    if (!Number.isNaN(value)) goTo(value - 1);
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

  pageImage.addEventListener("load", () => stage.removeAttribute("aria-busy"));
  pageImage.addEventListener("error", () => stage.removeAttribute("aria-busy"));

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
        goTo(state.totalPages - 1);
        break;
      case "f":
        toggleFullscreen();
        break;
      default:
        return;
    }
    event.preventDefault();
  });

  // --- startup: apply the saved theme, then resume an in-session book or
  // show the picker ---

  applyTheme(localStorage.getItem(THEME_KEY) === "light" ? "light" : "dark");

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
      startReading(info, saved.page);
    } catch {
      showPicker();
    }
  })();
})();
