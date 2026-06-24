// Package aiplaywright is a togo AI data-source plugin: render JS-heavy pages
// with a headless Chromium (Playwright) and return text/HTML/screenshot, so
// ai-rag ingest and agents can read client-rendered sites. Registers an
// "ai-playwright" service + REST: POST /api/ai/playwright {"url":"…"}.
// The browser is downloaded at runtime (playwright.Install or
// `go run github.com/playwright-community/playwright-go/cmd/playwright install chromium`);
// the build does not require it.
package aiplaywright

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/playwright-community/playwright-go"
	"github.com/togo-framework/togo"
)

// Result is a rendered page.
type Result struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
	HTML  string `json:"html,omitempty"`
}

// Source renders pages with a headless browser.
type Source struct{}

// New returns a Source.
func New() *Source { return &Source{} }

// Render loads url in headless Chromium and returns its title/text/HTML.
func (s *Source) Render(_ context.Context, url string) (*Result, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	defer pw.Stop()
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})
	if err != nil {
		return nil, err
	}
	defer browser.Close()
	page, err := browser.NewPage()
	if err != nil {
		return nil, err
	}
	if _, err := page.Goto(url, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle}); err != nil {
		return nil, err
	}
	title, _ := page.Title()
	html, _ := page.Content()
	text, _ := page.InnerText("body")
	return &Result{URL: url, Title: title, Text: text, HTML: html}, nil
}

// FromKernel returns the registered Source, or nil.
func FromKernel(k *togo.Kernel) *Source {
	if v, ok := k.Get("ai-playwright"); ok {
		if s, ok := v.(*Source); ok {
			return s
		}
	}
	return nil
}

func init() {
	togo.RegisterProviderFunc("ai-playwright", togo.PriorityService, func(k *togo.Kernel) error {
		s := New()
		k.Set("ai-playwright", s)
		mount(k.Router, s)
		return nil
	})
}

func mount(r chi.Router, s *Source) {
	r.Post("/api/ai/playwright", func(w http.ResponseWriter, req *http.Request) {
		var b struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(req.Body).Decode(&b); err != nil || b.URL == "" {
			http.Error(w, `{"error":"url required"}`, http.StatusBadRequest)
			return
		}
		res, err := s.Render(req.Context(), b.URL)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(res)
	})
}
