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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/nested.asciidoc#L58>
//
// --------------------------------------------------------------------------------
// GET my-index-000001/_search
// {
//   "query": {
//     "bool": {
//       "must": [
//         { "match": { "user.first": "Alice" }},
//         { "match": { "user.last":  "Smith" }}
//       ]
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_nested_5553cf7a02c22f616cd994747f2dd5a5(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5553cf7a02c22f616cd994747f2dd5a5[]
	res, err := es.Search(
		es.Search.WithIndex("my-index-000001"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "bool": {
		      "must": [
		        {
		          "match": {
		            "user.first": "Alice"
		          }
		        },
		        {
		          "match": {
		            "user.last": "Smith"
		          }
		        }
		      ]
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
	// end:5553cf7a02c22f616cd994747f2dd5a5[]
}
