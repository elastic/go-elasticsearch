// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

//go:build ignore
// +build ignore

// This example demonstrates indexing documents using the Elasticsearch "Bulk" API
// [https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html].
//
// You can configure the number of documents and the batch size with command line flags:
//
//     go run benchmark.go.go -count=10000 -batch=2500 -dataset=sample
//     -indexName index-name
//
// The example intentionally doesn't use any abstractions or helper functions, to
// demonstrate the low-level mechanics of working with the Bulk API: preparing
// the meta+data payloads, sending the payloads in batches,
// inspecting the error results, and printing a report.
//
//
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
)

type Data struct {
	ID       int                    `json:"id"`
	Document map[string]interface{} `json:"-"`
}

var (
	_         = fmt.Print
	count     int
	batch     int
	dataset   string
	indexName string
)

func init() {
	flag.IntVar(&count, "count", 1000, "Number of documents to generate")
	flag.IntVar(&batch, "batch", 255, "Number of documents to send in one batch")
	flag.StringVar(&dataset, "dataset", "sample", "Enter the data set name")
	flag.StringVar(&indexName, "IndexName", "index", "Enter the index name")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

// Creates a mapping templates from details present in the mapping.json and
func getDefaultTemplateMapping(indexName string, mapperContent map[string]interface{}) map[string]interface{} {
	mapper := make(map[string]interface{})

	// NOTE: If you want to apply the change only to this index, then remove this '*'
	mapper["index_patterns"] = []string{indexName + "*"}
	mapper["mappings"] = mapperContent
	return mapper
}

// Creates the Index with given indexName and mapping details as mapperContent
func createIndex(client *elasticsearch.Client, indexName string, mapperContent map[string]interface{}) {
	jsonBody, err := json.Marshal(getDefaultTemplateMapping(indexName, mapperContent))
	createRequest := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  bytes.NewReader(jsonBody),
	}

	res, err := createRequest.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error creating request to create client : %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error creating the index : %s, err - %s", indexName, res.String())
	}
	log.Printf("Successfully created the index : %s", indexName)

}

// Checks if the index is present, if it is then it deletes and re-creates it
func reCreateIndex(client *elasticsearch.Client, indexName string, mapperContent map[string]interface{}) {
	if _, err := client.Indices.Delete([]string{indexName}); err != nil {
		log.Fatalf("Cannot delete index: %s", err)
	}
	createIndex(client, indexName, mapperContent)
}

// Begin main method
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

		bulkData       []*Data
		rootDataFolder = "data/"

		numItems   int
		numErrors  int
		numIndexed int
		numBatches int
		currBatch  int
	)

	// Read the mapping.json file
	mappingFile, err := os.Open(filepath.Join("data", dataset, "mapping.json"))
	if err != nil {
		log.Fatalf("Error reading the mapping.json file : %s", err)
	}

	var mappingEnvelope map[string]interface{}
	json.NewDecoder(mappingFile).Decode(&mappingEnvelope)

	// Read the document.json file
	documentFile := filepath.Join(rootDataFolder, dataset, "document.json")
	dataContent, err := ioutil.ReadFile(documentFile)

	if err != nil {
		log.Fatalf("Error reading the file: %s, err - %v", documentFile, err)
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(dataContent, &data)

	if err != nil {
		log.Fatalf("Error while unmarshalling the data, err - %v", err)
	}

	log.Printf(
		"\x1b[1mBulk\x1b[0m: documents [%s] batch size [%s]",
		humanize.Comma(int64(count)), humanize.Comma(int64(batch)))
	log.Println(strings.Repeat("▁", 65))

	// Create the Elasticsearch client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	reCreateIndex(es, indexName, mappingEnvelope)

	for i := 1; i < count+1; i++ {
		bulkData = append(bulkData, &Data{ID: i, Document: data})
	}
	log.Printf("→ Generated %s documents", humanize.Comma(int64(len(bulkData))))
	fmt.Print("→ Sending batch ")

	if count%batch == 0 {
		numBatches = (count / batch)
	} else {
		numBatches = (count / batch) + 1
	}

	start := time.Now().UTC()

	// Loop over the collection
	for i, a := range bulkData {
		numItems++

		currBatch = i / batch
		if i == count-1 {
			currBatch++
		}

		// Prepare the metadata payload
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d", "_type" : "doc" } }%s`, a.ID, "\n"))

		data_, err := json.Marshal(a.Document)
		if err != nil {
			log.Fatalf("Cannot encode article %d: %s", a.ID, err)
		}

		data_ = append(data_, "\n"...)

		// Append payloads to the buffer (ignoring write errors)
		buf.Grow(len(meta) + len(data_))
		buf.Write(meta)
		buf.Write(data_)

		// When a threshold is reached, execute the Bulk() request with body from buffer
		if i > 0 && i%batch == 0 || i == count-1 {
			fmt.Printf("[%d/%d] ", currBatch, numBatches)

			res, err = es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
			if err != nil {
				log.Fatalf("Failure indexing batch %d: %s", currBatch, err)
			}
			// If the whole request failed, print error and mark all documents as failed
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
			} else {
				if err := json.NewDecoder(res.Body).Decode(&blk); err != nil {
					log.Fatalf("Failure to to parse response body: %s", err)
				} else {
					for _, d := range blk.Items {
						// ... so for any HTTP status above 201 ...
						if d.Index.Status > 201 {
							// ... increment the error counter ...
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
							numIndexed++
						}
					}
				}
			}

			// Close the response body, to prevent reaching the limit for goroutines or file handles
			res.Body.Close()

			// Reset the buffer and items counter
			buf.Reset()
			numItems = 0
		}
	}

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
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
