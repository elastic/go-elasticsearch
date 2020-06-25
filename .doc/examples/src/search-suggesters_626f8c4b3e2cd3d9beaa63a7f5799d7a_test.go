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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/suggesters.asciidoc#L8>
//
// --------------------------------------------------------------------------------
// POST twitter/_search
// {
//   "query" : {
//     "match": {
//       "message": "tring out Elasticsearch"
//     }
//   },
//   "suggest" : {
//     "my-suggestion" : {
//       "text" : "tring out Elasticsearch",
//       "term" : {
//         "field" : "message"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_suggesters_626f8c4b3e2cd3d9beaa63a7f5799d7a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:626f8c4b3e2cd3d9beaa63a7f5799d7a[]
	res, err := es.Search(
		es.Search.WithIndex("twitter"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "message": "tring out Elasticsearch"
		    }
		  },
		  "suggest": {
		    "my-suggestion": {
		      "text": "tring out Elasticsearch",
		      "term": {
		        "field": "message"
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
	// end:626f8c4b3e2cd3d9beaa63a7f5799d7a[]
}
