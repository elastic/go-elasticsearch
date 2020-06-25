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
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/scroll.asciidoc#L268>
//
// --------------------------------------------------------------------------------
// GET /twitter/_search?scroll=1m
// {
//     "slice": {
//         "field": "date",
//         "id": 0,
//         "max": 10
//     },
//     "query": {
//         "match" : {
//             "title" : "elasticsearch"
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_request_scroll_fdcaba9547180439ff4b6275034a5170(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:fdcaba9547180439ff4b6275034a5170[]
	res, err := es.Search(
		es.Search.WithIndex("twitter"),
		es.Search.WithBody(strings.NewReader(`{
		  "slice": {
		    "field": "date",
		    "id": 0,
		    "max": 10
		  },
		  "query": {
		    "match": {
		      "title": "elasticsearch"
		    }
		  }
		}`)),
		es.Search.WithScroll(time.Duration(60000000000)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:fdcaba9547180439ff4b6275034a5170[]
}
