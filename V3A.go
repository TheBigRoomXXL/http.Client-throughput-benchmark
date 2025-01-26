package main

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"
)

type CrawlerV3A struct {
	ctx  context.Context
	urls []string
}

// [ABANDONNED] Same as v2 but with  increased DialContext timeout
func NewCrawlerV3A(ctx context.Context, urls []string) Crawler {
	return &CrawlerV3A{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV3A) Run() {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = (&net.Dialer{
		Timeout: time.Minute,
	}).DialContext
	transport.DialTLSContext = (&net.Dialer{
		Timeout: time.Minute,
	}).DialContext

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
