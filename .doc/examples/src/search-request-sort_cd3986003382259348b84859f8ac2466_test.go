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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L30>
//
// --------------------------------------------------------------------------------
// GET /my-index-000001/_search
// {
//   "sort" : [
//     { "post_date" : {"order" : "asc"}},
//     "user",
//     { "name" : "desc" },
//     { "age" : "desc" },
//     "_score"
//   ],
//   "query" : {
//     "term" : { "user" : "kimchy" }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_cd3986003382259348b84859f8ac2466(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:cd3986003382259348b84859f8ac2466[]
	res, err := es.Search(
		es.Search.WithIndex("my-index-000001"),
		es.Search.WithBody(strings.NewReader(`{
		  "sort": [
		    {
		      "post_date": {
		        "order": "asc"
		      }
		    },
		    "user",
		    {
		      "name": "desc"
		    },
		    {
		      "age": "desc"
		    },
		    "_score"
		  ],
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
	// end:cd3986003382259348b84859f8ac2466[]
}
