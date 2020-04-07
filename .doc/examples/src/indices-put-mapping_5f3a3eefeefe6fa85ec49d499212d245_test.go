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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L257>
//
// --------------------------------------------------------------------------------
// PUT /my_index/_mapping
// {
//   "properties": {
//     "city": {
//       "type": "text",
//       "fields": {
//         "raw": {
//           "type": "keyword"
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_5f3a3eefeefe6fa85ec49d499212d245(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5f3a3eefeefe6fa85ec49d499212d245[]
	res, err := es.Indices.PutMapping(
		[]string{"my_index"},
		strings.NewReader(`{
		  "properties": {
		    "city": {
		      "type": "text",
		      "fields": {
		        "raw": {
		          "type": "keyword"
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
	// end:5f3a3eefeefe6fa85ec49d499212d245[]
}
