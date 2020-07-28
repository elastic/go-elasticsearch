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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/nested-query.asciidoc#L41>
//
// --------------------------------------------------------------------------------
// GET /my-index-000001/_search
// {
//   "query": {
//     "nested": {
//       "path": "obj1",
//       "query": {
//         "bool": {
//           "must": [
//             { "match": { "obj1.name": "blue" } },
//             { "range": { "obj1.count": { "gt": 5 } } }
//           ]
//         }
//       },
//       "score_mode": "avg"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_nested_query_641f75862c70e25e79d249d9e0a79f03(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:641f75862c70e25e79d249d9e0a79f03[]
	res, err := es.Search(
		es.Search.WithIndex("my-index-000001"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "nested": {
		      "path": "obj1",
		      "query": {
		        "bool": {
		          "must": [
		            {
		              "match": {
		                "obj1.name": "blue"
		              }
		            },
		            {
		              "range": {
		                "obj1.count": {
		                  "gt": 5
		                }
		              }
		            }
		          ]
		        }
		      },
		      "score_mode": "avg"
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
	// end:641f75862c70e25e79d249d9e0a79f03[]
}
