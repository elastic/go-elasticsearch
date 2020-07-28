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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L338>
//
// --------------------------------------------------------------------------------
// PUT /my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "user_id": {
//         "type": "keyword",
//         "ignore_above": 20
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_5f1ed9cfdc149763b444acfbe10b0e16(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5f1ed9cfdc149763b444acfbe10b0e16[]
	res, err := es.Indices.Create(
		"my-index-000001",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "user_id": {
		        "type": "keyword",
		        "ignore_above": 20
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
	// end:5f1ed9cfdc149763b444acfbe10b0e16[]
}
