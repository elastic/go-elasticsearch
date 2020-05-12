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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/function-score-query.asciidoc#L380>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "function_score": {
//             "gauss": {
//                 "date": {
//                       "origin": "2013-09-17", <1>
//                       "scale": "10d",
//                       "offset": "5d", <2>
//                       "decay" : 0.5 <2>
//                 }
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_function_score_query_ec27afee074001b0e4e393611010842b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:ec27afee074001b0e4e393611010842b[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "function_score": {
		      "gauss": {
		        "date": {
		          "origin": "2013-09-17",
		          "scale": "10d",
		          "offset": "5d",
		          "decay": 0.5
		        }
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
	// end:ec27afee074001b0e4e393611010842b[]
}
