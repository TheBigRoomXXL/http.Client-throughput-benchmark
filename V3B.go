package main

import (
	"context"
	"net/http"
	"sync"
)

type CrawlerV3B struct {
	ctx  context.Context
	urls []string
}

// Same as v2 but with  increased ForceAttemptHTTP2
func NewCrawlerV3B(ctx context.Context, urls []string) Crawler {
	return &CrawlerV3B{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV3B) Run() {
	transport := &http.Transport{
		ForceAttemptHTTP2: true,
	}
	client := http.Client{Transport: transport}
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
