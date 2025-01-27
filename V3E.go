package main

import (
	"context"
	"net/http"
	"sync"
)

type CrawlerV3E struct {
	ctx  context.Context
	urls []string
}

// Same as v1 but with semaphore to avoid overloading
func NewCrawlerV3E(ctx context.Context, urls []string) Crawler {
	return &CrawlerV3E{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV3E) Run() {
	transport := &http.Transport{
		MaxConnsPerHost: 5000,
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
