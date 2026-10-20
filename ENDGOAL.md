# ENDGOAL: AlphaFlux Web (alphaflux.net)

The end state this project exists to reach, in plain words. TODO.md and
PLAN.md are how we get here; this file is how we know we are done. Work
continues until this end goal is reached. When David restates or extends
the end goal, this file is updated the same day and PLAN/TODO follow.

## The end state

AlphaFlux Web is the public face of the platform whose end state is
written in alphaflux-console ENDGOAL.md, and its front door. A business
owner who lands on alphaflux.net understands in one visit what
AlphaFlux does, sees the guided business setup story (sign up, be
walked through everything, be operating within an hour or two), and
walks through the front door into signup or sign in without ever
leaving the platform's own domain, because this site's nginx container
is the door: it serves the public pages and proxies the console paths
to the console app behind it.

The site is the specification, and in the end state the specification
is simply true: every capability the site describes is a shipped
capability. The status ledger that today honestly splits the site into
LIVE, PARTIAL, and PLANNED reads LIVE all the way down, and
BUILD-MANIFEST.md proves it from the same content the pages are built
from.

## How we will know it is reached

- Every page and section in content/_status.json reads LIVE; no
  PARTIAL or PLANNED remains in BUILD-MANIFEST.md.
- The site presents the platform end state (guided setup, the four
  ways in, the builder, the marketplace, the social wall) and every
  claim on it is backed by a live console capability.
- The front door is verified: the public pages serve from this
  container and the console paths proxy through to the console app,
  both checked live.
- New pages keep arriving whole: rail entry, footer, sitemap, and both
  llms files, with OG cards regenerated, as the upkeep loop requires.

## Not in the end state

- Tenant websites do not live here; they are a console feature (the
  website builder). This repo is the platform's own site only.
- No page copy claims a feature the console has not shipped; the status
  ledger is the guard, and it stays in the build.
