# AlphaFlux Web Landing (filler)

## What it is

A single-file marketing landing page for the alphaflux.net apex domain while the
real alphaflux-web marketing site (blueprint Part 3) is still on the roadmap.
It brands AlphaFlux as an AI-native marketing and AI conversion company and is
built to feel like the product pitch: "Find the signal. Convert the demand."

The centerpiece is the Web GPS Layer, an interactive radar simulation. It is the
visual metaphor for the company's positioning (locate demand across the web,
score it, convert it) and doubles as a mini game: conversion nodes appear on the
radar, the sweep locks them, and the visitor clicks to capture them before the
signal fades. 45 second rounds, combos for chaining captures, live mission log
with typed telemetry, HUD clock, demand feed ticker, flux particle background.

## How it works

- One self-contained `index.html`. No external assets, no build step, works
  offline. Canvas background (starfield + flux particles + burst rings), radar
  canvas with its own render loop, HUD clock, rotating demand feed.
- Deployment: nginx container on the local server (Docker Swarm service
  `alphaflux-web` on the dokploy-network overlay), routed by traefik through a
  hand-written dynamic file `/etc/dokploy/traefik/dynamic/alphaflux-web.yml`.
- Routing: `Host(alphaflux.net) || Host(www.alphaflux.net)` on the web and
  websecure entrypoints with `certResolver: letsencrypt`. Routers carry an
  explicit `priority: 1000` so they beat the Dokploy routers that previously
  served those hosts (Dokploy stays reachable at dash.alphaflux.net, which is
  not touched).
- DNS: alphaflux.net and www.alphaflux.net already resolve to the local server
  public IP 107.90.116.157, so no DNS change was needed.

## Where the idea came from

- David's request (2026-09-03): "put some kind of filler landing page for the
  alpha flux domain... some badass... the web GPS layer or something... some
  kind of mini game... related to my alpha flux AI and marketing and AI
  conversion company."
- Blueprint C:\Users\David\Desktop\alphaflux-saas.md names the future repo
  alphaflux-web for exactly this role (marketing / landing site) and describes
  the SaaS where each company gets its own dashboard, so the radar doubles as
  a preview of the product concept.

## Deploy / rollback

Deploy:
1. `docker build -t alphaflux-web:latest /home/server/alphaflux-web`
2. `docker service create --name alphaflux-web --network dokploy-network --replicas 1 alphaflux-web:latest`
3. Write /etc/dokploy/traefik/dynamic/alphaflux-web.yml (traefik watches the dir).
4. Verify https://alphaflux.net serves the page and https://dash.alphaflux.net
   still shows the Dokploy console.

Rollback: `docker service rm alphaflux-web` and delete
/etc/dokploy/traefik/dynamic/alphaflux-web.yml. The Dokploy routers for the
apex are untouched in dokploy.yml, so alphaflux.net falls back to Dokploy
immediately (that was the pre-existing behavior).
