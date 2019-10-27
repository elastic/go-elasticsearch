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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping.asciidoc#L144>
//
// --------------------------------------------------------------------------------
// PUT /my-index
// {
//   "mappings": {
//     "properties": {
//       "age":    { "type": "integer" },  <1>
//       "email":  { "type": "keyword"  }, <2>
//       "name":   { "type": "text"  }     <3>
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_d8b2a88b5eca99d3691ad3cd40266736(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d8b2a88b5eca99d3691ad3cd40266736[]
	res, err := es.Indices.Create(
		"my-index",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "age": {
		        "type": "integer"
		      },
		      "email": {
		        "type": "keyword"
		      },
		      "name": {
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
	// end:d8b2a88b5eca99d3691ad3cd40266736[]
}
