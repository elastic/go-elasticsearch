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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/getting-started.asciidoc#L512>
//
// --------------------------------------------------------------------------------
// GET /bank/_search
// {
//   "query": {
//     "bool": {
//       "must": [
//         { "match": { "age": "40" } }
//       ],
//       "must_not": [
//         { "match": { "state": "ID" } }
//       ]
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_getting_started_47bb632c6091ad0cd94bc660bdd309a5(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:47bb632c6091ad0cd94bc660bdd309a5[]
	res, err := es.Search(
		es.Search.WithIndex("bank"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "bool": {
		      "must": [
		        {
		          "match": {
		            "age": "40"
		          }
		        }
		      ],
		      "must_not": [
		        {
		          "match": {
		            "state": "ID"
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
	// end:47bb632c6091ad0cd94bc660bdd309a5[]
}
