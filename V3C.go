package main

import (
	"context"
	"net/http"
	"sync"
)

type CrawlerV3C struct {
	ctx  context.Context
	urls []string
}

// Same as v2 but with DisableKeepAlives
func NewCrawlerV3C(ctx context.Context, urls []string) Crawler {
	return &CrawlerV3C{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV3C) Run() {
	transport := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{Transport: transport}
	sem := make(chan struct{}, SEMAPHORE_SIZE)
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
