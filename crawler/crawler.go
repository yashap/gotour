package crawler

import (
	"fmt"
)

type result struct {
	url   string
	body  string
	urls  []string
	err   error
	depth int
}

// Crawl uses fetcher to recursively crawl pages starting with url,
// to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	results := make(chan result)
	fetched := make(map[string]bool)

	// fetch url, put results in channel
	go fetch(fetcher, url, depth, results)
	fetched[url] = true

	for fetching := 1; fetching > 0; fetching-- {
		// get a result out of the channel
		r := <-results

		if r.err != nil {
			fmt.Println(r.err)
		} else {
			fmt.Printf("found: %s %q\n", r.url, r.body)

			// if max depth not reached, follow links
			if r.depth > 0 {
				for _, u := range r.urls {
					// This map access is safe because we execute all of this in the same thread
					// Only the fetching and writing to channel happens in a side goroutine
					if _, alreadyFetched := fetched[u]; !alreadyFetched {
						fetching++
						fetched[u] = true
						go fetch(fetcher, u, r.depth-1, results)
					}
				}
			}
		}
	}
}

func fetch(fetcher Fetcher, url string, depth int, results chan result) {
	body, urls, err := fetcher.Fetch(url)
	results <- result{url, body, urls, err, depth}
}
