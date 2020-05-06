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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L366>
//
// --------------------------------------------------------------------------------
// PUT twitter/_doc/2?routing=user1
// {
//     "counter" : 1,
//     "tags" : ["white"]
// }
// --------------------------------------------------------------------------------

func Test_docs_get_0ba0b2db24852abccb7c0fc1098d566e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0ba0b2db24852abccb7c0fc1098d566e[]
	res, err := es.Index(
		"twitter",
		strings.NewReader(`{
		  "counter": 1,
		  "tags": [
		    "white"
		  ]
		}`),
		es.Index.WithDocumentID("2"),
		es.Index.WithRouting("user1"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:0ba0b2db24852abccb7c0fc1098d566e[]
}
