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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/terms-query.asciidoc#L186>
//
// --------------------------------------------------------------------------------
// GET my_index/_search?pretty
// {
//   "query": {
//     "terms": {
//         "color" : {
//             "index" : "my_index",
//             "id" : "2",
//             "path" : "color"
//         }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_terms_query_d1bcf2eb63a462bfdcf01a68e68d5b4a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d1bcf2eb63a462bfdcf01a68e68d5b4a[]
	res, err := es.Search(
		es.Search.WithIndex("my_index"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "terms": {
		      "color": {
		        "index": "my_index",
		        "id": "2",
		        "path": "color"
		      }
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
	// end:d1bcf2eb63a462bfdcf01a68e68d5b4a[]
}
