package main

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type CrawlerV5 struct {
	ctx  context.Context
	urls []string
}

// Same as v2 but with DisableKeepAlives
func NewCrawlerV5(ctx context.Context, urls []string) Crawler {
	return &CrawlerV5{
		ctx:  ctx,
		urls: urls,
	}
}

func (c *CrawlerV5) Run() {

	sem := make(chan struct{}, SEMAPHORE_SIZE)
	wg := &sync.WaitGroup{}
	wg.Add(len(c.urls))

	client := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, 10*time.Second)
		},
	}

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

			err := client.DoRedirects(req, resp, 10)
			resultChan <- err
		}()
	}
	wg.Wait()
}
