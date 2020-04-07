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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L263>
//
// --------------------------------------------------------------------------------
// POST /_search
// {
//    "query" : {
//       "term" : { "product" : "chocolate" }
//    },
//    "sort" : [
//        {
//           "offer.price" : {
//              "mode" :  "avg",
//              "order" : "asc",
//              "nested": {
//                 "path": "offer",
//                 "filter": {
//                    "term" : { "offer.color" : "blue" }
//                 }
//              }
//           }
//        }
//     ]
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_de139866a220124360e5e27d1a736ea4(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:de139866a220124360e5e27d1a736ea4[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "product": "chocolate"
		    }
		  },
		  "sort": [
		    {
		      "offer.price": {
		        "mode": "avg",
		        "order": "asc",
		        "nested": {
		          "path": "offer",
		          "filter": {
		            "term": {
		              "offer.color": "blue"
		            }
		          }
		        }
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
	// end:de139866a220124360e5e27d1a736ea4[]
}
