// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"context"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/_examples/fasthttp"
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

	// Test sending the body as POST
	//
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go func() {
		es.Info()
		if _, err := es.Search(
			es.Search.WithBody(strings.NewReader(`{"query":{"match":{"title":"foo"}}}`)),
			es.Search.WithPretty(),
		); err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		cancel()
	}()

	select {
	case <-ctx.Done():
		if ctx.Err() != context.Canceled {
			log.Fatalf("Timeout: %s", ctx.Err())
		}
	}

	// Run benchmark
	//
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
