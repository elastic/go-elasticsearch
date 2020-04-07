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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L445>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "sort" : [
//         {
//             "_geo_distance" : {
//                 "pin.location" : {
//                     "lat" : 40,
//                     "lon" : -70
//                 },
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

func Test_search_request_sort_979d25dff2d8987119410291ad47b0d1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:979d25dff2d8987119410291ad47b0d1[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "sort": [
		    {
		      "_geo_distance": {
		        "pin.location": {
		          "lat": 40,
		          "lon": -70
		        },
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
	// end:979d25dff2d8987119410291ad47b0d1[]
}
