package main

import (
	"context"
	"sync"

	"github.com/valyala/fasthttp"
)

type CrawlerV4 struct {
	ctx  context.Context
	urls []string
}

// Same as v2 but with DisableKeepAlives
func NewCrawlerV4(ctx context.Context, urls []string) Crawler {
	return &CrawlerV4{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV4) Run() {

	sem := make(chan struct{}, SEMAPHORE_SIZE)
	wg := &sync.WaitGroup{}
	wg.Add(len(c.urls))

	for _, url := range c.urls {
		go func() {
			defer wg.Done()

			// aquire / release semaphore
			sem <- struct{}{}
			defer func() { <-sem }()

			req := fasthttp.AcquireRequest()
			resp := fasthttp.AcquireResponse()
			defer fasthttp.ReleaseRequest(req)
			defer fasthttp.ReleaseResponse(resp)

			req.SetRequestURI(url)

			err := fasthttp.DoRedirects(req, resp, 10)
			resultChan <- err
		}()
	}
	wg.Wait()
}
