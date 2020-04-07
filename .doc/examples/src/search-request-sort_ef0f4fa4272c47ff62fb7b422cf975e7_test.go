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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L346>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "sort" : [
//         { "price" : {"missing" : "_last"} }
//     ],
//     "query" : {
//         "term" : { "product" : "chocolate" }
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_ef0f4fa4272c47ff62fb7b422cf975e7(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:ef0f4fa4272c47ff62fb7b422cf975e7[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "sort": [
		    {
		      "price": {
		        "missing": "_last"
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
	// end:ef0f4fa4272c47ff62fb7b422cf975e7[]
}
