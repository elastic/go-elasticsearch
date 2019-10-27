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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/bool-query.asciidoc#L130>
//
// --------------------------------------------------------------------------------
// GET _search
// {
//   "query": {
//     "constant_score": {
//       "filter": {
//         "term": {
//           "status": "active"
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_bool_query_162b5b693b713f0bfab1209d59443c46(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:162b5b693b713f0bfab1209d59443c46[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "constant_score": {
		      "filter": {
		        "term": {
		          "status": "active"
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
	// end:162b5b693b713f0bfab1209d59443c46[]
}
