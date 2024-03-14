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

//go:build bulk_default
// +build bulk_default

// This example demonstrates indexing documents using the Elasticsearch "Bulk" API
// [https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html].
//
// You can configure the number of documents and the batch size with command line flags:
//
//	go run default.go -count=10000 -batch=2500
//
// The example intentionally doesn't use any abstractions or helper functions, to
// demonstrate the low-level mechanics of working with the Bulk API: preparing
// the meta+data payloads, sending the payloads in batches,
// inspecting the error results, and printing a report.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Published time.Time `json:"published"`
	Author    Author    `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	_     = fmt.Print
	count int
	batch int
)

func init() {
	flag.IntVar(&count, "count", 1000, "Number of documents to generate")
	flag.IntVar(&batch, "batch", 255, "Number of documents to send in one batch")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.SetFlags(0)

	type bulkResponse struct {
		Errors bool `json:"errors"`
		Items  []struct {
			Index struct {
				ID     string `json:"_id"`
				Result string `json:"result"`
				Status int    `json:"status"`
				Error  struct {
					Type   string `json:"type"`
					Reason string `json:"reason"`
					Cause  struct {
						Type   string `json:"type"`
						Reason string `json:"reason"`
					} `json:"caused_by"`
				} `json:"error"`
			} `json:"index"`
		} `json:"items"`
	}

	var (
		buf bytes.Buffer
		res *esapi.Response
		err error
		raw map[string]interface{}
		blk *bulkResponse

		articles  []*Article
		indexName = "articles"

		numItems   int
		numErrors  int
		numIndexed int
		numBatches int
		currBatch  int
	)

	log.Printf(
		"\x1b[1mBulk\x1b[0m: documents [%s] batch size [%s]",
		humanize.Comma(int64(count)), humanize.Comma(int64(batch)))
	log.Println(strings.Repeat("▁", 65))

	// Create the Elasticsearch client
	//
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Generate the articles collection
	//
	names := []string{"Alice", "John", "Mary"}
	for i := 1; i < count+1; i++ {
		articles = append(articles, &Article{
			ID:        i,
			Title:     strings.Join([]string{"Title", strconv.Itoa(i)}, " "),
			Body:      "Lorem ipsum dolor sit amet...",
			Published: time.Now().Round(time.Second).UTC().AddDate(0, 0, i),
			Author: Author{
				FirstName: names[rand.Intn(len(names))],
				LastName:  "Smith",
			},
		})
	}
	log.Printf("→ Generated %s articles", humanize.Comma(int64(len(articles))))
	fmt.Print("→ Sending batch ")

	// Re-create the index
	//
	if res, err = es.Indices.Delete([]string{indexName}); err != nil {
		log.Fatalf("Cannot delete index: %s", err)
	}
	res, err = es.Indices.Create(indexName)
	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Cannot create index: %s", res)
	}

	if count%batch == 0 {
		numBatches = (count / batch)
	} else {
		numBatches = (count / batch) + 1
	}

	start := time.Now().UTC()

	// Loop over the collection
	//
	for i, a := range articles {
		numItems++

		currBatch = i / batch
		if i == count-1 {
			currBatch++
		}

		// Prepare the metadata payload
		//
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, a.ID, "\n"))
		// fmt.Printf("%s", meta) // <-- Uncomment to see the payload

		// Prepare the data payload: encode article to JSON
		//
		data, err := json.Marshal(a)
		if err != nil {
			log.Fatalf("Cannot encode article %d: %s", a.ID, err)
		}

		// Append newline to the data payload
		//
		data = append(data, "\n"...) // <-- Comment out to trigger failure for batch
		// fmt.Printf("%s", data) // <-- Uncomment to see the payload

		// // Uncomment next block to trigger indexing errors -->
		// if a.ID == 11 || a.ID == 101 {
		// 	data = []byte(`{"published" : "INCORRECT"}` + "\n")
		// }
		// // <--------------------------------------------------

		// Append payloads to the buffer (ignoring write errors)
		//
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)

		// When a threshold is reached, execute the Bulk() request with body from buffer
		//
		if i > 0 && i%batch == 0 || i == count-1 {
			fmt.Printf("[%d/%d] ", currBatch, numBatches)

			res, err = es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
			if err != nil {
				log.Fatalf("Failure indexing batch %d: %s", currBatch, err)
			}
			// If the whole request failed, print error and mark all documents as failed
			//
			if res.IsError() {
				numErrors += numItems
				if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
					log.Fatalf("Failure to to parse response body: %s", err)
				} else {
					log.Printf("  Error: [%d] %s: %s",
						res.StatusCode,
						raw["error"].(map[string]interface{})["type"],
						raw["error"].(map[string]interface{})["reason"],
					)
				}
				// A successful response might still contain errors for particular documents...
				//
			} else {
				if err := json.NewDecoder(res.Body).Decode(&blk); err != nil {
					log.Fatalf("Failure to to parse response body: %s", err)
				} else {
					for _, d := range blk.Items {
						// ... so for any HTTP status above 201 ...
						//
						if d.Index.Status > 201 {
							// ... increment the error counter ...
							//
							numErrors++

							// ... and print the response status and error information ...
							log.Printf("  Error: [%d]: %s: %s: %s: %s",
								d.Index.Status,
								d.Index.Error.Type,
								d.Index.Error.Reason,
								d.Index.Error.Cause.Type,
								d.Index.Error.Cause.Reason,
							)
						} else {
							// ... otherwise increase the success counter.
							//
							numIndexed++
						}
					}
				}
			}

			// Close the response body, to prevent reaching the limit for goroutines or file handles
			//
			res.Body.Close()

			// Reset the buffer and items counter
			//
			buf.Reset()
			numItems = 0
		}
	}

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	//
	fmt.Print("\n")
	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)

	if numErrors > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(numIndexed)),
			humanize.Comma(int64(numErrors)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(numIndexed))),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(numIndexed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(numIndexed))),
		)
	}
}
