<!--
INSTRUCTIONS FOR THIS FILE (user and AI: read before editing. Keep this header.)

WHAT THIS FILE IS: the human-readable milestone trail of this project:
what shipped, dated, newest first. Git history is the raw record; this
is the curated one a person can read in a minute.

EXACT STRUCTURE:
- "## <YYYY-MM-DD>: <milestone in one line>" entries, newest first.
- Under each: what changed, and the proof or commit it landed in.

RULES:
- An entry is added when something SHIPS (lands on the main branch or
  goes live), not when it is started.
- One entry per milestone, not per commit.
- No secrets, passwords, or tokens in this file, ever.
-->
# CHANGELOG: AlphaFlux Web (alphaflux.net)

Curated milestones from the git history of this repo, newest first. Git history is the raw record.

## 2026-10-20: Repo standard complete, end goal recorded
- TODO, IDEA, PLAN, ENDGOAL, README, and docs/RULES.md brought to the repo standard with stable TODO ids, PLAN cross-references, and instruction headers (commits 4844b61 to 54fb444)

## 2026-09-29: Go directive bumped
- Go directive moved to 1.27.1 and the local build binary ignored (commit 1eab8f8)

## 2026-09-19: Privacy policy and terms added
- The missing privacy policy and terms of use were added, and the tracking claim in the privacy policy corrected (commits 36b52c7, 93e2e1b)

## 2026-09-16: Build manifest generated
- The generator began producing BUILD-MANIFEST.md from content and status, so the site can act as the feature spec, and the README documented the manifest and status file (commits 9644570, 1be60b8)

## 2026-09-15: Multi-page product site rebuilt
- alphaflux.net rebuilt as a multi-page product site, the generator made deterministic, deploy settings and front-door routing recorded in the README (commits 685fa0e, 51569f8, 2fc8572)

## 2026-09-13: Marketing site first shipped
- AlphaFlux marketing and landing site (nginx static, served at alphaflux.net) (commit bb377bb)
