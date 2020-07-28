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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L369>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "sort" : [
//     { "price" : {"unmapped_type" : "long"} }
//   ],
//   "query" : {
//     "term" : { "product" : "chocolate" }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_0bcd380315ef4691b8c79df6ca53a85f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0bcd380315ef4691b8c79df6ca53a85f[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "sort": [
		    {
		      "price": {
		        "unmapped_type": "long"
		      }
		    }
		  ],
		  "query": {
		    "term": {
		      "product": "chocolate"
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
	// end:0bcd380315ef4691b8c79df6ca53a85f[]
}
