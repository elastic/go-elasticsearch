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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/term-query.asciidoc#L94>
//
// --------------------------------------------------------------------------------
// PUT my_index
// {
//     "mappings" : {
//         "properties" : {
//             "full_text" : { "type" : "text" }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_term_query_2a1de18774f9c68cafa169847832b2bc(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:2a1de18774f9c68cafa169847832b2bc[]
	res, err := es.Indices.Create(
		"my_index",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "full_text": {
		        "type": "text"
		      }
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:2a1de18774f9c68cafa169847832b2bc[]
}
