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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update.asciidoc#L86>
//
// --------------------------------------------------------------------------------
// PUT test/_doc/1
// {
//   "counter" : 1,
//   "tags" : ["red"]
// }
// --------------------------------------------------------------------------------

func Test_docs_update_b0fa301cd3c6b9db128e34114f0c1e8f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b0fa301cd3c6b9db128e34114f0c1e8f[]
	res, err := es.Index(
		"test",
		strings.NewReader(`{
		  "counter": 1,
		  "tags": [
		    "red"
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
	// end:b0fa301cd3c6b9db128e34114f0c1e8f[]
}
