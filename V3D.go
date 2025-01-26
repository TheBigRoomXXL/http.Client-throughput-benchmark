package main

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type CrawlerV3D struct {
	ctx  context.Context
	urls []string
}

// Same as v2 but with DisableKeepAlives
func NewCrawlerV3D(ctx context.Context, urls []string) Crawler {
	return &CrawlerV3D{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV3D) Run() {
	transport := &http.Transport{
		TLSHandshakeTimeout: 15 * time.Second,
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
