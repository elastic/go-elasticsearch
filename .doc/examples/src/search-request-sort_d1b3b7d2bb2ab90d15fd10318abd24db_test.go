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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L11>
//
// --------------------------------------------------------------------------------
// PUT /my_index
// {
//     "mappings": {
//         "properties": {
//             "post_date": { "type": "date" },
//             "user": {
//                 "type": "keyword"
//             },
//             "name": {
//                 "type": "keyword"
//             },
//             "age": { "type": "integer" }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_d1b3b7d2bb2ab90d15fd10318abd24db(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d1b3b7d2bb2ab90d15fd10318abd24db[]
	res, err := es.Indices.Create(
		"my_index",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "post_date": {
		        "type": "date"
		      },
		      "user": {
		        "type": "keyword"
		      },
		      "name": {
		        "type": "keyword"
		      },
		      "age": {
		        "type": "integer"
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
	// end:d1b3b7d2bb2ab90d15fd10318abd24db[]
}
