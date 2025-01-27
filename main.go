package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"
	"time"
)

const NB_REQUEST = 50_000
const SEMAPHORE_SIZE = 4000

var resultChan = make(chan error)

var crawlers = map[string]func(ctx context.Context, urls []string) Crawler{
	"V1":  NewCrawlerV1,
	"V2":  NewCrawlerV2,
	"V3A": NewCrawlerV3A,
	"V3B": NewCrawlerV3B,
	"V3C": NewCrawlerV3C,
	"V3D": NewCrawlerV3D,
	"V3E": NewCrawlerV3E,
	"V4":  NewCrawlerV4,
	"V5":  NewCrawlerV5,
}

type Crawler interface {
	Run()
}

func main() {
	urls := readCSV()
	if len(os.Args) > 2 {
		log.Fatal("expect the crawler version as argument")
	}
	v := strings.ToUpper(os.Args[1])
	newCrawler, ok := crawlers[v]
	if !ok {
		log.Fatal("invalid crawler version")
	}

	// Concelation context
	ctx, stop := context.WithCancel(context.Background())
	time.AfterFunc(30*time.Second, stop)  // stop after 30s
	signalChan := make(chan os.Signal, 1) // Stop on interupt
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		stop()
	}()

	fmt.Printf("NB_REQUEST = %d\n", NB_REQUEST)
	fmt.Printf("SEMAPHORE_SIZE = %d\n", SEMAPHORE_SIZE)
	benchmark(ctx, v, newCrawler(ctx, urls))

}

func benchmark(ctx context.Context, label string, crawler Crawler) {
	err := 0
	errTimeouts := 0
	errTimeoutsLookup := 0
	errCertificate := 0
	errNoSuchHost := 0
	errNetworkUnreachable := 0
	total := 0

	// Run
	start := time.Now()
	go crawler.Run()

OuterLoop:
	for {
		select {
		case <-ctx.Done():
			break OuterLoop
		case result := <-resultChan:
			total++
			if result != nil {
				errorMsg := result.Error()
				slog.Error(errorMsg)
				if strings.Contains(errorMsg, "timeout") {
					errTimeouts++
					if strings.Contains(errorMsg, "dial tcp: lookup") {
						errTimeoutsLookup++
					}
				} else if strings.Contains(errorMsg, "failed to verify certificate") {
					errCertificate++
				} else if strings.Contains(errorMsg, "no such host") {
					errNoSuchHost++
				} else if strings.Contains(errorMsg, "network is unreachable") {
					errNetworkUnreachable++
				}
				err++
			}

			if total == NB_REQUEST {
				break OuterLoop
			}
		}
	}
	duration := time.Since(start)
	errOthers := err - errTimeouts - errNoSuchHost - errCertificate - errNetworkUnreachable

	// Report
	fmt.Printf("%s results:\n", label)
	fmt.Printf("%d requests done in %.2fs - %.2freq/s\n", total, duration.Seconds(), float64(total)/duration.Seconds())
	fmt.Printf("%d errs (%.3f%%)\n", err, percent(err, total))
	fmt.Printf(" └─┬─┬ %d timeouts (%.3f%%)\n", errTimeouts, percent(errTimeouts, err))
	fmt.Printf("   │ └─ %d timeouts on host lookup (%.3f%%)\n", errTimeoutsLookup, percent(errTimeouts, err))
	fmt.Printf("   ├── %d no such host (%.3f%%)\n", errNoSuchHost, percent(errNoSuchHost, err))
	fmt.Printf("   ├── %d bad certificate (%.3f%%)\n", errCertificate, percent(errCertificate, err))
	fmt.Printf("   ├── %d network is unreachable (%.3f%%)\n", errNetworkUnreachable, percent(errNetworkUnreachable, err))
	fmt.Printf("   └── %d others (%.3f%%)\n", errOthers, percent(errOthers, err))
}

func readCSV() []string {
	f, err := os.Open("urls.csv")
	if err != nil {
		log.Fatal("Unable to read inputs file: ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse inputs file: ", err)
	}

	urls := make([]string, NB_REQUEST)
	for i := 1; i < NB_REQUEST; i++ {
		urls[i] = records[i][0] + "://" + ReverseHostname(records[i][1]) + records[i][2]
	}
	return urls
}

func ReverseHostname(hostname string) string {
	labels := strings.Split(hostname, ".")
	slices.Reverse(labels)
	return strings.Join(labels, ".")
}

func percent(a int, b int) float32 {
	return 100 * float32(a) / float32(b)
}
