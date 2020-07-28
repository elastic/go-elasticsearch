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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/terms-query.asciidoc#L160>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001/_doc/2
// {
//   "color":   "blue"
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_terms_query_344b4144244d57f87c6aa4652b100b25(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:344b4144244d57f87c6aa4652b100b25[]
	res, err := es.Index(
		"my-index-000001",
		strings.NewReader(`{
		  "color": "blue"
		}`),
		es.Index.WithDocumentID("2"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:344b4144244d57f87c6aa4652b100b25[]
}
