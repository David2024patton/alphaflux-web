<!--
INSTRUCTIONS FOR THIS FILE (user and AI: read before editing. Keep this header.)

WHAT THIS FILE IS: the settled-decision log for this project. IDEA.md is
the raw capture of what David said; this file is only the calls that are
FINAL, so no one re-asks a settled question or quietly reverses a call.

EXACT STRUCTURE: one entry per decision, newest first:
## <YYYY-MM-DD>: <decision in one line>
- **Decided:** <what was picked>
- **Why:** <the reason, in plain words>
- **Supersedes:** <the earlier call it replaces, or "none">
- **By:** David | <who, with David's approval noted>

RULES:
- A decision enters here the day it is made, in the same pass as any
  PLAN or TODO change it causes.
- Reversing a decision means a NEW entry that names what it supersedes;
  never edit an old entry into a different decision.
- No secrets, passwords, or tokens in this file, ever.
-->
# DECISIONS: AlphaFlux Web (alphaflux.net)

## 2026-10-20: This site is the public face and the front door of the platform
- **Decided:** ENDGOAL.md records the end state: this repo presents the platform end goal dictated for alphaflux-console, and its nginx container is the platform front door, serving the public pages and proxying the console paths
- **Why:** David's platform end goal dictation on 2026-10-20, written into ENDGOAL.md the same day
- **Supersedes:** none
- **By:** David

## 2026-09-16: The site is the specification, and the manifest is the honest ledger
- **Decided:** The site describes the product intended to be sold, including parts designed and not yet built; BUILD-MANIFEST.md, generated from the same content as the pages, is the honest answer for where each page, section, and feature stands (LIVE, PARTIAL, PLANNED, N/A)
- **Why:** So the site can act as the feature spec without ever claiming a capability that has not shipped (commits 9644570, 1be60b8)
- **Supersedes:** none
- **By:** David

## 2026-09-15: The site is a generated multi-page site with committed output
- **Decided:** A Go generator (standard library only) renders one JSON content file per page into committed static HTML; the container build needs no Go toolchain and nginx serves plain static files; the generator is deterministic
- **Why:** The served site is plain files, the copy source of truth is content, and what is committed is exactly what is served (commits 685fa0e, 51569f8)
- **Supersedes:** the first single-page landing site of 2026-09-13
- **By:** David

## 2026-09-15: This container is the front door for the console paths
- **Decided:** alphaflux.net routing points at this service, and its nginx proxies the console paths (/login, /console, /api, /phone, /twilio, /scan, /patrol, /admin, /waitlist, /accept) to the alphaflux-console app on port 8080; routing lives in the Traefik dynamic file on the server, documented in README.md
- **Why:** One public domain for the marketing site and the platform, recorded in README.md and alphaflux-console DEPLOY.md
- **Supersedes:** none
- **By:** David
