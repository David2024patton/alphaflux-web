# AlphaFlux Manifest

Public build log for the AlphaFlux project. Written to be machine readable:
any AI assistant working with David or Skyler should read this file first.

Last updated: 2026-09-15

## Who

- David Patton. Runs marketing and tech for Patriot Pest Control (owner:
  Skyler Rose), builds the AlphaFlux platform, and runs iTaK.
- Stack: Go backends, SurrealDB, Dokploy, Twilio voice and SMS, FieldRoutes
  CRM, Meta lead ads, bare metal CUDA inference.
- Working rules: modular systems, code must compile and run, integer math
  for money, no raw credentials in chat, text, or code. Secrets live in
  secure credential storage only.

## What is built

### Sameday AI voice agent, Patriot Pest Control (live 2026-09-14)

- Agent "Emily" live at 100 percent of traffic on the business line
  (385) 438-3231.
- Full IVR flow: silent caller ID, admin / existing-customer /
  unknown-caller branches, billing and secure payment links, scheduling
  and rescheduling, emergency escalation, cancellation retention, Spanish
  support, solicitor message capture.
- Tools verified live: findOffice, checkJobType, saveCustomer, findSpot,
  scheduleAppointment. Appointment booking works.
- FieldRoutes connected, no errors or warnings.
- Human and emergency transfers route to Skyler Rose. David's personal
  number is never a transfer target.
- Original agent parked as "Emily (Original)" at 0 percent traffic.

### Meta Leads connector (2026-09-14)

- Reads the Patriot Pest Control Solutions Facebook page: 1 page,
  7 lead forms, raw lead data (name, phone, email, pest issue,
  created time).
- A Meta token was exposed over text on 2026-09-14 and must be treated
  as compromised. Rotation is part of the open token work below.

### GitHub skill (2026-09-14, extended 2026-09-15)

- Workspace skill for GitHub repo read plus file write via the contents
  API. This skill maintains this manifest.

### patriotpest-ivr Go service (private repo David2024patton/patriotpest-ivr)

- Twilio voice and SMS backend: caller triage, admin menu,
  existing-customer flow, new-lead sales flow, retention flow with a
  loyalty credit and a service pause option, voicemail transcription,
  Slack and email alerts, abandoned-call callback queue.
- Health endpoint: https://ivr.alphaflux.net/health

## In progress

### Sales / Leads dashboard section (requested by Skyler Rose, 2026-09-14)

- Store unconverted leads: door-to-door leads and online leads where
  contact info was taken but the customer never signed up or never made
  the roster.
- Needs: lead capture with source tagging (D2D vs online), a status
  pipeline, and follow-up tracking.
- Status: specified, not built yet.

### Payroll tab (requested by David Patton, 2026-09-14)

- AlphaFlux owns the data side: employee records, hours, pay rates, pay
  run drafts, pay history, pay stubs.
- Tax calculation, filings, W-2/1099 generation, and direct deposit go
  through a payroll provider API (Gusto or ADP). AlphaFlux does not do
  tax compliance in-house.
- Status: scoped, not built yet.

### Meta ads read-only token (requested by Skyler Rose, 2026-09-14)

- Read-only Meta token for ad spend and CPL feeding the morning briefing.
- Status on 2026-09-15: Skyler could not complete the self-serve setup.
  David is generating a fresh token and will paste it into Skyler's
  assistant's secure credential page. Pending.

### IVR vs Sameday comparison (requested by Skyler Rose, 2026-09-15)

- Skyler's Claude assistant is dissatisfied with Sameday's responses and
  wants an apples to apples comparison against David's IVR system, and
  asks whether they are locked to Sameday.
- Status: Skyler asked David for a link to the system. Pending David's
  reply.

## Open threads

- BJ Bower callback and scheduling (owed by David, requested 2026-09-14).
- FieldRoutes skill verification: Skyler reports 578 customer records
  pulled through his skill; not independently verified yet.
- Visit with Mersades Moore pushed to next week; she is sick
  (2026-09-15).
- Scott Holt asked for David's Venmo on 2026-09-14; still awaiting
  David's reply.
