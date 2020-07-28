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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L180>
//
// --------------------------------------------------------------------------------
// PUT /index_double
// {
//   "mappings": {
//     "properties": {
//       "field": { "type": "date" }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_3bc872dbcdad8ff02cbaea39e7f38352(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3bc872dbcdad8ff02cbaea39e7f38352[]
	res, err := es.Indices.Create(
		"index_double",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "field": {
		        "type": "date"
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
	// end:3bc872dbcdad8ff02cbaea39e7f38352[]
}
