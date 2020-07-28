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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/term-query.asciidoc#L28>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "term": {
//       "user": {
//         "value": "Kimchy",
//         "boost": 1.0
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_term_query_e9dc1bd46146ce96910acc81fa649fc9(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:e9dc1bd46146ce96910acc81fa649fc9[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "user": {
		        "value": "Kimchy",
		        "boost": 1.0
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
	// end:e9dc1bd46146ce96910acc81fa649fc9[]
}
