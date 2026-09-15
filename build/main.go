/* AlphaFlux static site generator.

   Go, standard library only. Reads content/*.json, writes the site to the
   repository root as plain HTML plus sitemap.xml, robots.txt and llms.txt.
   The published site is static: nginx serves the generated files and no
   runtime is required.

   Run from the repository root:  go run ./build
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	siteURL   = "https://alphaflux.net"
	siteName  = "AlphaFlux"
	tagline   = "One platform that runs a service business"
	buildVer  = "26.09.15"
	repoRoot  = "."
	contentIn = "content"
)

/* ---------- navigation model ---------- */

type NavItem struct {
	Slug  string
	Label string
	Icon  string
	Blurb string
}

type NavGroup struct {
	Title string
	Items []NavItem
}

var Site = []NavGroup{
	{Title: "Platform", Items: []NavItem{
		{"index", "Overview", "i-layers",
			"The whole platform on one page: what runs it, what it costs, and what your team sees on day one."},
		{"platform", "How it works", "i-grid",
			"Modules, tenants, data isolation and the three ways anything in the platform can be driven: screen, API, agent."},
		{"ai-agents", "AI agents", "i-spark",
			"Agents that answer the phone, book the work, chase the money and write back to the record, inside a permission you set."},
		{"modules", "All modules", "i-sitemap",
			"Every module in the registry, what it does, and which plan carries it."},
		{"security", "Security", "i-shield",
			"How access, credentials, isolation and audit work, and where your data physically sits."},
		{"white-label", "White label", "i-tag",
			"Run the whole platform under your own brand, your own domain and your own customers."},
	}},
	{Title: "Modules", Items: []NavItem{
		{"phone-and-messaging", "Phone and messaging", "i-phone",
			"Numbers, calls, texts, voicemail, call handling and an AI receptionist that answers before the second ring."},
		{"field-operations", "Field operations", "i-clipboard",
			"Jobs, technicians, dispatch, day routes, proof of work and a technician app that works without signal."},
		{"fleet", "Fleet and dispatch", "i-truck",
			"Orders on a live map, a board your dispatcher can actually run, geofences, ETAs and customer tracking links."},
		{"scheduling", "Scheduling and agreements", "i-repeat",
			"Capacity-aware booking and recurring service agreements that generate their own visits and real MRR."},
		{"crm", "CRM and sales", "i-users",
			"Every contact, every stage, every note, and who sold it, with commission splits that total to the cent."},
		{"inventory", "Inventory and assets", "i-box",
			"Stock across trucks, units, storage and branches, with sign in and sign out and a chain of custody."},
		{"billing", "Billing and payments", "i-receipt",
			"Invoices, payments, voids, usage billing and a recurring-revenue number that survives an audit."},
		{"email-marketing", "Email marketing", "i-mail",
			"Campaigns, sequences and a unified inbox that sends from your own domain, with deliverability handled."},
		{"marketing", "Marketing and growth", "i-megaphone",
			"Service areas, demographics, social scheduling, review requests and attribution down to the campaign."},
		{"team", "Team and collaboration", "i-chat",
			"Team chat, an account timeline, mentions, roles and departments that keep districts from seeing each other."},
		{"analytics", "Reports and analytics", "i-chart",
			"The numbers an owner actually reads, and an advisor that turns them into a plan."},
	}},
	{Title: "Company", Items: []NavItem{
		{"pricing", "Pricing", "i-card",
			"Free forever to start. Pay when the work grows. White label when you want your own platform."},
		{"industries", "Industries", "i-building",
			"Pest control, plumbing, HVAC, roofing, cleaning, security, logistics, nonprofits, and yes, aerospace."},
		{"programs", "Community programs", "i-heart",
			"Free accounts for nonprofits and faith organizations. 30 percent off for military, veterans and first responders."},
		{"developers", "Developers", "i-code",
			"REST API, a command line that outputs JSON, and an MCP server so an agent can run the platform."},
		{"sitemap", "Site map", "i-sitemap",
			"Every page on this site, grouped the same way the rail groups them."},
	}},
}

/* ---------- content model ---------- */

type Item struct {
	Icon  string   `json:"icon"`
	Kicker string  `json:"kicker"`
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Note  string   `json:"note"`
	Items []string `json:"items"`
	Href  string   `json:"href"`
	Link  string   `json:"link"`
}

