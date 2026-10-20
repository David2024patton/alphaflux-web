# TODO: AlphaFlux Web (alphaflux.net)

The detailed task list for AlphaFlux Web. Rules: a box is checked only
when its proof exists and is named (a test, a log line, a bench note, a
live check). Emulation-proven and bench-proven are marked differently.
Plan steps live in PLAN.md; each phase here names its PLAN section.

Page and section statuses live in content/_status.json and are rendered
into BUILD-MANIFEST.md by the build. A page item here is done when its
status flips to live there, never before.

## Phase 0: Generator and hosting (shipped) (PLAN: Build order step 1)
- [x] Static generator: `go run ./build` renders content/*.json into committed HTML plus sitemap.xml, robots.txt, llms.txt, and llms-full.txt (proof: README.md; `build/`)
- [x] Honest status ledger: content/_status.json feeds BUILD-MANIFEST.md, so the site copy and its build status cannot drift (proof: BUILD-MANIFEST.md)
- [x] Deploy on push: GitHub webhook to Dokploy, swarm service updated, Traefik serves alphaflux.net (proof: README.md deploy section; the production site is live)
- [x] The public XML sitemap and llms files are published from this repo (the console is login-gated, so the public sitemap belongs here) (proof: sitemap.xml, llms.txt in the tree)
- [x] Self-hosted fonts and the icon sprite, no third party font or icon calls (proof: assets/fonts.css, assets/fonts/, assets/icons.svg)

## Phase 1: Planned pages and sections to build (PLAN: Build order step 3)
- [ ] White-label page: status planned (proof: status live in content/_status.json, BUILD-MANIFEST.md regenerated, page checked at phone size)
- [ ] Fleet page: status planned (proof: same as above)
- [ ] Field operations, Proof of work section: status planned (proof: section status live in content/_status.json)
- [ ] Field operations, Offline that recovers section: status planned (proof: same as above)
- [ ] Scheduling, Appointment reminders section: status planned (proof: same as above)
- [ ] Inventory, Assets as well as consumables section: status planned (proof: same as above)
- [ ] Billing, Card payments section: status planned (blocked on: the console money chain, alphaflux-console TODO Phase 1 T40)

## Phase 2: Partial pages to finish (PLAN: Build order step 2)
Current counts from BUILD-MANIFEST.md: 5 LIVE, 10 PARTIAL, 4 PLANNED, 5 N/A.
Each item: proof is the page status flipping to live in content/_status.json,
BUILD-MANIFEST.md regenerated, and the page checked at phone and desktop sizes.
- [ ] Overview (index): partial to live
- [ ] AI agents: partial to live
- [ ] Security: partial to live
- [ ] Field operations: partial to live
- [ ] Scheduling: partial to live
- [ ] Inventory: partial to live
- [ ] Billing: partial to live
- [ ] Email marketing: partial to live
- [ ] Marketing: partial to live
- [ ] Developers: partial to live

## Phase 3: Site upkeep (PLAN: Build order step 4)
- [ ] After any copy change, regenerate the Open Graph cards (scripts/make-og.sh) and commit the PNGs with the copy
- [ ] Run the responsive check (scripts/check-responsive.mjs) before calling any page done
- [ ] After adding a page, verify the build wrote it into the rail, the footer, sitemap.xml, and both llms files (the build fails on a rail entry with no page, a duplicate slug, or a missing icon)
- [ ] Keep the Dokploy and Traefik notes in README.md true: if the Dokploy app is ever recreated, the swarm service suffix changes and /etc/dokploy/traefik/dynamic/alphaflux-web.yml must be updated in the same move
- [ ] Terms and privacy pages reconciled with the console Terms of Service when it publishes (alphaflux-console T73)

## Parked (waiting on a decision, hardware, or hands)
- White-label page content beyond the planned stub (blocked on: console white-label theming; design captured, not started, in alphaflux-docs docs/whitelabel-theming.md)
