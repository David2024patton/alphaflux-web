/* ALPHAFLUX top sliver: screenshot, view as markdown, clear cache */
(function () {
  "use strict";
  var sliver = document.getElementById("sliver");
  if (!sliver) return;

  function toast(msg, bad) {
    var t = document.createElement("div");
    t.className = "sl-toast" + (bad ? " bad" : "");
    t.textContent = msg;
    document.body.appendChild(t);
    setTimeout(function () { t.style.opacity = "0"; setTimeout(function () { t.remove(); }, 400); }, 2600);
  }

  function overlay(title, bodyNode) {
    var o = document.createElement("div");
    o.className = "sl-modal";
    var head = document.createElement("div");
    head.className = "sl-modal-head";
    head.innerHTML = "<b>" + title + "</b>";
    var close = document.createElement("button");
    close.className = "sl-x";
    close.textContent = "Close";
    close.addEventListener("click", function () { o.remove(); });
    head.appendChild(close);
    o.appendChild(head);
    o.appendChild(bodyNode);
    document.body.appendChild(o);
  }

  function snapshotWebGL(canvas) {
    try {
      var d = canvas.toDataURL("image/png");
      if (d && d.length > 1000) return d;
    } catch (e) {}
    return null;
  }

  sliver.addEventListener("click", function (e) {
    var el = e.target;
    var act = el && el.getAttribute ? el.getAttribute("data-act") : null;
    if (!act) return;

    if (act === "shot") {
      if (!window.html2canvas) { toast("Screenshot engine not loaded", true); return; }
      toast("Capturing page...");
      window.html2canvas(document.body, {
        scale: Math.min(window.devicePixelRatio || 1, 2),
        useCORS: true,
        backgroundColor: "#05070f",
        logging: false,
        onclone: function (doc) {
          var real = document.querySelectorAll("canvas");
          var list = doc.querySelectorAll("canvas");
          for (var i = 0; i < list.length; i++) {
            var c = list[i];
            var src = real[i] ? snapshotWebGL(real[i]) : null;
            if (!src) continue;
            var img = doc.createElement("img");
            img.src = src;
            img.style.cssText = "position:absolute;left:0;top:0;width:100%;height:100%";
            img.width = c.width; img.height = c.height;
            c.parentNode.insertBefore(img, c);
            c.style.display = "none";
          }
        }
      }).then(function (canvas) {
        var a = document.createElement("a");
        a.download = "alphaflux-" + Date.now() + ".png";
        a.href = canvas.toDataURL("image/png");
        a.click();
        toast("Screenshot saved");
      }).catch(function (err) {
        toast("Screenshot failed: " + (err && err.message ? err.message : err), true);
      });
      return;
    }

    if (act === "md") {
      var parts = ["# AlphaFlux"];
      var seen = {};
      var nodes = document.querySelectorAll(".sl-top, .slide h2, .slide p, main h1, main h2, main h3, .eyebrow, .hero-sub, .kicker, .step h3, .step p, .chip h4, .chip p, .sec-head .tag");
      for (var i = 0; i < nodes.length; i++) {
        var n = nodes[i];
        var txt = (n.textContent || "").replace(/\s+/g, " ").trim();
        if (!txt || seen[txt]) continue;
        seen[txt] = 1;
        var tag = n.tagName.toLowerCase();
        var cls = typeof n.className === "string" ? n.className : "";
        if (tag === "h1") parts.push("\n# " + txt);
        else if (tag === "h2") parts.push("\n## " + txt);
        else if (tag === "h3") parts.push("\n### " + txt);
        else if (tag === "p") parts.push(txt);
        else if (cls.indexOf("sl-top") >= 0) parts.push("\n## " + txt);
        else if (cls.indexOf("eyebrow") >= 0) parts.push("\n> " + txt);
        else if (cls.indexOf("tag") >= 0) parts.push("_" + txt + "_");
        else parts.push(txt);
      }
      var md = parts.join("\n\n");
      var pre = document.createElement("textarea");
      pre.className = "sl-md";
      pre.value = md;
      pre.readOnly = true;
      var wrap = document.createElement("div");
      var row = document.createElement("div");
      row.className = "sl-md-actions";
      var copy = document.createElement("button");
      copy.className = "sl-mini";
      copy.textContent = "Copy";
      copy.addEventListener("click", function () { pre.select(); try { document.execCommand("copy"); } catch (e) {} toast("Copied"); });
      var dl = document.createElement("button");
      dl.className = "sl-mini";
      dl.textContent = "Download .md";
      dl.addEventListener("click", function () {
        var a = document.createElement("a");
        a.download = "alphaflux.md";
        a.href = "data:text/markdown;charset=utf-8," + encodeURIComponent(md);
        a.click();
      });
      row.appendChild(copy); row.appendChild(dl);
      wrap.appendChild(row);
      wrap.appendChild(pre);
      overlay("View as Markdown", wrap);
      return;
    }

    if (act === "cache") {
      toast("Purging browser cache...");
      var jobs = [];
      if (window.caches && window.caches.keys) {
        jobs.push(window.caches.keys().then(function (keys) {
          return Promise.all(keys.map(function (k) { return window.caches.delete(k); }));
        }));
      }
      Promise.all(jobs).then(function () {
        toast("Cache cleared, reloading fresh copy");
        setTimeout(function () { location.href = location.pathname + "?v=" + Date.now(); }, 400);
      }).catch(function (err) {
        toast("Could not clear cache: " + err, true);
      });
    }
  });
})();
