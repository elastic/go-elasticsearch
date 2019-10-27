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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping.asciidoc#L173>
//
// --------------------------------------------------------------------------------
// PUT /my-index/_mapping
// {
//   "properties": {
//     "employee-id": {
//       "type": "keyword",
//       "index": false
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_71ba9033107882f61cdc3b32fc73568d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:71ba9033107882f61cdc3b32fc73568d[]
	res, err := es.Indices.PutMapping(
		[]string{"my-index"},
		strings.NewReader(`{
		  "properties": {
		    "employee-id": {
		      "type": "keyword",
		      "index": false
		    }
		  }
		}`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:71ba9033107882f61cdc3b32fc73568d[]
}
