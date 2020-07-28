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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L423>
//
// --------------------------------------------------------------------------------
// PUT /users
// {
//   "mappings" : {
//     "properties": {
//       "user_id": {
//         "type": "long"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_bd5918ab903c0889bb1f09c8c2466e43(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:bd5918ab903c0889bb1f09c8c2466e43[]
	res, err := es.Indices.Create(
		"users",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "user_id": {
		        "type": "long"
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
	// end:bd5918ab903c0889bb1f09c8c2466e43[]
}
