package main

import (
	"context"
	"net/http"
	"sync"
)

type CrawlerV2 struct {
	ctx  context.Context
	urls []string
}

// Same as v1 but with semaphore to avoid overloading
func NewCrawlerV2(ctx context.Context, urls []string) Crawler {
	return &CrawlerV2{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV2) Run() {
	client := http.DefaultClient
	sem := make(chan struct{}, 400)
	wg := &sync.WaitGroup{}
	wg.Add(len(c.urls))
	for _, url := range c.urls {
		go func() {
			defer wg.Done()

			// aquire / release semaphore
			sem <- struct{}{}
			defer func() { <-sem }()

			_, err := client.Get(url)
			resultChan <- err
		}()
	}
	wg.Wait()
}
