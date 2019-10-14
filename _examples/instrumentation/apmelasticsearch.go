// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	"github.com/elastic/go-elasticsearch/v8"

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

	// Create new elasticsearch client ...
	//
	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
			// ... using the "apmelasticsearch" wrapper for instrumentation
			Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
			// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
		})
	if err != nil {
		log.Fatal("ERROR: %s", err)
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

				res, err := es.Info(es.Info.WithContext(ctx))
				if err != nil {
					boldRed.Printf("Error getting response: %s\n", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}
				defer res.Body.Close()

				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Set the response status as transaction result
				txn.Result = res.Status()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				faint.Println(res.Status())
			}()

		// -> Index
		//
		case t := <-tickers.Index.C:
			func(t time.Time) {
				// Fail some requests with empty body...
				var body io.Reader
				if t.Second()%4 == 0 {
					body = strings.NewReader(``)
				} else {
					body = strings.NewReader(`{"timestamp":"` + t.Format(time.RFC3339) + `"}`)
				}

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

				res, err := es.Index("test", body, es.Index.WithContext(ctx))
				if err != nil {
					boldRed.Printf("Error getting response: %s\n", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}
				defer res.Body.Close()

				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Set the response status as transaction result
				txn.Result = res.Status()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				faint.Println(res.Status())
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

				res, err := es.Cluster.Health(
					es.Cluster.Health.WithLevel("indices"),
					es.Cluster.Health.WithContext(ctx),
				)
				if err != nil {
					boldRed.Printf("Error getting response: %s\n", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}
				defer res.Body.Close()

				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Set the response status as transaction result
				txn.Result = res.Status()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				faint.Println(res.Status())
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

				res, err := es.Search(
					es.Search.WithIndex("test"),
					es.Search.WithSort("timestamp:desc"),
					es.Search.WithSize(1),
					// Annotate this search request; https://www.elastic.co/guide/en/elasticsearch/reference/current/search.html#stats-groups
					es.Search.WithStats("foo"),
					es.Search.WithContext(ctx),
				)
				if err != nil {
					boldRed.Printf("Error getting response: %s\n", err)
					// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
					// Capture the error
					apm.CaptureError(ctx, err).Send()
					// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
					return
				}
				defer res.Body.Close()

				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Set the response status as transaction result
				txn.Result = res.Status()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Create a custom span within the transaction
				sp := txn.StartSpan("JSON/Decode", "searching", nil)
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					sp.End()
					boldRed.Printf("Error parsing the response body: %s\n", err)
					apm.CaptureError(ctx, err).Send()
					return
				}
				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Mark the span as completed
				sp.End()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Create a custom span within the transaction
				sp = txn.StartSpan("UI/Render", "searching", nil)
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
				faint.Printf("%s ; %vms\n", res.Status(), r["took"])
				// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
				// Mark the span as completed
				sp.End()
				// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
			}()
		}
	}
}
