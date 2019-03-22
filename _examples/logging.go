// +build ignore

// This example demonstrates how to use the logging facility of the Elasticsearch client.
//
package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/estransport"
)

var (
	_ = estransport.LogFormatText
	_ = fmt.Print
	_ = http.MethodGet
	_ = bytes.NewBuffer
)

func main() {
	log.SetFlags(0)

	// Initialize the client with logger output set to STDOUT.
	//
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		LogOutput:       os.Stdout,
		LogRequestBody:  true,
		LogResponseBody: true,
		LogFormat:       estransport.LogFormatColor,
		// LogFormat: estransport.LogFormatCurl,
		// LogFormat: estransport.LogFormatJSON,
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	es.Info()
	es.Delete("test", "1")
	es.Exists("test", "1")
	es.Index("test", strings.NewReader(`{"title":"Title"}`), es.Index.WithDocumentID("1"), es.Index.WithPretty())

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

	log.Println(strings.Repeat("*", 80))
	log.Println(res)
	log.Println(strings.Repeat("*", 80))

	res, err = es.Search(es.Search.WithQuery("[FAIL"))

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

	log.Println(strings.Repeat("*", 80))
	log.Println(res)
	log.Println(strings.Repeat("*", 80))
}
