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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/params/format.asciidoc#L13>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "date": {
//         "type":   "date",
//         "format": "yyyy-MM-dd"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_params_format_41ad6077f9c1b8d8fefab6ea1660edcd(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:41ad6077f9c1b8d8fefab6ea1660edcd[]
	res, err := es.Indices.Create(
		"my-index-000001",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "date": {
		        "type": "date",
		        "format": "yyyy-MM-dd"
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
	// end:41ad6077f9c1b8d8fefab6ea1660edcd[]
}
