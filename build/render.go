package main

import (
	"encoding/json"
	"strings"
)

/* ---------- document shell ---------- */

func renderPage(p *Page) string {
	var b strings.Builder
	crumb := p.Crumb
	if crumb == "" {
		crumb = navItem(p.Slug).Label
	}
	path := urlOf(p.Slug)

	b.WriteString("<!doctype html>\n")
	b.WriteString(`<html lang="en" data-theme="dark">` + "\n")
	b.WriteString("<head>\n")
	b.WriteString(metaBlock(p, path))
	b.WriteString("</head>\n")
	b.WriteString("<body>\n")
	b.WriteString(skipLink())
	b.WriteString(`<div class="shell">` + "\n")
	b.WriteString(rail(p.Slug))
	b.WriteString(`<div class="main">` + "\n")
	b.WriteString(topbar(p, crumb))
	b.WriteString(`<main id="main">` + "\n")

	if p.Slug != "index" {
		b.WriteString(pageHead(p, crumb))
	}

	for i, s := range p.Sections {
		b.WriteString(renderSection(p, s, i))
	}

	if len(p.Plans) > 0 {
		b.WriteString(renderPlans(p))
	}

	if len(p.FAQ) > 0 {
		b.WriteString(renderFAQ(p))
	}

	b.WriteString(`</main>` + "\n")
	b.WriteString(ctaBand(p))
	b.WriteString(footer())
	b.WriteString("</div>\n</div>\n")
	b.WriteString(drawer())
	b.WriteString(`<script src="/assets/site.js?v=` + assetVersion + `" defer></script>` + "\n")
	b.WriteString("</body>\n</html>\n")
	return b.String()
}

func render404() string {
	var b strings.Builder
	b.WriteString("<!doctype html>\n<html lang=\"en\" data-theme=\"dark\">\n<head>\n")
	b.WriteString(`<meta charset="utf-8">` + "\n")
	b.WriteString(`<meta name="viewport" content="width=device-width, initial-scale=1">` + "\n")
	b.WriteString("<title>Page not found. AlphaFlux</title>\n")
	b.WriteString(`<meta name="robots" content="noindex">` + "\n")
	b.WriteString(`<link rel="icon" href="/assets/favicon.svg" type="image/svg+xml">` + "\n")
	b.WriteString(`<link rel="stylesheet" href="/assets/site.css?v=` + assetVersion + `">` + "\n")
	b.WriteString(themeScript())
	b.WriteString("</head>\n<body>\n")
	b.WriteString(skipLink())
	b.WriteString(`<div class="shell">` + "\n")
	b.WriteString(rail(""))
	b.WriteString(`<div class="main">` + "\n")
	b.WriteString(topbar(nil, "Not found"))
	b.WriteString(`<main id="main">` + "\n")
	b.WriteString(`<section class="section"><div class="wrap"><div class="split">` + "\n")
	b.WriteString(`<div class="stack-lg">` + "\n")
	b.WriteString(`<h1>That page is not here</h1>` + "\n")
	b.WriteString(`<p class="standfirst">The address may have been mistyped, or the page may have moved. ` +
		`Every module is listed in the rail, and the site map has the full list.</p>` + "\n")
	b.WriteString(`<div class="btn-row"><a class="btn btn-primary" href="/">Go to the overview</a>` +
		`<a class="btn btn-secondary" href="/sitemap">Open the site map</a></div>` + "\n")
	b.WriteString(`</div>` + "\n")
	b.WriteString(`<div>` + moduleMiniList(8) + `</div>` + "\n")
	b.WriteString(`</div></div></section>` + "\n")
	b.WriteString(`</main>` + "\n")
	b.WriteString(footer())
	b.WriteString("</div>\n</div>\n")
	b.WriteString(`</body>\n</html>\n`)
	return b.String()
}

/* ---------- head ---------- */

