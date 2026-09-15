import { createRequire } from 'node:module';
const require = createRequire(import.meta.url);
const { chromium } = await import('playwright');

const BASE = process.argv[2] || 'http://127.0.0.1:8099';
const PAGES = ['index','pricing','fleet','inventory','field-operations','email-marketing',
               'ai-agents','modules','programs','industries','white-label','developers','sitemap','security','team','platform','crm','billing','marketing','analytics','scheduling','phone-and-messaging'];
const WIDTHS = [320, 375, 414, 768, 834, 1024, 1280, 1920];

const b = await chromium.launch();
let fails = 0, checks = 0;
const rows = [];

for (const slug of PAGES) {
  for (const w of WIDTHS) {
    const p = await b.newPage({ viewport: { width: w, height: 900 }, hasTouch: w < 1024 });
    const issues = [];
    p.on('pageerror', e => issues.push('js: ' + e.message));
    await p.goto(`${BASE}/${slug}.html`, { waitUntil: 'load' });
    await p.waitForTimeout(350);
    const r = await p.evaluate(() => {
      const vw = window.innerWidth;
      const de = document.documentElement;
      const unclipped = [];
      const smallTargets = [];
      for (const el of document.querySelectorAll('body *')) {
        const rc = el.getBoundingClientRect();
        if (rc.width === 0 && rc.height === 0) continue;
        if (rc.right > vw + 1 || rc.left < -1) {
          let clipped = false;
          for (let a = el.parentElement; a && a !== document.body; a = a.parentElement) {
            const ox = getComputedStyle(a).overflowX;
            if (ox !== 'visible') { clipped = true; break; }
          }
          if (!clipped) unclipped.push(el.tagName.toLowerCase() + '.' + String(el.className).split(' ')[0] + ' r=' + Math.round(rc.right));
        }
      }
      // tap targets: interactive controls a finger must hit
      for (const el of document.querySelectorAll('a[href], button, [role="button"], summary, input, select')) {
        const rc = el.getBoundingClientRect();
        if (rc.width === 0 || rc.height === 0) continue;
        if (getComputedStyle(el).display === 'none' || getComputedStyle(el).visibility === 'hidden') continue;
        if (rc.height < 24 || rc.width < 24) smallTargets.push((el.textContent || el.tagName).trim().slice(0, 24) + ' ' + Math.round(rc.width) + 'x' + Math.round(rc.height));
      }
      const railVisible = !!document.querySelector('.rail') && getComputedStyle(document.querySelector('.rail')).display !== 'none';
      const toggleVisible = !!document.querySelector('.nav-toggle') && getComputedStyle(document.querySelector('.nav-toggle')).display !== 'none';
      return { vw, doc: de.scrollWidth, unclipped: unclipped.slice(0, 6), small: [...new Set(smallTargets)].slice(0, 8), railVisible, toggleVisible };
    });
    checks++;
    const bad = [];
    if (r.doc > r.vw + 1) bad.push(`doc scrollWidth ${r.doc} > ${r.vw}`);
    if (r.unclipped.length) bad.push('unclipped: ' + r.unclipped.join(' | '));
    if (r.small.length && w < 1024) bad.push('small targets: ' + r.small.join(' | '));
    // navigation must exist at every width: rail on wide, toggle on narrow
    if (w >= 1024 && !r.railVisible) bad.push('rail hidden at ' + w);
    if (w < 1024 && !r.toggleVisible) bad.push('no nav toggle at ' + w);
    if (issues.length) bad.push(issues.join(' | '));
    if (bad.length) { fails++; rows.push(`FAIL ${slug}@${w}: ${bad.join('; ')}`); }
    await p.close();
  }
}
await b.close();
console.log(rows.length ? rows.join('\n') : '');
console.log(`\n${checks} checks, ${fails} failed`);
