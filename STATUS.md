<!--
INSTRUCTIONS FOR THIS FILE (user and AI: read before editing. Keep this header.)

WHAT THIS FILE IS: the living right-now snapshot of this project. Anyone
opening the repo cold reads this first and is current in two minutes.
Updated at the end of every work session on the project; a stale STATUS
is worse than none, so when you finish work here, update it in the same
pass.

EXACT STRUCTURE:
- "## As of <YYYY-MM-DD>": the date this snapshot was written.
- "## Proven state": what is done, with its proof named.
- "## In flight": work currently underway and where it stands.
- "## Blocked": each blocker and exactly who or what it waits on.
- "## Next moves": the next three moves, in order, referencing TODO ids.

RULES:
- Claims here carry the same proof standard as TODO checkboxes.
- Keep it short: a snapshot, not a history (history is CHANGELOG.md).
- No secrets, passwords, or tokens in this file, ever.
-->
# STATUS: AlphaFlux Web (alphaflux.net)

## As of 2026-10-28

## Proven state
- Static generator shipped: go run ./build renders content/*.json into committed HTML plus sitemap.xml, robots.txt, llms.txt, and llms-full.txt (proof: README.md, build/; TODO T-0.1)
- Honest status ledger shipped: content/_status.json feeds BUILD-MANIFEST.md, so site copy and build status cannot drift (proof: BUILD-MANIFEST.md; TODO T-0.2)
- Deploy on push live: GitHub webhook to Dokploy, swarm service updated, Traefik serves alphaflux.net (proof: README.md deploy section; the production site is live; TODO T-0.3)
- Public sitemap and llms files published from this repo (proof: sitemap.xml, llms.txt in the tree; TODO T-0.4)
- Self-hosted fonts and icon sprite, no third party font or icon calls (proof: assets/fonts.css, assets/fonts/, assets/icons.svg; TODO T-0.5)
- Page truth today, from BUILD-MANIFEST.md: 5 pages LIVE, 10 PARTIAL, 4 PLANNED, 5 N/A (proof: BUILD-MANIFEST.md)
- Repo standard files complete in this repo, with instruction headers (proof: commits 4844b61 to 54fb444, 2026-10-20)
- Home page is a hand-built coming soon sneak peek (2026-10-28, David's order), modeled on the AlphaFlux poster: coming soon banner, module tile grid, four ways in, free plan box. Deliberate exception to the generator: content/index.json is unchanged and stays the relaunch source; the previous generated index is backed up at backup/index.full-2026-10-28.html (proof: index.html commit cb1ace8; backup commit 722d063)

## In flight
- Phase 2: the ten PARTIAL pages being finished to LIVE, one at a time (TODO Phase 2)
- Phase 1: planned pages and sections being built as the console features behind them land (TODO Phase 1)
- Phase 3 upkeep loop: OG cards, responsive checks, rail and sitemap verification (TODO Phase 3)

## Blocked
- Billing card payments section (waiting on: the console money chain, alphaflux-console TODO T-1.3 / T40; TODO T-1.7)
- White-label page content beyond the planned stub (waiting on: console white-label theming; design in alphaflux-docs docs/whitelabel-theming.md; TODO T-4.1)
- Terms and privacy reconciliation (waiting on: the console Terms of Service publishing after counsel review, alphaflux-console T73; TODO T-3.5)

## Next moves
1. Finish the Overview (index) page from partial to live (TODO T-2.1)
2. Finish the AI agents page from partial to live (TODO T-2.2)
3. Finish the Security page from partial to live (TODO T-2.3)