func metaBlock(p *Page, path string) string {
	canonical := siteURL + "/" + path
	title := p.Title
	desc := p.Desc
	if title == "" {
		title = p.H1 + " | " + siteName
	}
	var b strings.Builder

	b.WriteString(`<meta charset="utf-8">` + "\n")
	b.WriteString(`<meta name="viewport" content="width=device-width, initial-scale=1">` + "\n")
	b.WriteString("<title>" + esc(title) + "</title>\n")
	b.WriteString(`<meta name="description" content="` + esc(desc) + `">` + "\n")
	if p.NoIndex {
		b.WriteString(`<meta name="robots" content="noindex,follow">` + "\n")
	} else {
		b.WriteString(`<meta name="robots" content="index,follow,max-snippet:-1,max-image-preview:large">` + "\n")
	}
	b.WriteString(`<link rel="canonical" href="` + canonical + `">` + "\n")

	// Link previews
	b.WriteString(`<meta property="og:site_name" content="` + siteName + `">` + "\n")
	b.WriteString(`<meta property="og:type" content="website">` + "\n")
	b.WriteString(`<meta property="og:title" content="` + esc(title) + `">` + "\n")
	b.WriteString(`<meta property="og:description" content="` + esc(desc) + `">` + "\n")
	b.WriteString(`<meta property="og:url" content="` + canonical + `">` + "\n")
	b.WriteString(`<meta property="og:image" content="` + siteURL + `/assets/og/` + ogName(p.Slug) + `">` + "\n")
	b.WriteString(`<meta property="og:image:width" content="1200">` + "\n")
	b.WriteString(`<meta property="og:image:height" content="630">` + "\n")
	b.WriteString(`<meta name="twitter:card" content="summary_large_image">` + "\n")

	b.WriteString(`<meta name="theme-color" content="#ffffff" media="(prefers-color-scheme: light)">` + "\n")
	b.WriteString(`<meta name="theme-color" content="#0a1220" media="(prefers-color-scheme: dark)">` + "\n")
	b.WriteString(`<link rel="icon" href="/assets/favicon.svg" type="image/svg+xml">` + "\n")
	b.WriteString(`<link rel="apple-touch-icon" href="/assets/apple-touch-icon.png">` + "\n")
	b.WriteString(`<link rel="alternate" type="text/plain" href="/llms.txt" title="llms.txt">` + "\n")

	// Fonts, self hosted. Only the two faces used above the fold are preloaded.
	b.WriteString(`<link rel="preload" href="/assets/fonts/archivo-latin.woff2" as="font" type="font/woff2" crossorigin>` + "\n")
	b.WriteString(`<link rel="preload" href="/assets/fonts/plexmono-400-latin.woff2" as="font" type="font/woff2" crossorigin>` + "\n")
	b.WriteString(`<link rel="stylesheet" href="/assets/fonts.css?v=` + assetVersion + `">` + "\n")
	b.WriteString(`<link rel="stylesheet" href="/assets/site.css?v=` + assetVersion + `">` + "\n")
	b.WriteString(themeScript())
	b.WriteString(jsonLD(p, canonical))

	return b.String()
}

func ogName(slug string) string {
	if slug == "index" {
		return "home.png"
	}
	return slug + ".png"
}

// themeScript runs before paint so the stored or preferred theme is applied
// without a flash. Nothing else is synchronous in the head.
func themeScript() string {
	return `<script>
(function(){try{var s=localStorage.getItem("af-theme");var d=window.matchMedia("(prefers-color-scheme: light)").matches?"light":"dark";document.documentElement.setAttribute("data-theme",s||d);}catch(e){}})();
</script>` + "\n"
}

func jsonLD(p *Page, canonical string) string {
	var graphs []map[string]any

	org := map[string]any{
		"@type":       "Organization",
		"@id":         siteURL + "/#org",
		"name":        siteName,
		"url":         siteURL + "/",
		"description": "A self-hosted, modular operations platform for service businesses, driver accessible by AI agents.",
		"logo": map[string]any{
			"@type": "ImageObject",
			"url":   siteURL + "/assets/apple-touch-icon.png",
		},
	}
	graphs = append(graphs, org)

	if p.Slug == "index" {
		graphs = append(graphs, map[string]any{
			"@type":     "WebSite",
			"@id":       siteURL + "/#website",
			"url":       siteURL + "/",
			"name":      siteName,
			"publisher": map[string]any{"@id": siteURL + "/#org"},
		})
	}

	// The breadcrumb trail in structured data matches the visible trail, and
	// like the visible one it is omitted on the root page.
	slug := p.Slug
	if slug != "index" {
		pos := 1
		items := []map[string]any{{
			"@type":    "ListItem",
			"position": pos,
			"name":     "Overview",
			"item":     siteURL + "/",
		}}
		if g := group(slug); g != "" {
			pos++
			items = append(items, map[string]any{
				"@type":    "ListItem",
				"position": pos,
				"name":     g,
			})
		}
		pos++
		items = append(items, map[string]any{
			"@type":    "ListItem",
			"position": pos,
			"name":     p.Crumb,
			"item":     canonical,
		})
		graphs = append(graphs, map[string]any{
			"@type":           "BreadcrumbList",
			"itemListElement": items,
		})
	}

	if len(p.FAQ) > 0 {
		var qs []map[string]any
		for _, q := range p.FAQ {
			qs = append(qs, map[string]any{
				"@type": "Question",
				"name":  q.Q,
				"acceptedAnswer": map[string]any{
					"@type": "Answer",
					"text":  q.A,
				},
			})
		}
		graphs = append(graphs, map[string]any{
			"@type":      "FAQPage",
			"mainEntity": qs,
		})
	}

	if slug == "modules" {
		var l []map[string]any
		for i, s := range order {
			l = append(l, map[string]any{
				"@type":    "ListItem",
				"position": i + 1,
				"name":     navItem(s).Label,
				"url":      siteURL + "/" + urlOf(s),
			})
		}
		graphs = append(graphs, map[string]any{
			"@type":           "ItemList",
			"name":            "AlphaFlux modules",
			"itemListElement": l,
		})
	}

	if len(p.Plans) > 0 {
		var offers []map[string]any
		for _, pl := range p.Plans {
			o := map[string]any{
				"@type":       "Offer",
				"name":        pl.Name,
				"description": pl.For,
				"price":       strings.TrimPrefix(strings.ReplaceAll(pl.Price, ",", ""), "$"),
				"priceCurrency": "USD",
				"url":         canonical,
			}
			if pl.Price == "Custom" {
				o["price"] = "0"
				o["description"] = pl.For + " Pricing on application."
			}
			offers = append(offers, o)
		}
		graphs = append(graphs, map[string]any{
			"@type":               "SoftwareApplication",
			"name":                siteName,
			"applicationCategory": "BusinessApplication",
			"operatingSystem":     "Web",
			"description":         p.Desc,
			"url":                 siteURL + "/",
			"offers":              offers,
		})
	}

	doc := map[string]any{
		"@context": "https://schema.org",
		"@graph":   graphs,
	}
	out, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return ""
	}
	return "<script type=\"application/ld+json\">\n" + string(out) + "\n</script>\n"
}

