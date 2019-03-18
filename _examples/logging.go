// +build ignore

// This example demonstrates how to use the logging facility of the Elasticsearch client.
//
package main

import (
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

func main() {
	log.SetFlags(0)

	// Initialize the client with logger output set to STDOUT.
	//
	es, err := elasticsearch.NewClient(elasticsearch.Config{LogOutput: os.Stdout})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Index a document
	//
	res, err := es.Index(
		"test",
		strings.NewReader(`{"title" : "logging"}`),
		es.Index.WithRefresh("true"),
		es.Index.WithPretty(),
		es.Index.WithFilterPath("result", "_id"),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	// Search a document
	//
	res, err = es.Search(
		es.Search.WithIndex("test"),
		es.Search.WithBody(strings.NewReader(`{"query" : {"match" : { "title" : "logging" } } }`)),
		es.Search.WithSize(1),
		es.Search.WithPretty(),
		es.Search.WithFilterPath("took", "hits.hits"),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
}
