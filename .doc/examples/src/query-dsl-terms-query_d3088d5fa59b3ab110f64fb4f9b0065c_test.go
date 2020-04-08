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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/terms-query.asciidoc#L145>
//
// --------------------------------------------------------------------------------
// PUT my_index/_doc/1
// {
//   "color":   ["blue", "green"]
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_terms_query_d3088d5fa59b3ab110f64fb4f9b0065c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d3088d5fa59b3ab110f64fb4f9b0065c[]
	res, err := es.Index(
		"my_index",
		strings.NewReader(`{
		  "color": [
		    "blue",
		    "green"
		  ]
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:d3088d5fa59b3ab110f64fb4f9b0065c[]
}