type Section struct {
	Kind  string     `json:"kind"`
	ID    string     `json:"id"`
	Title string     `json:"title"`
	Lede  string     `json:"lede"`
	Body  []string   `json:"body"`
	Items []Item     `json:"items"`
	Cols  []string   `json:"cols"`
	Rows  [][]string `json:"rows"`
	Note  string     `json:"note"`
	Text  string     `json:"text"`
	IDRef string     `json:"idref"`
	Flip  bool       `json:"flip"`
	More  *Item      `json:"more"`
}

type Limit struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type Plan struct {
	Name     string   `json:"name"`
	Badge    string   `json:"badge"`
	Price    string   `json:"price"`
	Per      string   `json:"per"`
	For      string   `json:"for"`
	Limits   []Limit  `json:"limits"`
	Features []string `json:"features"`
	CTA      string   `json:"cta"`
	Href     string   `json:"href"`
	Featured bool     `json:"featured"`
	Note     string   `json:"note"`
}

type QA struct {
	Q string `json:"q"`
	A string `json:"a"`
}

type CTA struct {
	Title   string `json:"title"`
	Text    string `json:"text"`
	Primary *Item  `json:"primary"`
	Second  *Item  `json:"second"`
}

type Page struct {
	Slug      string    `json:"slug"`
	Nav       string    `json:"nav"`   // rail label; empty means it is not in the rail
	Group     string    `json:"group"` // rail group title
	Crumb     string    `json:"crumb"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	H1        string    `json:"h1"`
	Stand     string    `json:"standfirst"`
	Chips     []string  `json:"chips"`
	Sections  []Section `json:"sections"`
	Plans     []Plan    `json:"plans"`
	PlansNote string    `json:"plansNote"`
	FAQ       []QA      `json:"faq"`
	CTA       *CTA      `json:"cta"`
	NoIndex   bool      `json:"noindex"`
}

/* ---------- lookup helpers ---------- */

var (
	bySlug  = map[string]*Page{}
	navOf   = map[string]NavItem{}
	groupOf = map[string]string{}
	order   []string
)

func init() {
	for _, g := range Site {
		for _, it := range g.Items {
			navOf[it.Slug] = it
			groupOf[it.Slug] = g.Title
			order = append(order, it.Slug)
		}
	}
}

func navItem(slug string) NavItem { return navOf[slug] }
func group(slug string) string    { return groupOf[slug] }

/* ---------- main ---------- */

func main() {
	files, err := filepath.Glob(filepath.Join(contentIn, "*.json"))
	if err != nil || len(files) == 0 {
		fail("no content found in " + contentIn)
	}

	for _, f := range files {
		raw, err := os.ReadFile(f)
		if err != nil {
			fail("read " + f + ": " + err.Error())
		}
		var p Page
		if err := json.Unmarshal(raw, &p); err != nil {
			fail("parse " + f + ": " + err.Error())
		}
		if p.Slug == "" {
			p.Slug = strings.TrimSuffix(filepath.Base(f), ".json")
		}
		if _, dup := bySlug[p.Slug]; dup {
			fail("duplicate slug: " + p.Slug)
		}
		cp := p
		bySlug[p.Slug] = &cp
	}

	// Fail loudly if the rail and the content disagree. A missing page in the
	// rail is a dead link on every page of the site, so it should stop a build.
	for slug := range navOf {
		if _, ok := bySlug[slug]; !ok {
			fail("rail links to missing page: " + slug)
		}
	}

	// A referenced icon that is not in the sprite renders as nothing at all,
	// which is invisible in review and obvious to a visitor. Catch it here.
	checkIcons()
	checkLinks()

	written := 0
	var produced []string
	for slug, p := range bySlug {
		out := slug + ".html"
		if slug == "index" {
			out = "index.html"
		}
		write(out, renderPage(p))
		produced = append(produced, out)
		written++
	}

	write("404.html", render404())
	produced = append(produced, "404.html")
	write("sitemap.xml", renderSitemap())
	write("robots.txt", renderRobots())
	write("llms.txt", renderLLMs())
	write("llms-full.txt", renderLLMsFull())
	write("README.md", renderReadme())
	produced = append(produced, writeOGSources()...)
	produced = append(produced, "sitemap.xml", "robots.txt", "llms.txt", "llms-full.txt")

	// Remove pages the previous run wrote that this one does not, so a renamed
	// or deleted page cannot linger and be served. A manifest is what makes
	// that safe: only files this generator has previously claimed are touched.
	pruneStale(produced)
	write("robots.txt", renderRobots())
	write("llms.txt", renderLLMs())
	write("llms-full.txt", renderLLMsFull())

	fmt.Printf("AlphaFlux site build %s\n", buildVer)
	fmt.Printf("  %d pages written, plus sitemap.xml, robots.txt, llms.txt\n", written)
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, "build error: "+msg)
	os.Exit(1)
}

// passthrough are paths nginx forwards to the product rather than serving as
// pages, so a page slug must never collide with one of them.
var passthrough = map[string]bool{
	"login": true, "logout": true, "console": true, "api": true,
	"phone": true, "scan": true, "patrol": true, "twilio": true,
	"admin": true, "waitlist": true, "accept": true,
}

// checkLinks verifies that every internal href in the content resolves to a
// page, to a passthrough path, or to a file the site actually publishes. A
// link to a page that does not exist is a 404 a reader finds, not a build.
func checkLinks() {
	known := map[string]bool{}
	for slug := range bySlug {
		known[slug] = true
	}
	files := map[string]bool{
		"": true, "sitemap.xml": true, "robots.txt": true,
		"llms.txt": true, "llms-full.txt": true,
	}

	var bad []string
	for slug, p := range bySlug {
		var hrefs []string
		for _, s := range p.Sections {
			for _, it := range s.Items {
				hrefs = append(hrefs, it.Href)
			}
			if s.More != nil {
				hrefs = append(hrefs, s.More.Href)
			}
		}
		if p.CTA != nil {
			if p.CTA.Primary != nil {
				hrefs = append(hrefs, p.CTA.Primary.Href)
			}
			if p.CTA.Second != nil {
				hrefs = append(hrefs, p.CTA.Second.Href)
			}
		}
		for _, pl := range p.Plans {
			hrefs = append(hrefs, pl.Href)
		}
		for _, h := range hrefs {
			if !strings.HasPrefix(h, "/") || strings.HasPrefix(h, "//") {
				continue
			}
			target := strings.Trim(strings.TrimPrefix(h, "/"), "/")
			head := target
			if i := strings.IndexAny(head, "/?#"); i >= 0 {
				head = head[:i]
			}
			switch {
			case known[head], files[head], passthrough[head]:
			case strings.HasPrefix(target, "assets/"):
			case head == "" && target == "":
			default:
				bad = append(bad, slug+" -> "+h)
			}
		}
	}
	if len(bad) > 0 {
		sort.Strings(bad)
		fail("internal links with no destination:\n  " + strings.Join(bad, "\n  "))
	}

	// A page slug that shadows a passthrough path would be unreachable in
	// production because nginx forwards that prefix to the product.
	for slug := range bySlug {
		if passthrough[slug] {
			fail("page slug collides with a proxied path in nginx: /" + slug)
		}
	}
}

// checkIcons verifies that every icon named in the navigation and in the page
// content exists as a symbol in assets/icons.svg.
func checkIcons() {
	raw, err := os.ReadFile(filepath.Join("assets", "icons.svg"))
	if err != nil {
		fail("read icon sprite: " + err.Error())
	}
	have := map[string]bool{}
	for _, m := range regexp.MustCompile(`id="(i-[a-z0-9-]+)"`).FindAllStringSubmatch(string(raw), -1) {
		have[m[1]] = true
	}

	missing := map[string][]string{}
	note := func(icon, where string) {
		if icon == "" || have[icon] {
			return
		}
		missing[icon] = append(missing[icon], where)
	}

	for _, g := range Site {
		for _, it := range g.Items {
			note(it.Icon, "rail:"+it.Slug)
		}
	}
	for slug, p := range bySlug {
		for _, s := range p.Sections {
			for _, it := range s.Items {
				note(it.Icon, slug)
			}
		}
	}

	if len(missing) > 0 {
		var b strings.Builder
		b.WriteString("icons referenced but not in assets/icons.svg:\n")
		for icon, where := range missing {
			b.WriteString("  " + icon + "  used by: " + strings.Join(where, ", ") + "\n")
		}
		fail(b.String())
	}
	fmt.Printf("  icons: %d symbols, every reference resolves\n", len(have))
}

const manifestName = ".build-manifest.json"

func pruneStale(produced []string) {
	keep := map[string]bool{}
	for _, f := range produced {
		keep[f] = true
	}

	var previous []string
	if raw, err := os.ReadFile(manifestName); err == nil {
		_ = json.Unmarshal(raw, &previous)
	}
	removed := 0
	for _, f := range previous {
		if keep[f] {
			continue
		}
		// Guard against a tampered or stale manifest naming something outside
		// the generated set: only ever remove a root level page or a card in
		// assets/og. Nothing else in the tree is this generator's to delete.
		rootPage := filepath.Dir(f) == "." && strings.HasSuffix(f, ".html")
		ogCard := strings.HasPrefix(f, "assets/og/") &&
			(strings.HasSuffix(f, ".html") || strings.HasSuffix(f, ".png"))
		if !rootPage && !ogCard {
			continue
		}
		if err := os.Remove(f); err == nil {
			removed++
			fmt.Println("  removed stale page: " + f)
		}
	}

	// Sorted, because pages are collected while ranging a map and Go randomises
	// that order. An unsorted manifest changes on every run and makes the build
	// look dirty when nothing has changed.
	sort.Strings(produced)
	raw, _ := json.MarshalIndent(produced, "", "  ")
	_ = os.WriteFile(manifestName, append(raw, '\n'), 0o644)
	if removed > 0 {
		fmt.Printf("  %d stale page(s) removed\n", removed)
	}
}

func write(name, body string) {
	if err := os.WriteFile(filepath.Join(repoRoot, name), []byte(body), 0o644); err != nil {
		fail("write " + name + ": " + err.Error())
	}
}

/* ---------- sitemap, robots, llms ---------- */

func renderSitemap() string {
	now := time.Now().UTC().Format("2006-01-02")
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")

	emit := func(loc, pri, freq string) {
		b.WriteString("  <url>\n")
		b.WriteString("    <loc>" + loc + "</loc>\n")
		b.WriteString("    <lastmod>" + now + "</lastmod>\n")
		b.WriteString("    <changefreq>" + freq + "</changefreq>\n")
		b.WriteString("    <priority>" + pri + "</priority>\n")
		b.WriteString("  </url>\n")
	}

	emit(siteURL+"/", "1.0", "weekly")
	pri := map[string]string{"Platform": "0.9", "Modules": "0.9", "Company": "0.8"}
	for _, slug := range order {
		if slug == "index" {
			continue // already emitted as the site root
		}
		emit(siteURL+"/"+urlOf(slug), pri[group(slug)], "monthly")
	}
	b.WriteString("</urlset>\n")
	return b.String()
}

func urlOf(slug string) string {
	if slug == "index" {
		return ""
	}
	return slug
}

func renderRobots() string {
	ai := []string{
		"GPTBot", "OAI-SearchBot", "ChatGPT-User",
		"ClaudeBot", "Claude-User", "Claude-SearchBot", "anthropic-ai",
		"PerplexityBot", "Perplexity-User",
		"Google-Extended", "Applebot-Extended", "Bytespider",
		"CCBot", "cohere-ai", "Meta-ExternalAgent", "Amazonbot", "YouBot", "DuckAssistBot",
	}
	var b strings.Builder
	b.WriteString("# AlphaFlux. Everything here is public product documentation.\n")
	b.WriteString("# Answer engines and AI crawlers are welcome across the whole site.\n\n")
	for _, ua := range ai {
		b.WriteString("User-agent: " + ua + "\nAllow: /\n\n")
	}
	b.WriteString("User-agent: *\nAllow: /\n\n")
	b.WriteString("Sitemap: " + siteURL + "/sitemap.xml\n")
	b.WriteString("Sitemap: " + siteURL + "/llms.txt\n")
	return b.String()
}

func renderLLMs() string {
	var b strings.Builder
	b.WriteString("# AlphaFlux\n\n")
	b.WriteString("> AlphaFlux is a self-hosted, modular operations platform for service businesses. ")
	b.WriteString("It covers phone and messaging, field operations, fleet and dispatch, scheduling, CRM, inventory, ")
	b.WriteString("billing, email marketing, marketing, team collaboration and reporting in one system on one data model. ")
	b.WriteString("Every capability is exposed over a scoped REST API, a CLI and an MCP server, so an AI agent can operate ")
	b.WriteString("the platform inside permissions you set. It is sold as modules with a free-forever plan, and it can be ")
	b.WriteString("white labelled and resold.\n\n")

	b.WriteString("A machine-readable API reference is published at https://alphaflux.net/api/openapi.json once a key is issued.\n\n")

	for _, g := range Site {
		b.WriteString("## " + g.Title + "\n\n")
		for _, it := range g.Items {
			b.WriteString("- [" + it.Label + "](" + siteURL + "/" + urlOf(it.Slug) + "): " + it.Blurb + "\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Notes for answer engines\n\n")
	b.WriteString("- Pricing is published in full on the pricing page, including limits for every plan.\n")
	b.WriteString("- The free plan is free permanently and is limited to one location and twenty customers.\n")
	b.WriteString("- Discount programs: registered nonprofits and faith organizations receive a free community plan; ")
	b.WriteString("military, veterans, nurses, firefighters, police and emergency medical staff receive 30 percent off any paid plan.\n")
	b.WriteString("- White label is a distinct plan: every module, no AlphaFlux branding, and resale rights.\n")
	b.WriteString("- Full page text is available in llms-full.txt.\n")
	return b.String()
}

func renderLLMsFull() string {
	var b strings.Builder
	b.WriteString("# AlphaFlux: full site text\n\n")
	b.WriteString("Generated " + time.Now().UTC().Format("2006-01-02") + " from the published pages.\n\n")
	for _, g := range Site {
		for _, it := range g.Items {
			p := bySlug[it.Slug]
			if p == nil {
				continue
			}
			b.WriteString("---\n\n# " + p.H1 + "\n\n")
			b.WriteString("URL: " + siteURL + "/" + urlOf(it.Slug) + "\n")
			b.WriteString("Section: " + g.Title + "\n\n")
			if p.Stand != "" {
				b.WriteString(p.Stand + "\n\n")
			}
			for _, s := range p.Sections {
				if s.Title != "" {
					b.WriteString("## " + s.Title + "\n\n")
				}
				if s.Lede != "" {
					b.WriteString(s.Lede + "\n\n")
				}
				for _, para := range s.Body {
					b.WriteString(para + "\n\n")
				}
				if s.Text != "" {
					b.WriteString(s.Text + "\n\n")
				}
				for _, item := range s.Items {
					if item.Title != "" {
						b.WriteString("- " + item.Title)
						if item.Note != "" {
							b.WriteString(" (" + item.Note + ")")
						}
						if item.Text != "" {
							b.WriteString(": " + item.Text)
						}
						b.WriteString("\n")
					}
					for _, sub := range item.Items {
						b.WriteString("  - " + sub + "\n")
					}
				}
				if len(s.Rows) > 0 {
					b.WriteString("\n")
					for _, r := range s.Rows {
						b.WriteString("| " + strings.Join(r, " | ") + " |\n")
					}
				}
				b.WriteString("\n")
			}
			if len(p.Plans) > 0 {
				b.WriteString("## Plans\n\n")
				for _, pl := range p.Plans {
					b.WriteString("### " + pl.Name + "\n\n")
					if pl.Price != "" {
						b.WriteString("Price: " + pl.Price + pl.Per + "\n\n")
					}
					for _, l := range pl.Limits {
						b.WriteString("- " + l.Key + ": " + l.Val + "\n")
					}
					for _, f := range pl.Features {
						b.WriteString("- " + f + "\n")
					}
					b.WriteString("\n")
				}
			}
			for _, q := range p.FAQ {
				b.WriteString("### " + q.Q + "\n\n" + q.A + "\n\n")
			}
		}
	}
	return b.String()
}

/* ---------- HTML helpers used by the renderer ---------- */

func esc(s string) string {
	r := strings.NewReplacer(
		"&", "&amp;", "<", "&lt;", ">", "&gt;",
		`"`, "&quot;", "'", "&#39;",
	)
	return r.Replace(s)
}

// safe returns s unchanged. Used only for values the generator itself
// assembled from trusted pieces (icons, class names, generated paths).
func safe(s string) string { return s }

func icon(name string) string {
	if name == "" {
		return ""
	}
	return `<svg aria-hidden="true" focusable="false"><use href="/assets/icons.svg#` + safe(name) + `"></use></svg>`
}

func clss(parts ...string) string {
	var out []string
	for _, p := range parts {
		if strings.TrimSpace(p) != "" {
			out = append(out, p)
		}
	}
	return strings.Join(out, " ")
}

func sortedSlugs() []string {
	s := append([]string(nil), order...)
	sort.Strings(s)
	return s
}
