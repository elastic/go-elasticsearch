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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L164>
//
// --------------------------------------------------------------------------------
// PUT /my_index/_mapping
// {
//   "properties": {
//     "name": {
//       "properties": {
//         "last": {
//           "type": "text"
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_0bbd30b9be3e54ff3028b9f4459634d2(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0bbd30b9be3e54ff3028b9f4459634d2[]
	res, err := es.Indices.PutMapping(
		[]string{"my_index"},
		strings.NewReader(`{
		  "properties": {
		    "name": {
		      "properties": {
		        "last": {
		          "type": "text"
		        }
		      }
		    }
		  }
		}`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:0bbd30b9be3e54ff3028b9f4459634d2[]
}
