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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/suggesters.asciidoc#L51>
//
// --------------------------------------------------------------------------------
// POST _search
// {
//   "suggest": {
//     "my-suggest-1" : {
//       "text" : "tring out Elasticsearch",
//       "term" : {
//         "field" : "message"
//       }
//     },
//     "my-suggest-2" : {
//       "text" : "kmichy",
//       "term" : {
//         "field" : "user"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_suggesters_2533e4b36ae837eaecda08407ecb6383(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:2533e4b36ae837eaecda08407ecb6383[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "suggest": {
		    "my-suggest-1": {
		      "text": "tring out Elasticsearch",
		      "term": {
		        "field": "message"
		      }
		    },
		    "my-suggest-2": {
		      "text": "kmichy",
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
	// end:2533e4b36ae837eaecda08407ecb6383[]
}
