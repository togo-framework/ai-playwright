# ai-playwright — documentation

Headless-browser (Playwright) data-source for the togo AI kit

## Overview

Package aiplaywright is a togo AI data-source plugin: render JS-heavy pages
with a headless Chromium (Playwright) and return text/HTML/screenshot, so
ai-rag ingest and agents can read client-rendered sites. Registers an
"ai-playwright" service + REST: POST /api/ai/playwright {"url":"…"}.
The browser is downloaded at runtime (playwright.Install or
`go run github.com/playwright-community/playwright-go/cmd/playwright install chromium`);

## Install

```bash
togo install togo-framework/ai-playwright
```

A capability plugin — it self-registers on boot; no driver selector needed.

## Configuration

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

_No environment variables read directly (uses the kernel/base config or the app DB)._

## Usage

```go
// A data source for ai-rag / agents: fetch/scrape/search web content.
docs, err := playwright.FromKernel(k).Fetch(ctx, "https://example.com")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-playwright
- Full README: ../README.md
