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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/function-score-query.asciidoc#L41>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "function_score": {
//       "query": { "match_all": {} },
//       "boost": "5", <1>
//       "functions": [
//         {
//           "filter": { "match": { "test": "bar" } },
//           "random_score": {}, <2>
//           "weight": 23
//         },
//         {
//           "filter": { "match": { "test": "cat" } },
//           "weight": 42
//         }
//       ],
//       "max_boost": 42,
//       "score_mode": "max",
//       "boost_mode": "multiply",
//       "min_score": 42
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_function_score_query_edcfadbfb14d97a2f5e6e21ef7039818(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:edcfadbfb14d97a2f5e6e21ef7039818[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "function_score": {
		      "query": {
		        "match_all": {}
		      },
		      "boost": "5",
		      "functions": [
		        {
		          "filter": {
		            "match": {
		              "test": "bar"
		            }
		          },
		          "random_score": {},
		          "weight": 23
		        },
		        {
		          "filter": {
		            "match": {
		              "test": "cat"
		            }
		          },
		          "weight": 42
		        }
		      ],
		      "max_boost": 42,
		      "score_mode": "max",
		      "boost_mode": "multiply",
		      "min_score": 42
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
	// end:edcfadbfb14d97a2f5e6e21ef7039818[]
}
