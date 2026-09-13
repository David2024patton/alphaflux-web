/* Parallax scene + Atropos cards for the "Enter site" view */
(function () {
  "use strict";
  var NS = "http://www.w3.org/2000/svg";

  function ridgePath(cx, amp, peaks, seed) {
    var pts = [];
    for (var i = 0; i <= peaks; i++) {
      var x = (i / peaks) * 1700 - 130;
      var wave = Math.abs(Math.sin(i * 1.7 + seed)) * 0.6 + Math.abs(Math.sin(i * 0.9 + seed * 2.3)) * 0.4;
      var y = cx - wave * amp;
      pts.push(x.toFixed(0) + "," + y.toFixed(0));
    }
    return "M" + pts.join(" L") + " L1700 520 L-130 520 Z";
  }

  function el(name, attrs, parent) {
    var e = document.createElementNS(NS, name);
    for (var k in attrs) { if (attrs.hasOwnProperty(k)) e.setAttribute(k, attrs[k]); }
    if (parent) parent.appendChild(e);
    return e;
  }

  function buildScene() {
    var host = document.getElementById("parScene");
    if (!host || host.childNodes.length) return;
    var svg = el("svg", { viewBox: "0 0 1440 900", preserveAspectRatio: "xMidYMax slice" }, host);
    svg.style.cssText = "position:absolute;inset:0;width:100%;height:100%";

    var defs = el("defs", {}, svg);
    defs.innerHTML =
      '<linearGradient id="psky" x1="0" y1="0" x2="0" y2="1"><stop offset="0" stop-color="#03050d"/><stop offset="0.55" stop-color="#0b1230"/><stop offset="1" stop-color="#1d1440"/></linearGradient>' +
      '<radialGradient id="pglow" cx="50%" cy="50%" r="50%"><stop offset="0" stop-color="#35d6ff" stop-opacity="0.5"/><stop offset="1" stop-color="#35d6ff" stop-opacity="0"/></radialGradient>' +
      '<radialGradient id="pplanet"><stop offset="0" stop-color="#bff4ff"/><stop offset="0.5" stop-color="#28e3ff"/><stop offset="1" stop-color="#8b6bff"/></radialGradient>';
    el("rect", { width: "1440", height: "900", fill: "url(#psky)" }, svg);

    function layer(spd) {
      return el("g", { class: "par-layer", "data-spd": String(spd) }, svg);
    }

    var stars = layer(0.06);
    for (var i = 0; i < 150; i++) {
      el("circle", {
        cx: (Math.random() * 1440).toFixed(0), cy: (Math.random() * 680).toFixed(0),
        r: (Math.random() * 1.4 + 0.3).toFixed(2), fill: "#9fc4ff",
        opacity: (Math.random() * 0.7 + 0.2).toFixed(2)
      }, stars);
    }
    var neb = layer(0.1);
    for (var n = 0; n < 4; n++) {
      el("circle", {
        cx: (180 + Math.random() * 1020).toFixed(0), cy: (100 + Math.random() * 400).toFixed(0),
        r: (180 + Math.random() * 220).toFixed(0), fill: n % 2 ? "#8b6bff" : "#28e3ff", opacity: "0.06"
      }, neb);
    }
    var planet = layer(0.14);
    el("circle", { cx: "1180", cy: "300", r: "120", fill: "url(#pplanet)" }, planet);
    el("ellipse", { cx: "1180", cy: "300", rx: "210", ry: "42", fill: "none", stroke: "#35d6ff", "stroke-width": "6", opacity: "0.55", transform: "rotate(-16 1180 300)" }, planet);
    el("circle", { cx: "260", cy: "190", r: "270", fill: "url(#pglow)" }, planet);

    function ridge(spd, cx, amp, peaks, seed, fill) {
      var g = layer(spd);
      el("path", { d: ridgePath(cx, amp, peaks, seed), fill: fill }, g);
      return g;
    }
    ridge(0.22, 620, 120, 24, 3.1, "#2b2150");
    ridge(0.34, 750, 160, 20, 7.4, "#1c1740");
    ridge(0.48, 880, 110, 30, 1.2, "#0d0f24");

    var core = layer(0.62);
    el("circle", { cx: "720", cy: "560", r: "26", fill: "#ffffff", opacity: "0.9" }, core);
    el("circle", { cx: "720", cy: "560", r: "66", fill: "url(#pglow)" }, core);
    el("path", { d: "M0 585 Q360 500 720 560 T1440 585 L1440 900 L0 900 Z", fill: "#0a1024", opacity: "0.9" }, core);
  }

  function tick() {
    var sy = window.scrollY || 0;
    var hero = document.getElementById("hero");
    var end = hero ? Math.max(hero.offsetHeight - window.innerHeight * 0.55, 200) : window.innerHeight * 1.4;
    var scene = document.getElementById("parScene");
    var layers = scene ? scene.querySelectorAll(".par-layer") : [];
    for (var i = 0; i < layers.length; i++) {
      var spd = parseFloat(layers[i].getAttribute("data-spd") || "0.2");
      layers[i].style.transform = "translate3d(0," + (-sy * spd * 0.6).toFixed(1) + "px,0)";
    }
    if (scene) {
      var fade = 1 - Math.max(0, sy - end) / Math.max(window.innerHeight * 0.5, 1);
      scene.style.opacity = String(Math.max(0, Math.min(1, fade)));
    }
    requestAnimationFrame(tick);
  }

  function initAtropos() {
    if (!window.Atropos) return;
    var cards = document.querySelectorAll(".atropos");
    for (var c = 0; c < cards.length; c++) {
      window.Atropos({ el: cards[c], rotateTouch: "scroll-y", shadow: true, highlight: true, activeOffset: 30 });
    }
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", function () { buildScene(); initAtropos(); });
  } else {
    buildScene();
    initAtropos();
  }
  requestAnimationFrame(tick);
})();
