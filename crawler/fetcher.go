package crawler

import "fmt"

// Fetcher fetches and parses URLs.
type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// FakeFetcher is Fetcher that returns canned results.
type FakeFetcher map[string]*fakeResult

// Fetch returns the body of URL and a slice of URLs found on that page.
func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type fakeResult struct {
	body string
	urls []string
}

// NewFakeFetcher creates a new FakeFetcher
func NewFakeFetcher() FakeFetcher {
	return FakeFetcher{
		"https://golang.org/": &fakeResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &fakeResult{
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}
}
