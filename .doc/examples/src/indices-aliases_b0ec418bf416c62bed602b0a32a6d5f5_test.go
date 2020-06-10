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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/aliases.asciidoc#L431>
//
// --------------------------------------------------------------------------------
// PUT /alias1/_doc/1
// {
//     "foo": "bar"
// }
// --------------------------------------------------------------------------------

func Test_indices_aliases_b0ec418bf416c62bed602b0a32a6d5f5(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b0ec418bf416c62bed602b0a32a6d5f5[]
	res, err := es.Index(
		"alias1",
		strings.NewReader(`{
		  "foo": "bar"
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:b0ec418bf416c62bed602b0a32a6d5f5[]
}
