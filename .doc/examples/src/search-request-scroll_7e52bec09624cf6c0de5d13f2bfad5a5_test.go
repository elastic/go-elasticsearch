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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/scroll.asciidoc#L45>
//
// --------------------------------------------------------------------------------
// POST /twitter/_search?scroll=1m
// {
//     "size": 100,
//     "query": {
//         "match" : {
//             "title" : "elasticsearch"
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_request_scroll_7e52bec09624cf6c0de5d13f2bfad5a5(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:7e52bec09624cf6c0de5d13f2bfad5a5[]
	res, err := es.Search(
		es.Search.WithIndex("twitter"),
		es.Search.WithBody(strings.NewReader(`{
		  "size": 100,
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
	// end:7e52bec09624cf6c0de5d13f2bfad5a5[]
}
