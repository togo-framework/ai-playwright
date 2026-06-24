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
