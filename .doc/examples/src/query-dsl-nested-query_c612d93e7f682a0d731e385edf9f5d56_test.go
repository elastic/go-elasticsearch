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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/nested-query.asciidoc#L23>
//
// --------------------------------------------------------------------------------
// PUT /my_index
// {
//     "mappings" : {
//         "properties" : {
//             "obj1" : {
//                 "type" : "nested"
//             }
//         }
//     }
// }
//
// --------------------------------------------------------------------------------

func Test_query_dsl_nested_query_c612d93e7f682a0d731e385edf9f5d56(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:c612d93e7f682a0d731e385edf9f5d56[]
	res, err := es.Indices.Create(
		"my_index",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "obj1": {
		        "type": "nested"
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
	// end:c612d93e7f682a0d731e385edf9f5d56[]
}
