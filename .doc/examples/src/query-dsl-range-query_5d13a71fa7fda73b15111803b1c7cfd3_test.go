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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/range-query.asciidoc#L214>
//
// --------------------------------------------------------------------------------
// GET _search
// {
//     "query": {
//         "range" : {
//             "timestamp" : {
//                 "time_zone": "+01:00", <1>
//                 "gte": "2015-01-01 00:00:00", <2>
//                 "lte": "now" <3>
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_range_query_5d13a71fa7fda73b15111803b1c7cfd3(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5d13a71fa7fda73b15111803b1c7cfd3[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "range": {
		      "timestamp": {
		        "time_zone": "+01:00",
		        "gte": "2015-01-01 00:00:00",
		        "lte": "now"
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
	// end:5d13a71fa7fda73b15111803b1c7cfd3[]
}
