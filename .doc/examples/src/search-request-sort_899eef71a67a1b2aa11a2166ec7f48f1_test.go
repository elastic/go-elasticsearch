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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L370>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "sort" : [
//         { "price" : {"unmapped_type" : "long"} }
//     ],
//     "query" : {
//         "term" : { "product" : "chocolate" }
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_899eef71a67a1b2aa11a2166ec7f48f1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:899eef71a67a1b2aa11a2166ec7f48f1[]
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
	// end:899eef71a67a1b2aa11a2166ec7f48f1[]
}
