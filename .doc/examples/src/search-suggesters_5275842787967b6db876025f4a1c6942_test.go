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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/suggesters.asciidoc#L127>
//
// --------------------------------------------------------------------------------
// POST _search
// {
//   "suggest": {
//     "text" : "tring out Elasticsearch",
//     "my-suggest-1" : {
//       "term" : {
//         "field" : "message"
//       }
//     },
//     "my-suggest-2" : {
//        "term" : {
//         "field" : "user"
//        }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_suggesters_5275842787967b6db876025f4a1c6942(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5275842787967b6db876025f4a1c6942[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "suggest": {
		    "text": "tring out Elasticsearch",
		    "my-suggest-1": {
		      "term": {
		        "field": "message"
		      }
		    },
		    "my-suggest-2": {
		      "term": {
		        "field": "user"
		      }
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
	// end:5275842787967b6db876025f4a1c6942[]
}
