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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/multi-match-query.asciidoc#L130>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "dis_max": {
//       "queries": [
//         { "match": { "subject": "brown fox" }},
//         { "match": { "message": "brown fox" }}
//       ],
//       "tie_breaker": 0.3
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_multi_match_query_b0eaf67e5cce24ef8889bf20951ccec1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b0eaf67e5cce24ef8889bf20951ccec1[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "dis_max": {
		      "queries": [
		        {
		          "match": {
		            "subject": "brown fox"
		          }
		        },
		        {
		          "match": {
		            "message": "brown fox"
		          }
		        }
		      ],
		      "tie_breaker": 0.3
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
	// end:b0eaf67e5cce24ef8889bf20951ccec1[]
}
