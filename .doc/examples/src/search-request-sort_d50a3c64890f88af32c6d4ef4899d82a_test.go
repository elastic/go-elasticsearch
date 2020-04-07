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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L471>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "sort" : [
//         {
//             "_geo_distance" : {
//                 "pin.location" : "40,-70",
//                 "order" : "asc",
//                 "unit" : "km"
//             }
//         }
//     ],
//     "query" : {
//         "term" : { "user" : "kimchy" }
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_d50a3c64890f88af32c6d4ef4899d82a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d50a3c64890f88af32c6d4ef4899d82a[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "sort": [
		    {
		      "_geo_distance": {
		        "pin.location": "40,-70",
		        "order": "asc",
		        "unit": "km"
		      }
		    }
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
	// end:d50a3c64890f88af32c6d4ef4899d82a[]
}
