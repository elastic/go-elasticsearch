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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L94>
//
// --------------------------------------------------------------------------------
// PUT /my_index/_doc/1?refresh
// {
//    "product": "chocolate",
//    "price": [20, 4]
// }
//
// POST /_search
// {
//    "query" : {
//       "term" : { "product" : "chocolate" }
//    },
//    "sort" : [
//       {"price" : {"order" : "asc", "mode" : "avg"}}
//    ]
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_b997885974522ef439d5e345924cc5ba(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b997885974522ef439d5e345924cc5ba[]
	{
		res, err := es.Index(
			"my_index",
			strings.NewReader(`{
		  "product": "chocolate",
		  "price": [
		    20,
		    4
		  ]
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "product": "chocolate"
		    }
		  },
		  "sort": [
		    {
		      "price": {
		        "order": "asc",
		        "mode": "avg"
		      }
		    }
		  ]
		}`)),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:b997885974522ef439d5e345924cc5ba[]
}
