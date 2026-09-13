/* ============================================================
   ALPHAFLUX // full-page scroll story (Enter site)
   One fixed scene. A tall track. Scrolling morphs the scene and
   feeds in chapters of the site, gist style. Hand-rolled, no GSAP.
   ============================================================ */
(function () {
  "use strict";
  var NS = "http://www.w3.org/2000/svg";
  var body = document.body;

  function lerp(a, b, t) { return a + (b - a) * t; }
  function hex(h) { return [parseInt(h.slice(1, 3), 16), parseInt(h.slice(3, 5), 16), parseInt(h.slice(5, 7), 16)]; }
  function css(rgb) { return "rgb(" + (rgb[0] | 0) + "," + (rgb[1] | 0) + "," + (rgb[2] | 0) + ")"; }
  function lerpHex(a, b, t) {
    var A = hex(a), B = hex(b);
    return css([lerp(A[0], B[0], t), lerp(A[1], B[1], t), lerp(A[2], B[2], t)]);
  }

  var SKIES = ["#05070f", "#041a24", "#140a2c", "#280a22", "#05070f"];
  var R1 = ["#2b2150", "#12303a", "#241247", "#3a1130", "#2b2150"];
  var R2 = ["#1c1740", "#0c2430", "#190f38", "#2c0e26", "#1c1740"];
  var R3 = ["#0d0f24", "#081a22", "#120b28", "#200a1e", "#0d0f24"];

  /* scene element refs */
  var refs = {};

  function el(name, attrs, parent) {
    var e = document.createElementNS(NS, name);
    for (var k in attrs) { if (attrs.hasOwnProperty(k)) e.setAttribute(k, attrs[k]); }
    if (parent) parent.appendChild(e);
    return e;
  }

  function buildScene() {
    var host = document.getElementById("storyScene");
    if (!host || host.childNodes.length) return;
    var svg = el("svg", { viewBox: "0 0 1440 900", preserveAspectRatio: "xMidYMax slice" }, host);
    svg.style.cssText = "position:absolute;inset:0;width:100%;height:100%";

    var defs = el("defs", {}, svg);
    defs.innerHTML =
      '<radialGradient id="sglow" cx="50%" cy="50%" r="50%"><stop offset="0" stop-color="#35d6ff" stop-opacity="0.45"/><stop offset="1" stop-color="#35d6ff" stop-opacity="0"/></radialGradient>' +
      '<radialGradient id="splanet"><stop offset="0" stop-color="#bff4ff"/><stop offset="0.5" stop-color="#28e3ff"/><stop offset="1" stop-color="#8b6bff"/></radialGradient>' +
      '<linearGradient id="swave" x1="0" y1="0" x2="0" y2="1"><stop offset="0" stop-color="#35d6ff" stop-opacity="0.35"/><stop offset="1" stop-color="#8b6bff" stop-opacity="0"/></linearGradient>';

    refs.sky = el("rect", { width: "1440", height: "900", fill: SKIES[0] }, svg);

    /* stars */
    var stars = el("g", {}, svg);
    for (var i = 0; i < 170; i++) {
      el("circle", { cx: (Math.random() * 1440).toFixed(0), cy: (Math.random() * 900).toFixed(0), r: (Math.random() * 1.5 + 0.3).toFixed(2), fill: "#9fc4ff", opacity: (Math.random() * 0.7 + 0.2).toFixed(2) }, stars);
    }
    refs.stars = stars;

    /* ringed planet group */
    var planet = el("g", {}, svg);
    el("circle", { cx: "1150", cy: "240", r: "110", fill: "url(#splanet)" }, planet);
    el("ellipse", { cx: "1150", cy: "240", rx: "190", ry: "36", fill: "none", stroke: "#35d6ff", "stroke-width": "5", opacity: "0.5", transform: "rotate(-14 1150 240)" }, planet);
    el("circle", { cx: "300", cy: "150", r: "300", fill: "url(#sglow)" }, planet);
    refs.planet = planet;

    /* three ridges */
    function ridge(attrs) { return el("path", attrs, svg); }
    refs.r1 = ridge({ fill: R1[0] });
    refs.r2 = ridge({ fill: R2[0] });
    refs.r3 = ridge({ fill: R3[0] });

    /* foreground signal wave + core */
    refs.wave = el("path", { fill: "url(#swave)" }, svg);
    var coreG = el("g", {}, svg);
    refs.core = el("circle", { cx: "720", cy: "600", r: "22", fill: "#ffffff", opacity: "0.9" }, coreG);
    el("circle", { cx: "720", cy: "600", r: "70", fill: "url(#sglow)" }, coreG);
    refs.coreG = coreG;

    refs.host = host;
  }

  function ridgePath(cx, amp, peaks, seed) {
    var pts = [];
    for (var i = 0; i <= peaks; i++) {
      var x = (i / peaks) * 1700 - 130;
      var w = Math.abs(Math.sin(i * 1.7 + seed)) * 0.6 + Math.abs(Math.sin(i * 0.9 + seed * 2.3)) * 0.4;
      pts.push(x.toFixed(0) + "," + (cx - w * amp).toFixed(0));
    }
    return "M" + pts.join(" L") + " L1700 520 L-130 520 Z";
  }

  /* per-chapter slide content (the "parts of the website") */
  var SLIDES = [
    { top: "ALPHAFLUX", title: "Find the signal. Convert the demand.", sub: "An AI-native marketing and conversion engine. Deployed on your hardware. Your data stays yours." },
    { top: "WHAT IT IS", title: "AI that works while you sleep", sub: "Watches intent across the web, scores every lead with your own model, and closes the loop with automation." },
    { top: "HOW IT CONVERTS", title: "Signal. Score. Convert.", sub: "Locate demand, rank readiness, close the loop. Every conversion feeds back into the model." },
    { top: "CAPABILITIES", title: "Six engines, one stack", sub: "AI campaigns, conversion intelligence, lead GPS, automation, own-your-stack, privacy first." },
    { top: "ENTER", title: "Ready to explore?", sub: "Fly the galaxy or dive into the capabilities." }
  ];

  function buildSlides() {
    var wrap = document.getElementById("storySlides");
    if (!wrap) return;
    for (var i = 0; i < SLIDES.length; i++) {
      var s = SLIDES[i];
      var d = document.createElement("div");
      d.className = "slide";
      d.innerHTML = '<div class="slide-in"><div class="sl-top">' + s.top + '</div>' +
        '<h2>' + s.title + '</h2><p>' + s.sub + '</p></div>';
      wrap.appendChild(d);
      if (i === 0) {
        var go = document.createElement("button");
        go.className = "story-cta";
        go.textContent = "Begin";
        go.addEventListener("click", function () {
          var track = document.getElementById("storyTrack");
          if (track) window.scrollTo({ top: track.offsetTop + window.innerHeight * 1.1, behavior: "smooth" });
        });
        d.querySelector(".slide-in").appendChild(go);
      }
      if (i === SLIDES.length - 1) {
        var row = document.createElement("div");
        row.className = "story-btns";
        var g = document.createElement("button");
        g.className = "story-cta primary";
        g.textContent = "Enter the Galaxy";
        g.addEventListener("click", function () { if (window.__enterGalaxy) window.__enterGalaxy(); });
        var s2 = document.createElement("button");
        s2.className = "story-cta";
        s2.textContent = "Back to top";
        s2.addEventListener("click", function () { window.scrollTo({ top: 0, behavior: "smooth" }); });
        row.appendChild(g); row.appendChild(s2);
        d.querySelector(".slide-in").appendChild(row);
      }
    }
  }

  /* progress -> chapter index + local t */
  function chapter(p, n) {
    var c = Math.min(n - 1, Math.floor(p * n));
    var t = p * n - c;
    return { c: c, t: t };
  }

  function render(p) {
    var ch = chapter(p, SLIDES.length);
    var sky = SKIES[ch.c], skyNext = SKIES[Math.min(ch.c + 1, SKIES.length - 1)];
    refs.sky.setAttribute("fill", lerpHex(sky, skyNext, ch.t));
    refs.r1.setAttribute("fill", lerpHex(R1[ch.c], R1[Math.min(ch.c + 1, R1.length - 1)], ch.t));
    refs.r2.setAttribute("fill", lerpHex(R2[ch.c], R2[Math.min(ch.c + 1, R2.length - 1)], ch.t));
    refs.r3.setAttribute("fill", lerpHex(R3[ch.c], R3[Math.min(ch.c + 1, R3.length - 1)], ch.t));

    /* parallax drift by depth */
    refs.stars.setAttribute("transform", "translate(0," + (-p * 40).toFixed(1) + ")");
    refs.planet.setAttribute("transform", "translate(" + (p * 90).toFixed(1) + "," + (p * -60).toFixed(1) + ")");
    refs.r1.setAttribute("d", ridgePath(lerp(520, 430, p), 130, 24, 3.1));
    refs.r2.setAttribute("d", ridgePath(lerp(660, 560, p), 170, 20, 7.4));
    refs.r3.setAttribute("d", ridgePath(lerp(800, 700, p), 120, 30, 1.2));
    refs.r1.setAttribute("transform", "translate(" + (-p * 60).toFixed(1) + ",0)");
    refs.r2.setAttribute("transform", "translate(" + (p * 70).toFixed(1) + ",0)");

    /* signal wave grows with progress */
    var amp = 10 + p * 130;
    var d = "M0 640 ";
    for (var x = 0; x <= 1440; x += 60) {
      var y = 640 - Math.sin(x * 0.008 + p * 9) * amp * 0.6 - Math.sin(x * 0.02 - p * 6) * amp * 0.4;
      d += "L" + x + " " + y.toFixed(0) + " ";
    }
    d += "L1440 900 L0 900 Z";
    refs.wave.setAttribute("d", d);
    var coreS = 1 + p * 0.6;
    refs.core.setAttribute("r", (22 * coreS).toFixed(1));
    refs.coreG.setAttribute("transform", "translate(0," + (-p * 120).toFixed(1) + ")");
  }

  var slides = [];
  function frame() {
    var track = document.getElementById("storyTrack");
    if (track && slides.length === 0) {
      slides = Array.prototype.slice.call(document.querySelectorAll(".slide"));
    }
    if (track) {
      var total = track.offsetHeight - window.innerHeight;
      var sy = window.scrollY - track.offsetTop;
      var p = Math.min(Math.max(sy / Math.max(total, 1), 0), 1);
      if (refs.sky) render(p);
      /* slides opacity by window */
      var n = slides.length;
      for (var i = 0; i < n; i++) {
        /* slide i owns its chapter band, with soft cross-fades; the first
           slide is visible at the very top and the last until the end */
        var a = i / n - 0.04, b = (i + 1) / n + 0.14;
        var e = 0.07;
        var up = Math.min(Math.max((p - a) / e, 0), 1);
        var dn = Math.min(Math.max((b - p) / e, 0), 1);
        var fade = Math.min(up, dn);
        slides[i].style.opacity = String(Math.max(0, Math.min(1, fade)));
        slides[i].style.transform = "translateY(" + ((0.5 - p) * 70 * (i % 2 ? -1 : 1)).toFixed(0) + "px)";
      }
      body.style.setProperty("--story-p", p.toFixed(3));
    }
    requestAnimationFrame(frame);
  }

  window.__enterSite = function () {
    body.classList.add("story");
    buildScene();
    buildSlides();
  };

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", function () { if (body.classList.contains("story")) { buildScene(); buildSlides(); } });
  }
  requestAnimationFrame(frame);
})();
