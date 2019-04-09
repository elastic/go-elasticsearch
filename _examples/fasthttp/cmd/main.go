package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/_examples/fasthttp"
)

var (
	count     int
	durations []time.Duration
)

func init() {
	if envCount := os.Getenv("COUNT"); envCount != "" {
		i, err := strconv.Atoi(envCount)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
		count = i
	} else {
		count = 1000
	}
}

func main() {
	log.SetFlags(0)

	es, err := elasticsearch.NewClient(
		elasticsearch.Config{Transport: &fasthttp.Transport{}},
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	t := time.Now()
	for i := 0; i < count; i++ {
		t0 := time.Now()
		res, err := es.Info()
		durations = append(durations, time.Now().Sub(t0))
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		res.Body.Close()
	}

	sorted := append(make([]time.Duration, 0), durations...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	log.Printf(
		"%d requests in %.2f (%.1f req/s) | min: %s / max: %s / mean: %s",
		count,
		time.Now().Sub(t).Seconds(),
		float64(count)/time.Now().Sub(t).Seconds(),
		sorted[0],
		sorted[(len(sorted)-1)],
		sorted[(len(sorted)-1)/2],
	)
}
