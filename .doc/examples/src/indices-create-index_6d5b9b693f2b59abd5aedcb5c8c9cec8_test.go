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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/create-index.asciidoc#L143>
//
// --------------------------------------------------------------------------------
// PUT /test
// {
//   "aliases": {
//     "alias_1": {},
//     "alias_2": {
//       "filter": {
//         "term": { "user": "kimchy" }
//       },
//       "routing": "kimchy"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_create_index_6d5b9b693f2b59abd5aedcb5c8c9cec8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6d5b9b693f2b59abd5aedcb5c8c9cec8[]
	res, err := es.Indices.Create(
		"test",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "aliases": {
		    "alias_1": {},
		    "alias_2": {
		      "filter": {
		        "term": {
		          "user": "kimchy"
		        }
		      },
		      "routing": "kimchy"
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:6d5b9b693f2b59abd5aedcb5c8c9cec8[]
}
