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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/wildcard-query.asciidoc#L21>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "wildcard": {
//       "user": {
//         "value": "ki*y",
//         "boost": 1.0,
//         "rewrite": "constant_score"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_wildcard_query_8ad4a5ea7657eba25dc9a0ea994cf07d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8ad4a5ea7657eba25dc9a0ea994cf07d[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "wildcard": {
		      "user": {
		        "value": "ki*y",
		        "boost": 1.0,
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
	// end:8ad4a5ea7657eba25dc9a0ea994cf07d[]
}
