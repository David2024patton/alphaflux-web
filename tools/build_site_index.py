#!/usr/bin/env python3
"""Build assets/site-index.json: the Ask AlphaFlux knowledge index.

Chunks are extracted from the real room pages (the same source of truth the
deep view uses), so the site guide's answers update whenever the pages do.
Run from the repo root: python3 tools/build_site_index.py
"""
import json
import re
from html.parser import HTMLParser
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent


class Node:
    def __init__(self, tag, attrs):
        self.tag = tag
        self.attrs = dict(attrs)
        self.children = []

    def text(self):
        parts = []
        for c in self.children:
            parts.append(c if isinstance(c, str) else c.text())
        return re.sub(r"\s+", " ", " ".join(parts)).strip()

    def find_all(self, tag=None, cls=None):
        out = []
        for c in self.children:
            if isinstance(c, Node):
                if (tag is None or c.tag == tag) and (
                    cls is None or cls in c.attrs.get("class", "").split()
                ):
                    out.append(c)
                out.extend(c.find_all(tag, cls))
        return out

    def first(self, tag=None, cls=None):
        found = self.find_all(tag, cls)
        return found[0] if found else None


class TreeBuilder(HTMLParser):
    VOID = {"img", "br", "hr", "input", "meta", "link", "source"}

    def __init__(self):
        super().__init__(convert_charrefs=True)
        self.root = Node("root", {})
        self.stack = [self.root]

    def handle_starttag(self, tag, attrs):
        node = Node(tag, attrs)
        self.stack[-1].children.append(node)
        if tag not in self.VOID:
            self.stack.append(node)

    def handle_startendtag(self, tag, attrs):
        self.stack[-1].children.append(Node(tag, attrs))

    def handle_endtag(self, tag):
        for i in range(len(self.stack) - 1, 0, -1):
            if self.stack[i].tag == tag:
                del self.stack[i:]
                break

    def handle_data(self, data):
        if data.strip():
            self.stack[-1].children.append(data)


def parse(path):
    parser = TreeBuilder()
    parser.feed(path.read_text(encoding="utf-8"))
    return parser.root


def section_chunks(href, label, main):
    chunks = []
    sections = [
        s
        for s in main.find_all("section")
        if {"card", "plan", "alacarte"} & set(s.attrs.get("class", "").split())
    ]
    for sec in sections:
        h2 = sec.first("h2")
        heading = h2.text() if h2 else ""
        if heading in ("Start on the free plan", "What AlphaFlux does"):
            continue
        blocks = []
        for det in sec.find_all("details"):
            q = det.first("summary")
            a = " ".join(p.text() for p in det.find_all("p"))
            if q:
                chunks.append(
                    {"href": href, "label": label, "h": "Question: " + q.text(), "t": a[:700]}
                )
        # block walk mirroring the deep view extraction
        def walk(node, in_li=False, in_table=False):
            for c in node.children:
                if not isinstance(c, Node):
                    continue
                if c.tag == "table":
                    continue
                if c is h2:
                    continue
                if c.tag == "details":
                    continue
                if c.tag == "li":
                    h3 = c.first("h3")
                    dt = c.first("dt")
                    if h3:
                        body = " ".join(p.text() for p in c.find_all("p"))
                        blocks.append(h3.text() + ": " + body)
                    elif dt:
                        body = " ".join(d.text() for d in c.find_all("dd"))
                        blocks.append(dt.text() + ": " + body)
                    else:
                        blocks.append(c.text())
                    continue
                if c.tag == "p" and not in_li:
                    blocks.append(c.text())
                    continue
                if c.tag == "h3" and not in_li:
                    blocks.append(c.text() + ":")
                    continue
                walk(c, in_li or c.tag == "li", in_table)

        walk(sec)
        if "plan" in sec.attrs.get("class", "").split():
            amt = sec.first(cls="amount")
            if amt:
                annual = amt.attrs.get("data-annual", "")
                if amt.text() == "0":
                    price = "Price: $0 per month, free for life"
                elif annual and annual != amt.text():
                    price = f"Price: ${amt.text()} per month, or ${annual} per month billed annually"
                else:
                    price = f"Price: ${amt.text()} per month"
                blocks.insert(0, price + ".")
        body = " ".join(b for b in blocks if b).strip()
        if body:
            chunks.append({"href": href, "label": label, "h": heading or label, "t": body[:900]})
    if not chunks:
        h1 = main.first("h1")
        body = main.text()
        if body:
            chunks.append(
                {"href": href, "label": label, "h": h1.text() if h1 else label, "t": body[:900]}
            )
    return chunks


