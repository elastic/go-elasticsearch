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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/regexp-query.asciidoc#L23>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "regexp": {
//             "user": {
//                 "value": "k.*y",
//                 "flags" : "ALL",
//                 "max_determinized_states": 10000,
//                 "rewrite": "constant_score"
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_regexp_query_618d5f3d35921d8cb7e9ccfbe9a4c3e3(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:618d5f3d35921d8cb7e9ccfbe9a4c3e3[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "regexp": {
		      "user": {
		        "value": "k.*y",
		        "flags": "ALL",
		        "max_determinized_states": 10000,
		        "rewrite": "constant_score"
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
	// end:618d5f3d35921d8cb7e9ccfbe9a4c3e3[]
}
