// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/estransport"
)

func main() {
	log.SetFlags(0)

	var es *elasticsearch.Client

	// This collection of examples demonstrates how to configure the default logger of the client.
	//
	// To enable logging, pass an io.Writer as the "LogOutput" option in the client configuration.
	//
	// ==============================================================================================
	//
	// The default logger formatter writes basic information about
	// the request and response as plain text to the output.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &estransport.TextLogger{Output: os.Stdout},
	})
	run(es, "Text")

	// ==============================================================================================
	//
	// The "LogFormatColor" formatter is optimized for displaying
	// information in the terminal during development.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &estransport.ColorLogger{Output: os.Stdout},
	})
	run(es, "Color")

	// ==============================================================================================
	//
	// To log the request and response bodies, use the respective configuration options.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &estransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
	})
	run(es, "Request/Response Body")

	// ==============================================================================================
	//
	// The "LogFormatCurl" formatter writes the information formatted as executable curl commands,
	// pretty-printing the response (when enabled), useful eg. for sharing or debugging.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &estransport.CurlLogger{Output: os.Stdout, EnableResponseBody: true},
	})
	run(es, "Curl")

	// ==============================================================================================
	//
	// The "LogFormatJSON" formatter writes the information as JSON, suitable for log ingestion.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &estransport.JSONLogger{Output: os.Stdout},
	})
	run(es, "JSON")
}

// ------------------------------------------------------------------------------------------------

func run(es *elasticsearch.Client, name string) {
	log.Println("███", fmt.Sprintf("\x1b[1m%s\x1b[0m", name), strings.Repeat("█", 75-len(name)))

	es.Delete("test", "1")
	es.Exists("test", "1")

	es.Index(
		"test",
		strings.NewReader(`{"title" : "logging"}`),
		es.Index.WithRefresh("true"),
		es.Index.WithPretty(),
		es.Index.WithFilterPath("result", "_id"),
	)

	es.Search(es.Search.WithQuery("[FAIL"))

	res, err := es.Search(
		es.Search.WithIndex("test"),
		es.Search.WithBody(strings.NewReader(`{"query" : {"match" : { "title" : "logging" } } }`)),
		es.Search.WithSize(1),
		es.Search.WithPretty(),
		es.Search.WithFilterPath("took", "hits.hits"),
	)

	resBody := res.String()
	if resBody == "" {
		log.Fatal("Response body is empty")
	}

	if err != nil {
		log.Fatalf("Error:   %s", err)
	}

	log.Print("\n")
}