def main():
    index_html = (ROOT / "index.html").read_text(encoding="utf-8")
    m = re.search(r"var QUICK = (\{.*?\});?\s*\n\s*var CHECK", index_html, re.S)
    quick = json.loads(m.group(1))
    chunks = []
    for href, data in quick.items():
        if not data.get("deep"):
            continue
        fname = href.strip("/") + ".html"
        path = ROOT / fname
        if not path.exists():
            continue
        doc = parse(path)
        main_node = doc.first("main")
        if main_node:
            chunks.extend(section_chunks(href, data["title"], main_node))
    # A plans-at-a-glance chunk so cost questions get the whole picture first.
    for i, c in enumerate(chunks):
        if c["href"] == "/pricing":
            chunks.insert(
                i,
                {
                    "href": "/pricing",
                    "label": "Pricing",
                    "h": "Plans at a glance",
                    "t": "AlphaFlux pricing has six plans. Free is $0 per month, free for life. Solo is $19 per month, or $15 per month billed annually. Starter is $69 per month, or $55 per month billed annually. Growth is $150 per month, or $120 billed annually. Pro is $300 per month, or $240 billed annually, and adds the AI receptionist. Scale is $420 per month, or $336 billed annually. White label comes in three sizes: $999, $1,773, or $2,999 per month. You do not have to buy a plan: modules can be added one at a time on the Free plan.",
                },
            )
            break
    # Real served pages that no home tile points at, still part of the story.
    for fname, href, label in [
        ("calls-and-messaging.html", "/calls-and-messaging", "AI Phone Answering"),
    ]:
        doc = parse(ROOT / fname)
        main_node = doc.first("main")
        if main_node:
            chunks.extend(section_chunks(href, label, main_node))
    chunks.extend(
        [
            {
                "href": "/about",
                "label": "About AlphaFlux",
                "h": "What this company is",
                "t": "AlphaFlux is the front office for service businesses. It answers your phones around the clock, captures every lead, books jobs onto your schedule, sends invoices and takes payments, and keeps your crew and your customers in the loop, so no call, lead, or job slips through the cracks. It is built for service businesses: pest control, plumbing, HVAC, roofing, cleaning, landscaping, security patrol, solar and mobile units, logistics, construction, manufacturing, and nonprofits.",
            },
            {
                "href": "/about",
                "label": "About AlphaFlux",
                "h": "How AlphaFlux helps your business",
                "t": "AlphaFlux answers your phones, captures your leads, and books your jobs while you are out doing the work. The phones, the schedule, the customer record, the invoices and the crew app all point at the same account, so nothing slips through the cracks and nobody types the same thing twice. You can start on the free plan: twenty customers, one location, every core module, free permanently, and move up when the work outgrows it.",
            },
            {
                "href": "/waitlist",
                "label": "AlphaVision",
                "h": "AlphaVision",
                "t": "AlphaVision is the AlphaFlux product line for TVs and big screens: the streaming style tile home screen this very site is modelled on. It is not on sale yet; join the waitlist on this site to hear when it opens.",
            },
            {
                "href": "https://cortex.alphaflux.net",
                "label": "Cortex",
                "h": "Cortex",
                "t": "Cortex is the AlphaFlux knowledge brain: it stores a business's documents, pages and records as connected nodes with search across all of it, at cortex.alphaflux.net.",
            },
            {
                "href": "https://critic.alphaflux.net",
                "label": "Critic",
                "h": "Critic",
                "t": "Critic is the AlphaFlux quality workbench: work is checked by a panel of critics before it is allowed in, at critic.alphaflux.net.",
            },
        ]
    )
    out = ROOT / "assets" / "site-index.json"
    out.write_text(json.dumps({"chunks": chunks}, ensure_ascii=False, indent=1), encoding="utf-8")
    print(f"wrote {out} with {len(chunks)} chunks, {out.stat().st_size} bytes")


if __name__ == "__main__":
    main()
