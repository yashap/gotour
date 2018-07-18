package main

import (
	"fmt"
	"time"

	"github.com/yashap/gotour/concurrent"
	"github.com/yashap/gotour/crawler"
	"github.com/yashap/gotour/model"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	v map[string]int
}

func main() {
	fmt.Println(">> Test model.Tree.Same")
	fmt.Println(model.NewTree(1).Same(model.NewTree(1)))
	fmt.Println(model.NewTree(1).Same(model.NewTree(2)))

	fmt.Println("\n>> Test concurrent.Counter")
	c := concurrent.NewCounter()
	key := "somekey"
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("%v: %v\n", key, c.Value(key))

	fmt.Println("\n>> Test crawler.Crawl")
	crawler.Crawl("https://golang.org/", 4, crawler.NewFakeFetcher())
}
