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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/match-query.asciidoc#L241>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "match" : {
//             "message" : {
//                 "query" : "to be or not to be",
//                 "operator" : "and",
//                 "zero_terms_query": "all"
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_match_query_0ac9916f47a2483b89c1416684af322a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0ac9916f47a2483b89c1416684af322a[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "message": {
		        "query": "to be or not to be",
		        "operator": "and",
		        "zero_terms_query": "all"
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
	// end:0ac9916f47a2483b89c1416684af322a[]
}