/* ---------- rail ---------- */

func rail(active string) string {
	var b strings.Builder
	b.WriteString(`<aside class="rail" aria-label="Site sections">` + "\n")
	b.WriteString(`<div class="rail-inner">` + "\n")
	b.WriteString(`<a class="brand" href="/">` + brandMark() + `<span class="brand-name">AlphaFlux</span></a>` + "\n")
	b.WriteString(`<nav class="rail-nav" aria-label="All sections">` + "\n")
	for _, g := range Site {
		b.WriteString(`<div class="rail-group">` + "\n")
		b.WriteString(`<h2 class="rail-group-title">` + esc(g.Title) + `</h2>` + "\n")
		b.WriteString(`<ul class="rail-list">` + "\n")
		for _, it := range g.Items {
			cur := ""
			if it.Slug == active {
				cur = ` aria-current="page"`
			}
			b.WriteString(`<li><a class="rail-link" href="/` + urlOf(it.Slug) + `"` + cur + `>` +
				icon(it.Icon) + `<span>` + esc(it.Label) + `</span></a></li>` + "\n")
		}
		b.WriteString("</ul>\n</div>\n")
	}
	b.WriteString("</nav>\n")
	b.WriteString(`<div class="rail-foot">` + "\n")
	b.WriteString(`<a class="btn btn-primary" href="/login">Sign in</a>` + "\n")
	b.WriteString(`<a class="btn btn-secondary" href="/pricing">Start free</a>` + "\n")
	b.WriteString("</div>\n")
	b.WriteString("</div>\n</aside>\n")
	return b.String()
}

func brandMark() string {
	return `<svg class="brand-mark" viewBox="0 0 24 24" aria-hidden="true" focusable="false" fill="none">` +
		`<circle cx="12" cy="12" r="9.25" stroke="currentColor" stroke-width="1.5" opacity=".35"/>` +
		`<circle cx="12" cy="12" r="5" stroke="currentColor" stroke-width="1.5" opacity=".6"/>` +
		`<path d="M12 12 19 5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/>` +
		`<circle cx="12" cy="12" r="1.75" fill="currentColor"/>` +
		`</svg>`
}

func drawer() string {
	var b strings.Builder
	b.WriteString(`<div class="drawer" data-drawer data-open="false" aria-hidden="true" role="dialog" aria-label="Site sections">` + "\n")
	b.WriteString(`<div class="drawer-scrim" data-drawer-scrim></div>` + "\n")
	b.WriteString(`<div class="drawer-panel">` + "\n")
	b.WriteString(`<div class="drawer-head"><a class="brand" href="/">` + brandMark() +
		`<span class="brand-name">AlphaFlux</span></a>` +
		`<button class="nav-toggle" type="button" data-drawer-close aria-label="Close menu">` +
		`<svg aria-hidden="true"><use href="/assets/icons.svg#i-close"></use></svg></button></div>` + "\n")
	for _, g := range Site {
		b.WriteString(`<div class="rail-group"><h2 class="rail-group-title">` + esc(g.Title) + `</h2><ul class="rail-list">` + "\n")
		for _, it := range g.Items {
			b.WriteString(`<li><a class="rail-link" href="/` + urlOf(it.Slug) + `">` + icon(it.Icon) +
				`<span>` + esc(it.Label) + `</span></a></li>` + "\n")
		}
		b.WriteString("</ul></div>\n")
	}
	b.WriteString("</div>\n</div>\n")
	return b.String()
}

