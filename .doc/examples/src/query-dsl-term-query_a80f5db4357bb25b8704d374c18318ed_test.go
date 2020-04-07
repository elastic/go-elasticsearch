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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/term-query.asciidoc#L165>
//
// --------------------------------------------------------------------------------
// GET my_index/_search?pretty
// {
//   "query": {
//     "match": {
//       "full_text": "Quick Brown Foxes!"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_term_query_a80f5db4357bb25b8704d374c18318ed(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:a80f5db4357bb25b8704d374c18318ed[]
	res, err := es.Search(
		es.Search.WithIndex("my_index"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
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
	// end:a80f5db4357bb25b8704d374c18318ed[]
}
