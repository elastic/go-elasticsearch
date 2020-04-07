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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/range-query.asciidoc#L157>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "range" : {
//             "timestamp" : {
//                 "gte" : "now-1d/d",
//                 "lt" :  "now/d"
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_range_query_67ceac4bf2d9ac7cc500390544cdcb41(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:67ceac4bf2d9ac7cc500390544cdcb41[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "range": {
		      "timestamp": {
		        "gte": "now-1d/d",
		        "lt": "now/d"
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
	// end:67ceac4bf2d9ac7cc500390544cdcb41[]
}