/* ---------- top bar and breadcrumbs ---------- */

func topbar(p *Page, crumb string) string {
	var b strings.Builder
	b.WriteString(`<header class="topbar">` + "\n")
	b.WriteString(`<button class="nav-toggle" type="button" data-drawer-open aria-expanded="false" aria-label="Open menu">` +
		`<svg aria-hidden="true"><use href="/assets/icons.svg#i-menu"></use></svg></button>` + "\n")
	b.WriteString(`<a class="topbar-brand" href="/">` + brandMark() + `</a>` + "\n")

	// The overview page is the root of the trail, so a trail here would read
	// "Overview / Platform / Overview". A breadcrumb that names the page you are
	// already on is worse than none, so the root gets none.
	if p != nil && p.Slug != "index" {
		b.WriteString(`<nav class="crumbs" aria-label="Breadcrumb"><ol>` + "\n")
		b.WriteString(`<li><a href="/">Overview</a></li>` + "\n")
		if g := group(p.Slug); g != "" {
			b.WriteString(`<li><span class="sep" aria-hidden="true">/</span><span>` + esc(g) + `</span></li>` + "\n")
		}
		b.WriteString(`<li><span class="sep" aria-hidden="true">/</span><span aria-current="page">` + esc(crumb) + `</span></li>` + "\n")
		b.WriteString("</ol></nav>\n")
	} else if p == nil {
		b.WriteString(`<span class="small muted">` + esc(crumb) + `</span>` + "\n")
	}

	b.WriteString(`<span class="topbar-spacer"></span>` + "\n")
	b.WriteString(`<div class="topbar-actions">` + "\n")
	b.WriteString(`<button class="theme-toggle" type="button" data-theme-toggle hidden aria-label="Switch colour theme">` +
		`<svg class="i-sun" aria-hidden="true"><use href="/assets/icons.svg#i-sun"></use></svg>` +
		`<svg class="i-moon" aria-hidden="true"><use href="/assets/icons.svg#i-moon"></use></svg>` +
		`</button>` + "\n")
	b.WriteString(`<a class="btn btn-ghost btn-sm" href="/login">Sign in</a>` + "\n")
	b.WriteString(`<a class="btn btn-primary btn-sm" href="/pricing">Start free</a>` + "\n")
	b.WriteString("</div>\n</header>\n")
	return b.String()
}

func skipLink() string {
	return `<a class="skip-link" href="#main">Skip to content</a>` + "\n"
}

func pageHead(p *Page, crumb string) string {
	var b strings.Builder
	b.WriteString(`<div class="wrap">` + "\n")
	b.WriteString(`<header class="page-head">` + "\n")
	if len(p.Chips) > 0 {
		b.WriteString(`<div class="eyebrow-line">` + "\n")
		for _, c := range p.Chips {
			b.WriteString(`<span class="chip"><span class="chip-dot"></span>` + esc(c) + `</span>` + "\n")
		}
		b.WriteString("</div>\n")
	}
	b.WriteString("<h1>" + esc(p.H1) + "</h1>\n")
	if p.Stand != "" {
		b.WriteString(`<p class="standfirst">` + esc(p.Stand) + `</p>` + "\n")
	}
	b.WriteString("</header>\n</div>\n")
	return b.String()
}

/* ---------- sections ---------- */

