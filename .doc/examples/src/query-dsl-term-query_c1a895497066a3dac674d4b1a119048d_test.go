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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/term-query.asciidoc#L132>
//
// --------------------------------------------------------------------------------
// GET my-index-000001/_search?pretty
// {
//   "query": {
//     "term": {
//       "full_text": "Quick Brown Foxes!"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_term_query_c1a895497066a3dac674d4b1a119048d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:c1a895497066a3dac674d4b1a119048d[]
	res, err := es.Search(
		es.Search.WithIndex("my-index-000001"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "full_text": "Quick Brown Foxes!"
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
	// end:c1a895497066a3dac674d4b1a119048d[]
}
