// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request-body.asciidoc#L7>
//
// --------------------------------------------------------------------------------
// GET /twitter/_search
// {
//   "query" : {
//     "term" : { "user" : "kimchy" }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_request_body_6f9d516e7d0c8c0c8a4b0b9c732e720d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6f9d516e7d0c8c0c8a4b0b9c732e720d[]
	res, err := es.Search(
		es.Search.WithIndex("twitter"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "user": "kimchy"
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:6f9d516e7d0c8c0c8a4b0b9c732e720d[]
}