func renderSection(p *Page, s Section, idx int) string {
	id := s.ID
	if id == "" {
		id = s.Kind
	}
	// Deliberately no scroll reveal. Hiding every section until it enters the
	// viewport means the page is blank to anything that does not scroll it
	// (a crawler, a full page capture, a reader whose script failed), and a
	// fade-up on every section is scattered motion rather than direction. The
	// orchestrated moments are the route line and the order strip, both of
	// which show something the page is actually claiming.
	reveal := ""
	_ = reveal
	_ = idx
	var b strings.Builder

	switch s.Kind {
	case "hero":
		b.WriteString(`<section class="section section-flush" id="top"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(`<div class="split">` + "\n")
		b.WriteString(`<div class="stack-lg">` + "\n")
		if len(p.Chips) > 0 {
			b.WriteString(`<div class="eyebrow-line">`)
			for _, c := range p.Chips {
				b.WriteString(`<span class="chip"><span class="chip-dot"></span>` + esc(c) + `</span>`)
			}
			b.WriteString("</div>\n")
		}
		b.WriteString(`<h1 class="display">` + esc(s.Title) + `</h1>` + "\n")
		if s.Lede != "" {
			b.WriteString(`<p class="standfirst">` + esc(s.Lede) + `</p>` + "\n")
		}
		if len(s.Body) > 0 {
			b.WriteString(`<div class="prose">` + "\n")
			for _, para := range s.Body {
				b.WriteString("<p>" + esc(para) + "</p>\n")
			}
			b.WriteString("</div>\n")
		}
		b.WriteString(`<div class="btn-row">` + "\n")
		if s.More != nil {
			b.WriteString(`<a class="btn btn-primary btn-lg" href="` + esc(s.More.Href) + `">` + esc(s.More.Title) + `</a>` + "\n")
		}
		if s.IDRef != "" {
			b.WriteString(s.IDRef + "\n")
		}
		b.WriteString("</div>\n</div>\n")
		b.WriteString(`<div class="stack-lg">` + "\n")
		b.WriteString(orderStrip(Section{Items: s.Items, Note: s.Note}))
		if s.Text != "" {
			b.WriteString(`<p class="small muted">` + esc(s.Text) + `</p>` + "\n")
		}
		b.WriteString("</div>\n")
		b.WriteString("</div></div></section>\n")

	case "modules":
		b.WriteString(`<section class="section" id="modules"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(moduleGrid())
		b.WriteString("</div></section>\n")

	case "sitemap":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(`<div class="sitemap-cols">` + "\n")
		for _, g := range Site {
			b.WriteString("<div>\n<h2>" + esc(g.Title) + "</h2>\n<ul>\n")
			for _, it := range g.Items {
				b.WriteString(`<li><a href="/` + urlOf(it.Slug) + `">` + esc(it.Label) + `</a></li>` + "\n")
			}
			b.WriteString("</ul>\n</div>\n")
		}
		b.WriteString("</div>\n</div></section>\n")

	case "statement":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(`<div class="statement"><p>` + esc(s.Text) + `</p></div>` + "\n")
		b.WriteString("</div></section>\n")

	case "figures":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(`<div class="figures">` + "\n")
		for _, it := range s.Items {
			b.WriteString(`<div class="figure">` + "\n")
			b.WriteString(`<div class="value numeric">` + esc(it.Title))
			if it.Kicker != "" {
				b.WriteString(`<span class="unit">` + esc(it.Kicker) + `</span>`)
			}
			b.WriteString("</div>\n")
			b.WriteString(`<p class="caption">` + esc(it.Text) + `</p>` + "\n")
			b.WriteString("</div>\n")
		}
		b.WriteString("</div>\n</div></section>\n")

	case "order":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(orderStrip(s))
		b.WriteString("</div></section>\n")

	case "prose":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(`<div class="prose stack-lg">` + "\n")
		for _, para := range s.Body {
			b.WriteString("<p>" + esc(para) + "</p>\n")
		}
		b.WriteString("</div>\n")
		if s.More != nil {
			b.WriteString(`<p class="stack" style="margin-block-start:var(--space-6)">` + linkMore(*s.More) + `</p>` + "\n")
		}
		b.WriteString("</div></section>\n")

	case "split":
		flip := ""
		if s.Flip {
			flip = " split-flip"
		}
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap"><div class="split` + flip + `">` + "\n")
		b.WriteString(`<div class="stack-lg">` + "\n")
		if s.Title != "" {
			b.WriteString("<h2>" + esc(s.Title) + "</h2>\n")
		}
		if s.Lede != "" {
			b.WriteString(`<p class="standfirst">` + esc(s.Lede) + `</p>` + "\n")
		}
		if len(s.Body) > 0 {
			b.WriteString(`<div class="prose">` + "\n")
			for _, para := range s.Body {
				b.WriteString("<p>" + esc(para) + "</p>\n")
			}
			b.WriteString("</div>\n")
		}
		if s.More != nil {
			b.WriteString(linkMore(*s.More) + "\n")
		}
		b.WriteString("</div>\n")
		b.WriteString(`<div class="stack-lg">` + "\n")
		b.WriteString(featureItems(s.Items))
		b.WriteString("</div>\n")
		b.WriteString("</div></div></section>\n")

	case "features":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(featureItems(s.Items))
		b.WriteString("</div></section>\n")

	case "list":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(`<dl class="spec-list">` + "\n")
		for _, it := range s.Items {
			b.WriteString("<li><dt>" + esc(it.Title) + "</dt><dd>" + esc(it.Text))
			if len(it.Items) > 0 {
				b.WriteString("<ul>")
				for _, sub := range it.Items {
					b.WriteString("<li>" + esc(sub) + "</li>")
				}
				b.WriteString("</ul>")
			}
			if it.Note != "" {
				b.WriteString(`<p class="item-note" style="margin-block-start:var(--space-2)">` + esc(it.Note) + "</p>")
			}
			b.WriteString("</dd></li>\n")
		}
		b.WriteString("</dl>\n</div></section>\n")

	case "steps":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(`<ol class="steps">` + "\n")
		for _, it := range s.Items {
			b.WriteString("<li><h3>" + esc(it.Title) + "</h3><p>" + esc(it.Text) + "</p></li>\n")
		}
		b.WriteString("</ol>\n</div></section>\n")

	case "table":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		// A wide table scrolls sideways on a phone rather than collapsing into
		// cards, because cards destroy the comparison the table exists for. A
		// scroll container has to be reachable and operable by keyboard, so it
		// takes a tab stop and a label rather than being a mouse-only region.
		label := s.Title
		if label == "" {
			label = "Comparison table"
		}
		b.WriteString(`<div class="table-scroll" role="region" aria-label="` + esc(label) +
			`, scrolls horizontally" tabindex="0"><table class="data">` + "\n")
		if s.Title != "" {
			b.WriteString("<caption>" + esc(s.Lede) + "</caption>\n")
		}
		if len(s.Cols) > 0 {
			b.WriteString("<thead><tr>")
			for _, c := range s.Cols {
				b.WriteString("<th scope=\"col\">" + esc(c) + "</th>")
			}
			b.WriteString("</tr></thead>\n")
		}
		b.WriteString("<tbody>\n")
		for _, r := range s.Rows {
			b.WriteString("<tr>")
			for i, c := range r {
				if i == 0 {
					b.WriteString("<th scope=\"row\">" + esc(c) + "</th>")
				} else {
					b.WriteString("<td>" + cellHTML(c) + "</td>")
				}
			}
			b.WriteString("</tr>\n")
		}
		b.WriteString("</tbody></table></div>\n")
		if s.Note != "" {
			b.WriteString(`<p class="small muted" style="margin-block-start:var(--space-4)">` + esc(s.Note) + "</p>" + "\n")
		}
		b.WriteString("</div></section>\n")

	case "code":
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString(`<pre class="code" id="` + esc(s.IDRef) + `"><code>` + esc(s.Text) + `</code></pre>` + "\n")
		if s.More != nil {
			b.WriteString(`<div style="margin-block-start:var(--space-4)">`)
			if s.IDRef != "" {
				b.WriteString(`<button class="btn btn-secondary btn-sm" type="button" data-copy="` + esc(s.IDRef) + `">Copy</button>`)
			}
			b.WriteString(" " + linkMore(*s.More) + "</div>\n")
		}
		b.WriteString("</div></section>\n")

	default:
		b.WriteString(`<section class="section"` + reveal + `><div class="wrap">` + "\n")
		b.WriteString(sectionHead(s))
		b.WriteString("</div></section>\n")
	}
	return b.String()
}

func cellHTML(c string) string {
	switch c {
	case "yes":
		return `<span class="yes">Yes</span>`
	case "no":
		return `<span class="no">Not included</span>`
	case "opt":
		return `<span class="muted">Optional</span>`
	}
	if strings.HasPrefix(c, "!") {
		return `<span class="cell-strong">` + esc(strings.TrimPrefix(c, "!")) + `</span>`
	}
	return esc(c)
}

func sectionHead(s Section) string {
	if s.Title == "" && s.Lede == "" {
		return ""
	}
	var b strings.Builder
	b.WriteString(`<div class="section-head">` + "\n")
	if s.Title != "" {
		b.WriteString("<h2>" + esc(s.Title) + "</h2>\n")
	}
	if s.Lede != "" {
		b.WriteString(`<p class="lede">` + esc(s.Lede) + `</p>` + "\n")
	}
	b.WriteString("</div>\n")
	return b.String()
}

func featureItems(items []Item) string {
	if len(items) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(`<ul class="grid grid-3" style="list-style:none">` + "\n")
	for _, it := range items {
		b.WriteString(`<li class="item">` + "\n")
		if it.Icon != "" {
			b.WriteString(`<span class="item-icon">` + icon(it.Icon) + `</span>` + "\n")
		}
		b.WriteString("<h3>" + esc(it.Title) + "</h3>\n")
		b.WriteString("<p>" + esc(it.Text) + "</p>\n")
		if it.Note != "" {
			b.WriteString(`<p class="item-note">` + esc(it.Note) + `</p>` + "\n")
		}
		if len(it.Items) > 0 {
			b.WriteString(`<ul class="plan-features">`)
			for _, sub := range it.Items {
				b.WriteString("<li>" + esc(sub) + "</li>")
			}
			b.WriteString("</ul>\n")
		}
		if it.Link != "" && it.Href != "" {
			b.WriteString(`<a class="link-more" href="` + esc(it.Href) + `">` + esc(it.Link) +
				icon("i-arrow") + `</a>` + "\n")
		}
		b.WriteString("</li>\n")
	}
	b.WriteString("</ul>\n")
	return b.String()
}

func linkMore(it Item) string {
	if it.Href == "" {
		return ""
	}
	return `<a class="link-more" href="` + esc(it.Href) + `">` + esc(it.Title) + icon("i-arrow") + `</a>`
}

/* The work-order strip: a real record moving through its stages. It is reused
   on the overview and the operations pages because it explains the product
   better than a screenshot would. */
func orderStrip(s Section) string {
	stages := s.Items
	if len(stages) == 0 {
		stages = []Item{
			{Title: "Scheduled", Note: "06:40"},
			{Title: "Dispatched", Note: "07:05"},
			{Title: "On site", Note: "07:52"},
			{Title: "Complete", Note: "08:31"},
		}
	}
	var b strings.Builder
	b.WriteString(`<div class="order-strip" data-order-strip>` + "\n")
	b.WriteString(`<div class="order-head">` + "\n")
	b.WriteString(`<span class="order-id">JOB-24188</span>` + "\n")
	b.WriteString(`<span class="chip chip-ok"><span class="chip-dot"></span>Capacity cleared</span>` + "\n")
	b.WriteString(`<span class="order-meta">Quarterly treatment. Technician, truck 4. Zone, North Spokane.</span>` + "\n")
	b.WriteString("</div>\n")
	b.WriteString(`<ol class="order-track">` + "\n")
	for i, st := range stages {
		state := "todo"
		if i == 0 {
			state = "now"
		}
		b.WriteString(`<li data-state="` + state + `">` + "\n")
		b.WriteString(`<span class="stage">` + esc(st.Title) + `</span>` + "\n")
		if st.Note != "" {
			b.WriteString(`<span class="at">` + esc(st.Note) + `</span>` + "\n")
		}
		b.WriteString("</li>\n")
	}
	b.WriteString("</ol>\n</div>\n")
	if s.Note != "" {
		b.WriteString(`<p class="small muted" style="margin-block-start:var(--space-4)">` + esc(s.Note) + `</p>` + "\n")
	}
	return b.String()
}

/* ---------- module index ---------- */

func moduleGrid() string {
	var b strings.Builder
	b.WriteString(`<ul class="module-grid">` + "\n")
	for _, slug := range order {
		if group(slug) != "Modules" {
			continue
		}
		it := navItem(slug)
		b.WriteString(`<li><a class="module-card" href="/` + urlOf(slug) + `">` + "\n")
		b.WriteString(`<span class="item-icon">` + icon(it.Icon) + `</span>` + "\n")
		b.WriteString("<h3>" + esc(it.Label) + "</h3>\n")
		b.WriteString("<p>" + esc(it.Blurb) + "</p>\n")
		b.WriteString("</a></li>\n")
	}
	b.WriteString("</ul>\n")
	return b.String()
}

func moduleMiniList(n int) string {
	var b strings.Builder
	b.WriteString(`<ul class="spec-list">` + "\n")
	count := 0
	for _, slug := range order {
		if count >= n {
			break
		}
		if group(slug) != "Modules" {
			continue
		}
		it := navItem(slug)
		b.WriteString(`<li><dt><a href="/` + urlOf(slug) + `">` + esc(it.Label) + `</a></dt>` +
			`<dd>` + esc(it.Blurb) + `</dd></li>` + "\n")
		count++
	}
	b.WriteString("</ul>\n")
	return b.String()
}

/* ---------- plans ---------- */

func renderPlans(p *Page) string {
	cols := "plans"
	if len(p.Plans) == 4 {
		cols = "plans plans-4"
	}
	var b strings.Builder
	b.WriteString(`<section class="section" id="pricing">` + "\n")
	b.WriteString(`<div class="wrap">` + "\n")
	b.WriteString(`<div class="` + cols + `">` + "\n")
	for _, pl := range p.Plans {
		class := "plan"
		if pl.Featured {
			class += " plan-featured"
		}
		b.WriteString(`<div class="` + class + `">` + "\n")
		if pl.Badge != "" {
			b.WriteString(`<span class="plan-badge">` + esc(pl.Badge) + `</span>` + "\n")
		}
		b.WriteString(`<h2 class="plan-name">` + esc(pl.Name) + `</h2>` + "\n")
		b.WriteString(`<div class="plan-price numeric">` + esc(pl.Price))
		if pl.Per != "" {
			b.WriteString(`<span class="per">` + esc(pl.Per) + `</span>`)
		}
		b.WriteString("</div>\n")
		b.WriteString(`<p class="plan-for">` + esc(pl.For) + `</p>` + "\n")
		if len(pl.Limits) > 0 {
			b.WriteString(`<dl class="plan-limits">` + "\n")
			for _, l := range pl.Limits {
				b.WriteString(`<div style="display:flex;justify-content:space-between;gap:var(--space-4)">` +
					`<dt class="k">` + esc(l.Key) + `</dt><dd class="v">` + esc(l.Val) + `</dd></div>` + "\n")
			}
			b.WriteString("</dl>\n")
		}
		if len(pl.Features) > 0 {
			b.WriteString(`<ul class="plan-features">` + "\n")
			for _, f := range pl.Features {
				b.WriteString("<li>" + esc(f) + "</li>\n")
			}
			b.WriteString("</ul>\n")
		}
		btn := "btn-secondary"
		if pl.Featured {
			btn = "btn-primary"
		}
		href := pl.Href
		if href == "" {
			href = "/login"
		}
		b.WriteString(`<a class="btn ` + btn + `" href="` + esc(href) + `">` + esc(pl.CTA) + `</a>` + "\n")
		if pl.Note != "" {
			b.WriteString(`<p class="small muted" style="margin-block-start:var(--space-3)">` + esc(pl.Note) + `</p>` + "\n")
		}
		b.WriteString("</div>\n")
	}
	b.WriteString("</div>\n")
	if p.PlansNote != "" {
		b.WriteString(`<p class="small muted" style="margin-block-start:var(--space-6);max-inline-size:70ch">` +
			esc(p.PlansNote) + `</p>` + "\n")
	}
	b.WriteString("</div></section>\n")
	return b.String()
}

/* ---------- faq and cta ---------- */

func renderFAQ(p *Page) string {
	var b strings.Builder
	b.WriteString(`<section class="section" id="faq">` + "\n")
	b.WriteString(`<div class="wrap"><div class="split">` + "\n")
	b.WriteString(`<div><h2>Questions people ask before signing up</h2></div>` + "\n")
	b.WriteString(`<div class="faq">` + "\n")
	for _, q := range p.FAQ {
		b.WriteString(`<details><summary>` + esc(q.Q) + `</summary><div class="answer"><p>` + esc(q.A) + `</p></div></details>` + "\n")
	}
	b.WriteString("</div>\n</div></div></section>\n")
	return b.String()
}

func ctaBand(p *Page) string {
	c := p.CTA
	if c == nil {
		c = &CTA{
			Title: "Start on the free plan",
			Text:  "Twenty customers, one location, every core module, free permanently. Move up when the work outgrows it.",
			Primary: &Item{Title: "Create a free account", Href: "/login"},
			Second:  &Item{Title: "See what each plan includes", Href: "/pricing"},
		}
	}
	var b strings.Builder
	b.WriteString(`<section class="band"><div class="wrap"><div class="split">` + "\n")
	b.WriteString(`<div class="stack-lg">` + "\n")
	b.WriteString("<h2>" + esc(c.Title) + "</h2>\n")
	b.WriteString("<p>" + esc(c.Text) + "</p>\n")
	b.WriteString("</div>\n<div class=\"btn-row\" style=\"align-self:center\">\n")
	if c.Primary != nil {
		b.WriteString(`<a class="btn btn-primary btn-lg" href="` + esc(c.Primary.Href) + `">` + esc(c.Primary.Title) + `</a>` + "\n")
	}
	if c.Second != nil {
		b.WriteString(`<a class="btn btn-secondary btn-lg" href="` + esc(c.Second.Href) + `">` + esc(c.Second.Title) + `</a>` + "\n")
	}
	b.WriteString("</div>\n</div></div></section>\n")
	return b.String()
}

func footer() string {
	var b strings.Builder
	b.WriteString(`<footer class="footer"><div class="wrap">` + "\n")
	b.WriteString(`<div class="footer-cols">` + "\n")
	b.WriteString(`<div><a class="brand" href="/">` + brandMark() +
		`<span class="brand-name">AlphaFlux</span></a>` +
		`<p class="small muted" style="margin-block-start:var(--space-4);max-inline-size:34ch">` +
		`One platform for the phone, the field, the money and the marketing, driven by you or by your agents.</p></div>` + "\n")
	for _, g := range Site {
		b.WriteString("<div><h2>" + esc(g.Title) + "</h2><ul>\n")
		for _, it := range g.Items {
			b.WriteString(`<li><a href="/` + urlOf(it.Slug) + `">` + esc(it.Label) + `</a></li>` + "\n")
		}
		b.WriteString("</ul></div>\n")
	}
	b.WriteString("</div>\n")
	b.WriteString(`<div class="footer-note">` + "\n")
	b.WriteString(`<span>AlphaFlux. Built and hosted by the companies that use it.</span>` + "\n")
	// The build stamp is on every page so a deploy is visible from the outside
	// without opening the console or asking the server. The product does the
	// same thing for the same reason.
	b.WriteString(`<span class="ident">build ` + buildVer + `.1</span>` + "\n")
	b.WriteString(`<span><a href="/sitemap">Site map</a> and <a href="/llms.txt">llms.txt</a> for machines.</span>` + "\n")
	b.WriteString("</div>\n</div></footer>\n")
	return b.String()
}
