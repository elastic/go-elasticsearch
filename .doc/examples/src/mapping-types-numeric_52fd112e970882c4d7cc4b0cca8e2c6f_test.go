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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/numeric.asciidoc#L22>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "number_of_bytes": {
//         "type": "integer"
//       },
//       "time_in_seconds": {
//         "type": "float"
//       },
//       "price": {
//         "type": "scaled_float",
//         "scaling_factor": 100
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_numeric_52fd112e970882c4d7cc4b0cca8e2c6f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:52fd112e970882c4d7cc4b0cca8e2c6f[]
	res, err := es.Indices.Create(
		"my-index-000001",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "number_of_bytes": {
		        "type": "integer"
		      },
		      "time_in_seconds": {
		        "type": "float"
		      },
		      "price": {
		        "type": "scaled_float",
		        "scaling_factor": 100
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
	// end:52fd112e970882c4d7cc4b0cca8e2c6f[]
}
