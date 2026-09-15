/* AlphaFlux site behaviour. No dependencies, deferred.
   The page is readable and navigable with this file absent. */

(function () {
  "use strict";

  var root = document.documentElement;

  /* --- Theme. The inline head script has already applied the stored choice,
     so this only handles the toggle. --------------------------------------- */
  var toggle = document.querySelector("[data-theme-toggle]");
  if (toggle) {
    toggle.hidden = false;
    toggle.addEventListener("click", function () {
      var next = root.getAttribute("data-theme") === "dark" ? "light" : "dark";
      root.setAttribute("data-theme", next);
      try { localStorage.setItem("af-theme", next); } catch (e) {}
      var meta = document.querySelector('meta[name="theme-color"]:not([media])');
      if (meta) meta.setAttribute("content", next === "dark" ? "#0a1220" : "#ffffff");
    });
  }

  /* --- Mobile drawer ----------------------------------------------------- */
  var drawer = document.querySelector("[data-drawer]");
  var openBtn = document.querySelector("[data-drawer-open]");
  var closeBtn = document.querySelector("[data-drawer-close]");

  function setDrawer(open) {
    if (!drawer) return;
    drawer.setAttribute("data-open", open ? "true" : "false");
    drawer.setAttribute("aria-hidden", open ? "false" : "true");
    if (openBtn) openBtn.setAttribute("aria-expanded", open ? "true" : "false");
    document.body.style.overflow = open ? "hidden" : "";
    if (open && closeBtn) closeBtn.focus();
    else if (!open && openBtn) openBtn.focus();
  }

  if (openBtn) openBtn.addEventListener("click", function () { setDrawer(true); });
  if (closeBtn) closeBtn.addEventListener("click", function () { setDrawer(false); });
  if (drawer) {
    drawer.addEventListener("click", function (e) {
      if (e.target.hasAttribute("data-drawer-scrim")) setDrawer(false);
    });
  }
  document.addEventListener("keydown", function (e) {
    if (e.key === "Escape" && drawer && drawer.getAttribute("data-open") === "true") setDrawer(false);
  });

  var reduced = window.matchMedia("(prefers-reduced-motion: reduce)").matches;

  /* --- Copy buttons on code blocks --------------------------------------- */
  document.querySelectorAll("[data-copy]").forEach(function (btn) {
    btn.addEventListener("click", function () {
      var el = document.getElementById(btn.getAttribute("data-copy"));
      if (!el) return;
      var text = el.innerText;
      var done = function () {
        var prev = btn.textContent;
        btn.textContent = "Copied";
        setTimeout(function () { btn.textContent = prev; }, 1600);
      };
      if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(text).then(done, function () {});
      } else {
        var ta = document.createElement("textarea");
        ta.value = text;
        document.body.appendChild(ta);
        ta.select();
        try { document.execCommand("copy"); done(); } catch (e) {}
        document.body.removeChild(ta);
      }
    });
  });

  /* --- The work-order strip advances one stage while it is on screen, so a
     visitor sees what the record does rather than reading that it moves. --- */
  var strip = document.querySelector("[data-order-strip]");
  if (strip && !reduced) {
    var stages = strip.querySelectorAll("[data-state]");
    var i = 0;
    var advance = function () {
      stages.forEach(function (el, n) {
        el.setAttribute("data-state", n < i ? "done" : n === i ? "now" : "todo");
      });
      i = (i + 1) % (stages.length + 2);
    };
    var timer = null;
    var start = function () { if (!timer) timer = setInterval(advance, 1400); };
    var stop = function () { clearInterval(timer); timer = null; };
    if ("IntersectionObserver" in window) {
      new IntersectionObserver(function (entries) {
        entries[0].isIntersecting ? start() : stop();
      }, { threshold: 0.25 }).observe(strip);
    }
  }
})();
