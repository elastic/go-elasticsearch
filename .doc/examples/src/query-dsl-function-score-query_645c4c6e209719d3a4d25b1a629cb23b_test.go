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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/function-score-query.asciidoc#L241>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "function_score": {
//             "random_score": {
//                 "seed": 10,
//                 "field": "_seq_no"
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_function_score_query_645c4c6e209719d3a4d25b1a629cb23b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:645c4c6e209719d3a4d25b1a629cb23b[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "function_score": {
		      "random_score": {
		        "seed": 10,
		        "field": "_seq_no"
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
	// end:645c4c6e209719d3a4d25b1a629cb23b[]
}
