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

## As of 2026-11-07

## Proven state
- Static generator shipped: go run ./build renders content/*.json into committed HTML plus sitemap.xml, robots.txt, llms.txt, and llms-full.txt (proof: README.md, build/; TODO T-0.1)
- Honest status ledger shipped: content/_status.json feeds BUILD-MANIFEST.md, so site copy and build status cannot drift (proof: BUILD-MANIFEST.md; TODO T-0.2)
- Deploy on push live: GitHub webhook to Dokploy, swarm service updated, Traefik serves alphaflux.net (proof: README.md deploy section; the production site is live; TODO T-0.3)
- Public sitemap and llms files published from this repo (proof: sitemap.xml, llms.txt in the tree; TODO T-0.4)
- Display settings and promo rotations shipped on every page (2026-11-05, David's order): a gear button opens text size and tile size settings saved per device; a timed promo popup runs holiday rotations (Veterans Day, Thanksgiving, Christmas, New Year, Independence Day) and prefers super-admin promos from the console endpoint https://next.alphaflux.net/api/promos/active when it exists, falling back to the baked rotations. Both live in a uniform AFX extras block appended to every page (proof: the afx-gear / afx-promo markup in every *.html; live walk on alphaflux.net)
- Self-hosted fonts and icon sprite, no third party font or icon calls (proof: assets/fonts.css, assets/fonts/, assets/icons.svg; TODO T-0.5)
- Legal pages carry the A2P texting clauses (2026-11-07, David's order): privacy has a Text Messaging / SMS section with the no-share/no-sell promise for mobile and opt-in data, terms has a Text Messaging / SMS Program section (AlphaFlux program, STOP/HELP, support (502) 385-0241, consent not a condition of purchase), mirrored in content/privacy.json and content/terms.json; default.conf serves /privacy-policy and /terms-of-use from the same files (proof: live fetches of /privacy-policy and /terms-of-use returning the new sections)
- Page truth today, from BUILD-MANIFEST.md: 5 pages LIVE, 10 PARTIAL, 4 PLANNED, 5 N/A (proof: BUILD-MANIFEST.md)
- Repo standard files complete in this repo, with instruction headers (proof: commits 4844b61 to 54fb444, 2026-10-20)
- Home page is a hand-built coming soon sneak peek (2026-10-28, David's order), modeled on the AlphaFlux poster: coming soon banner, module tile grid, four ways in, free plan box. Deliberate exception to the generator: content/index.json is unchanged and stays the relaunch source; the previous generated index is backed up at backup/index.full-2026-10-28.html (proof: index.html commit cb1ace8; backup commit 722d063)
- Whole site converted to the Sunny Pop design (2026-11-01, David's order), extending the hand-built exception above to every page. The home page was rebuilt as a streaming style Sunny Pop page and went live first (commit 4576ff4), then the pricing page (commit 62a1526), then the rest of the site in the same shell: the 20 content pages (ai-agents, analytics, billing, calls-and-messaging, crm, developers, email-marketing, field-operations, fleet, industries, inventory, marketing, modules, platform, programs, scheduling, security, sitemap, team, white-label) restyled with their real text ported word for word (the only loss: a dead Copy button label on developers), privacy and terms restyled with their legal text word for word, a new Sunny Pop 404, and 17 new pages created so every home page tile lands on a real page (bookings, invoices, team-app, setup, import, connect, go-live, social, ads, reviews, websites, marketplace, help, tutorials, about, contact, signup, signin). content/ and build/ are untouched: the generator does not produce these pages until it is taught the Sunny Pop shell, so edit the committed HTML directly. The pre-conversion version of every replaced page is in git history at 62a1526 and archived with the project records.

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
