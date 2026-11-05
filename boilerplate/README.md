# Pop Boilerplate (AlphaFlux build)

The starting skeleton for any Pop UI page. One file: `pop-boilerplate.html`.
Copy it, swap the palette tokens at the top of the style block for your
product, and replace the blocks marked REPLACE. The build (layout, popups,
rails, settings, agent shell) never changes between products; only the
colors and the content do. The rules it follows live in the pop_ui skill.

## The grid

The page is rows and columns, like a spreadsheet. Content drops into cells
and the grid sizes it:

- `.shell` stacks the page rows.
- `.rail` runs tiles sideways in fixed-width columns with state-aware
  arrows: they fade at the end of a row and disappear when the row fits.
- `.tiles` is an auto-fit grid: add a tile and it takes the next cell at the
  right size on its own. Nobody hand-sizes a box.

The display settings (gear button, bottom left) scale the whole grid at
once by switching `html[data-view]` between sm, md, and lg, which restates
the root font size and the grid cell sizes. The choice saves on the device.

## What is baked in

- One sticky header bar: brand, nav pills, sign in. No phone button in it.
- The hero card: fixed geometry from a height clamp, the whole picture
  fitted inside it (object-fit: contain), words in a glass panel anchored
  low. The card never gets redesigned per page; only the picture and words
  change.
- Tiles: whole picture at its natural ratio, caption overlaid on the
  poster, the whole tile is the tap target. Rails show the top 5 plus a
  View All tile.
- Popups for everything: the base page always stays, every popup leaves a
  blurred sliver of the page around its edges, closes with a big X, the
  scrim, or Escape. Quick view fits whole; deeper layers scroll inside
  their own card. No Close word buttons anywhere.
- A promo popup with the honest price pattern: the full price crossed out
  beside the deal price, the yearly and monthly options shown the same way,
  a code box, and a SCAN ME code that lands on the booking or signup popup.
  Real offers only.
- An agent shell: floating button, chat panel with a signup pill in its
  header and a voice toggle slot after Send.
- Photo metadata: real alt text on every image and a JSON-LD block in the
  head (Organization and WebSite) so search listings and AI search bots
  read the page correctly. Extend the graph per page (plans and prices on
  a pricing page, questions on a help page).
- Photo framing as data: the `.framed` utility stores a photo's visible
  window as three CSS values (position x, position y, zoom) that a future
  settings UI can write after a drag or pinch. Default tiles and heroes do
  not use it; they show the whole picture.
- Footer: one copyright line.

## Starting a new page

1. Copy `pop-boilerplate.html` to the new page name.
2. Set the title, description, canonical, and JSON-LD for the page.
3. Replace the hero picture and words, the rail tiles, and the grid tiles.
   Keep every image's alt text true to what the picture shows.
4. Wire the quick view and promo content, and the agent's brain.
5. Walk the finished page at phone, tablet, laptop, and wide desktop
   sizes and confirm nothing scrolls sideways, nothing is cropped, and
   every popup fits with its sliver showing.
