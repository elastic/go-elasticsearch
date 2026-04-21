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

//go:build logging_default
// +build logging_default

// This collection of examples demonstrates how to configure the default logger of the typed client.
//
// To enable logging, pass a logger implementation to the "Logger" option in the client configuration.
//
// You can use one of the bundled loggers, or a custom "elastictransport.Logger" interface implementation.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

func main() {
	log.SetFlags(0)

	var es *elasticsearch.TypedClient

	// ==============================================================================================
	//
	// "TextLogger" writes basic information about the request and response as plain text to the output.
	//
	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.TextLogger{Output: os.Stdout}))
	run(es, "Text")
	_ = es.Close(context.Background())

	// ==============================================================================================
	//
	// "ColorLogger" is optimized for displaying information in the terminal during development.
	//
	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.ColorLogger{Output: os.Stdout}))
	run(es, "Color")
	_ = es.Close(context.Background())

	// ==============================================================================================
	//
	// To log the request and response bodies, use the respective configuration options.
	//
	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.ColorLogger{
		Output:             os.Stdout,
		EnableRequestBody:  true,
		EnableResponseBody: true,
	}))
	run(es, "Request/Response Body")
	_ = es.Close(context.Background())

	// ==============================================================================================
	//
	// "CurlLogger" writes the information formatted as runnable curl commands,
	// pretty-printing the response body (when enabled), useful eg. for sharing.
	//
	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.CurlLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true}))
	run(es, "Curl")
	_ = es.Close(context.Background())

	// ==============================================================================================
	//
	// "JSONLogger" writes the information as JSON and is suitable for production logging.
	//
	es, _ = elasticsearch.NewTyped(elasticsearch.WithLogger(&elastictransport.JSONLogger{Output: os.Stdout}))
	run(es, "JSON")
	_ = es.Close(context.Background())
}

// ------------------------------------------------------------------------------------------------

func run(es *elasticsearch.TypedClient, name string) {
	log.Println("███", fmt.Sprintf("\x1b[1m%s\x1b[0m", name), strings.Repeat("█", 75-len(name)))

	ctx := context.Background()

	es.Delete("test", "1").Do(ctx)
	es.Exists("test", "1").Do(ctx)

	es.Index("test").
		Document(map[string]string{"title": "logging"}).
		Refresh(refresh.True).
		Pretty(true).
		FilterPath("result", "_id").
		Do(ctx)

	// Intentionally send a bad query-string (?q=) to show an error being logged.
	es.Search().Q("{FAIL").Do(ctx)

	_, err := es.Search().
		Index("test").
		Query(esdsl.NewMatchQuery("title", "logging")).
		Size(1).
		Pretty(true).
		FilterPath("took", "hits.hits").
		Do(ctx)

	if err != nil {
		log.Fatalf("Error:   %s", err)
	}

	log.Print("\n")
}
