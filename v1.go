package main

import (
	"context"
	"net/http"
	"sync"
)

type CrawlerV1 struct {
	ctx  context.Context
	urls []string
}

func NewCrawlerV1(ctx context.Context, urls []string) Crawler {
	return &CrawlerV1{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV1) Run() {
	client := http.DefaultClient
	wg := &sync.WaitGroup{}
	wg.Add(len(c.urls))
	for _, url := range c.urls {
		go func() {
			defer wg.Done()
			_, err := client.Get(url)
			resultChan <- err
		}()
	}
	wg.Wait()
}
