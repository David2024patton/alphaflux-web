<!--
INSTRUCTIONS FOR THIS FILE (user and AI: read before editing. Keep this header.)

WHAT THIS FILE IS: the super-detailed task list for this project. It is the
only place work items live. PLAN.md says why and in what order; ENDGOAL.md
says what finished looks like; this file says exactly what remains to do.

EXACT STRUCTURE:
- Phases in order: "## Phase <n>: <name> (PLAN S<m>)".
- Parked work in its own phase at the end: "## Parked".
- Every item is one checkbox line: "- [ ] T-<phase>.<n> <task> (proof: <named
  proof>) (PLAN S<m>)". Parked items use T-P.<n> and name what blocks them.

RULES:
- A box is checked ONLY when its named proof exists (a test, a log line, a
  bench note, a live check). Never check a box on intention or partial work.
- Emulation-proven and bench-proven are different claims; write which one.
- Item IDs and (PLAN S<n>) tags are stable once published; never renumber
  casually. A deliberate renumber updates PLAN.md in the same pass.
- New work enters as new items at the end of their phase. Completed items
  stay, checked, as the record.
- No secrets, passwords, or tokens in this file, ever.
-->
# TODO: AlphaFlux Web (alphaflux.net)

The detailed task list for AlphaFlux Web. Rules: a box is checked only
when its proof exists and is named (a test, a log line, a bench note, a
live check). Emulation-proven and bench-proven are marked differently.
Every item carries a stable ID (T-<phase>.<n>) and ends with the PLAN.md
section it serves, so one item plus its one plan section is the whole job.

Page and section statuses live in content/_status.json and are rendered
into BUILD-MANIFEST.md by the build. A page item here is done when its
status flips to live there, never before.

## Phase 0: Generator and hosting (shipped) (PLAN S4)
- [x] T-0.1 Static generator: `go run ./build` renders content/*.json into committed HTML plus sitemap.xml, robots.txt, llms.txt, and llms-full.txt (proof: README.md; `build/`) (PLAN S4)
- [x] T-0.2 Honest status ledger: content/_status.json feeds BUILD-MANIFEST.md, so the site copy and its build status cannot drift (proof: BUILD-MANIFEST.md) (PLAN S4)
- [x] T-0.3 Deploy on push: GitHub webhook to Dokploy, swarm service updated, Traefik serves alphaflux.net (proof: README.md deploy section; the production site is live) (PLAN S4)
- [x] T-0.4 The public XML sitemap and llms files are published from this repo (the console is login-gated, so the public sitemap belongs here) (proof: sitemap.xml, llms.txt in the tree) (PLAN S4)
- [x] T-0.5 Self-hosted fonts and the icon sprite, no third party font or icon calls (proof: assets/fonts.css, assets/fonts/, assets/icons.svg) (PLAN S4)

## Phase 1: Planned pages and sections to build (PLAN S6)
- [ ] T-1.1 White-label page: status planned (proof: status live in content/_status.json, BUILD-MANIFEST.md regenerated, page checked at phone size) (PLAN S6)
- [ ] T-1.2 Fleet page: status planned (proof: same as above) (PLAN S6)
- [ ] T-1.3 Field operations, Proof of work section: status planned (proof: section status live in content/_status.json) (PLAN S6)
- [ ] T-1.4 Field operations, Offline that recovers section: status planned (proof: same as above) (PLAN S6)
- [ ] T-1.5 Scheduling, Appointment reminders section: status planned (proof: same as above) (PLAN S6)
- [ ] T-1.6 Inventory, Assets as well as consumables section: status planned (proof: same as above) (PLAN S6)
- [ ] T-1.7 Billing, Card payments section: status planned (blocked on: the console money chain, alphaflux-console TODO Phase 1 T40) (PLAN S6)

## Phase 2: Partial pages to finish (PLAN S5)
Current counts from BUILD-MANIFEST.md: 5 LIVE, 10 PARTIAL, 4 PLANNED, 5 N/A.
Each item: proof is the page status flipping to live in content/_status.json,
BUILD-MANIFEST.md regenerated, and the page checked at phone and desktop sizes.
- [ ] T-2.1 Overview (index): partial to live (PLAN S5)
- [ ] T-2.2 AI agents: partial to live (PLAN S5)
- [ ] T-2.3 Security: partial to live (PLAN S5)
- [ ] T-2.4 Field operations: partial to live (PLAN S5)
- [ ] T-2.5 Scheduling: partial to live (PLAN S5)
- [ ] T-2.6 Inventory: partial to live (PLAN S5)
- [ ] T-2.7 Billing: partial to live (PLAN S5)
- [ ] T-2.8 Email marketing: partial to live (PLAN S5)
- [ ] T-2.9 Marketing: partial to live (PLAN S5)
- [ ] T-2.10 Developers: partial to live (PLAN S5)

## Phase 3: Site upkeep (PLAN S7)
- [ ] T-3.1 After any copy change, regenerate the Open Graph cards (scripts/make-og.sh) and commit the PNGs with the copy (PLAN S7)
- [ ] T-3.2 Run the responsive check (scripts/check-responsive.mjs) before calling any page done (PLAN S7)
- [ ] T-3.3 After adding a page, verify the build wrote it into the rail, the footer, sitemap.xml, and both llms files (the build fails on a rail entry with no page, a duplicate slug, or a missing icon) (PLAN S7)
- [ ] T-3.4 Keep the Dokploy and Traefik notes in README.md true: if the Dokploy app is ever recreated, the swarm service suffix changes and /etc/dokploy/traefik/dynamic/alphaflux-web.yml must be updated in the same move (PLAN S7)
- [ ] T-3.5 Terms and privacy pages reconciled with the console Terms of Service when it publishes (alphaflux-console T73) (PLAN S7)

## Parked (waiting on a decision, hardware, or hands)
- [ ] T-4.1 White-label page content beyond the planned stub (blocked on: console white-label theming; design captured, not started, in alphaflux-docs docs/whitelabel-theming.md) (PLAN S6)
