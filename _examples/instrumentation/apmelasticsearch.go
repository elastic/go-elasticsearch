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
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"os/user"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/level"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmelasticsearch"
)

const count = 100

var (
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))

	faint   = color.New(color.Faint)
	bold    = color.New(color.Bold)
	boldRed = color.New(color.FgRed).Add(color.Bold)

	currentUser string
)

func init() {
	if tWidth < 0 {
		tWidth = 0
	}

	if u, err := user.Current(); err != nil {
		boldRed.Printf("ERROR: %s\n", err)
		currentUser = "N/A"
	} else {
		currentUser = u.Username
	}
}

func main() {
	log.SetFlags(0)
	start := time.Now()

	// Create new Elasticsearch typed client ...
	//
	es, err := elasticsearch.NewTyped(
		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		// ... using the "apmelasticsearch" wrapper for instrumentation
		elasticsearch.WithTransportOptions(elastictransport.WithTransport(apmelasticsearch.WrapRoundTripper(http.DefaultTransport))),
		// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// Set up the "done" channel
	//
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)

	// Set up tickers
	//
	tickers := struct {
		Info   *time.Ticker
		Index  *time.Ticker
		Health *time.Ticker
		Search *time.Ticker
	}{
		Info:   time.NewTicker(time.Second),
		Index:  time.NewTicker(500 * time.Millisecond),
		Health: time.NewTicker(5 * time.Second),
		Search: time.NewTicker(10 * time.Second),
	}
	defer tickers.Info.Stop()
	defer tickers.Index.Stop()
	defer tickers.Health.Stop()
	defer tickers.Search.Stop()

	// Initialize the context
	//
	ctx := context.Background()

	// Perform API calls
	//
	for {
		select {
		case <-done:
			fmt.Print("\n")
			fmt.Println(strings.Repeat("━", tWidth))
			faint.Printf("Finished in %s\n\n", time.Now().Sub(start).Truncate(time.Second))
			return

		// -> Info
		//
		case <-tickers.Info.C:
			func() {
				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Set up the APM transaction
				txn := apm.DefaultTracer.StartTransaction("Info()", "monitoring")
				// Add current user to the transaction metadata
				txn.Context.SetUsername(currentUser)
				// Store the transaction in a context
				ctx := apm.ContextWithTransaction(ctx, txn)
				// Mark the transaction as completed
				defer txn.End()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				_, err := es.Info().Do(ctx)
				if err != nil {
					boldRed.Printf("Error getting response: %s\n", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}

				faint.Println("OK")
			}()

		// -> Index
		//
		case t := <-tickers.Index.C:
			func(t time.Time) {
				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Set up the APM transaction
				txn := apm.DefaultTracer.StartTransaction("Index()", "indexing")
				// Add current user to the transaction metadata
				txn.Context.SetUsername(currentUser)
				// Store the transaction in a context
				ctx := apm.ContextWithTransaction(ctx, txn)
				// Mark the transaction as completed
				defer txn.End()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				// Fail some requests by sending an empty body.
				req := es.Index("test")
				if t.Second()%4 == 0 {
					req = req.Raw(strings.NewReader(``))
				} else {
					req = req.Document(map[string]string{"timestamp": t.Format(time.RFC3339)})
				}

				_, err := req.Do(ctx)
				if err != nil {
					boldRed.Printf("Error getting response: %s\n", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}

				faint.Println("OK")
			}(t)

		// -> Health
		//
		case <-tickers.Health.C:
			func() {
				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Set up the APM transaction
				txn := apm.DefaultTracer.StartTransaction("Cluster.Health()", "monitoring")
				// Add current user to the transaction metadata
				txn.Context.SetUsername(currentUser)
				// Store the transaction in a context
				ctx := apm.ContextWithTransaction(ctx, txn)
				// Mark the transaction as completed
				defer txn.End()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				_, err := es.Cluster.Health().Level(level.Indices).Do(ctx)
				if err != nil {
					boldRed.Printf("Error getting response: %s\n", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}

				faint.Println("OK")
			}()

		// -> Search
		//
		case <-tickers.Search.C:
			func() {
				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Set up the APM transaction
				txn := apm.DefaultTracer.StartTransaction("Search()", "searching")
				// Add current user to the transaction metadata
				txn.Context.SetUsername(currentUser)
				// Add a custom tag to the transaction metadata
				txn.Context.SetTag("stat_group", "foo")
				// Store the transaction in a context
				ctx := apm.ContextWithTransaction(ctx, txn)
				// Mark the transaction as completed
				defer txn.End()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				// Randomly trigger an error
				if rand.Intn(10) > 8 {
					err := errors.New("Unexpected error")
					boldRed.Println("ERROR:", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}

				res, err := es.Search().
					Index("test").
					Sort(esdsl.NewSortOptions().AddSortOption("timestamp", esdsl.NewFieldSort(sortorder.Desc))).
					Size(1).
					// Annotate this search request; https://www.elastic.co/guide/en/elasticsearch/reference/current/search.html#stats-groups
					Stats("foo").
					Do(ctx)
				if err != nil {
					boldRed.Printf("Error getting response: %s\n", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}

				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Create a custom span within the transaction
				sp := txn.StartSpan("UI/Render", "searching", nil)
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
				faint.Printf("OK ; %dms\n", res.Took)
				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Mark the span as completed
				sp.End()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
			}()
		}
	}
}
