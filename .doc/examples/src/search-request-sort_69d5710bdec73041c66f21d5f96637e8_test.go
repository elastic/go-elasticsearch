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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L192>
//
// --------------------------------------------------------------------------------
// PUT /index_long
// {
//   "mappings": {
//     "properties": {
//       "field": { "type": "date_nanos" }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_69d5710bdec73041c66f21d5f96637e8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:69d5710bdec73041c66f21d5f96637e8[]
	res, err := es.Indices.Create(
		"index_long",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "field": {
		        "type": "date_nanos"
		      }
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:69d5710bdec73041c66f21d5f96637e8[]
}
