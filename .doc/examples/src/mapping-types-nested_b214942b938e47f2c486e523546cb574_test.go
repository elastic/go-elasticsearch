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
// GET my_index/_search
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

func Test_mapping_types_nested_b214942b938e47f2c486e523546cb574(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b214942b938e47f2c486e523546cb574[]
	res, err := es.Search(
		es.Search.WithIndex("my_index"),
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
	// end:b214942b938e47f2c486e523546cb574[]
}
