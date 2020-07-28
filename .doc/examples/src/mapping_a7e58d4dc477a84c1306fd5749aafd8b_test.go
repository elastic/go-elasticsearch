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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping.asciidoc#L147>
//
// --------------------------------------------------------------------------------
// PUT /my-index-000001
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

func Test_mapping_a7e58d4dc477a84c1306fd5749aafd8b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:a7e58d4dc477a84c1306fd5749aafd8b[]
	res, err := es.Indices.Create(
		"my-index-000001",
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
	// end:a7e58d4dc477a84c1306fd5749aafd8b[]
}
