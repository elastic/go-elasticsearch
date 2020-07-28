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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/term-query.asciidoc#L113>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001/_doc/1
// {
//   "full_text":   "Quick Brown Foxes!"
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_term_query_111c69ca94162c1523b799a5c14723dd(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:111c69ca94162c1523b799a5c14723dd[]
	res, err := es.Index(
		"my-index-000001",
		strings.NewReader(`{
		  "full_text": "Quick Brown Foxes!"
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:111c69ca94162c1523b799a5c14723dd[]
}
