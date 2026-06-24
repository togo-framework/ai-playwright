<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/ai-playwright</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/ai-playwright"><img src="https://pkg.go.dev/badge/github.com/togo-framework/ai-playwright.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/ai-playwright
```

<!-- /togo-header -->

# ai-playwright

A togo **AI data-source** plugin — render JS-heavy pages with a headless **Chromium (Playwright)** and return text/HTML, so `ai-rag` ingest and agents can read client-rendered sites.

```
togo install togo-framework/ai-playwright
```
Install the browser once at runtime: `go run github.com/playwright-community/playwright-go/cmd/playwright install chromium`.

## Use
- Go: `aiplaywright.FromKernel(k).Render(ctx, "https://example.com")` → `Result{Title,Text,HTML}`
- REST: `POST /api/ai/playwright` `{"url":"…"}`

Part of the [togo AI kit](https://to-go.dev/ai). MIT.

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
