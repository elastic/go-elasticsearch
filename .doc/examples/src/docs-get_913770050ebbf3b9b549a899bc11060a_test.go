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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L306>
//
// --------------------------------------------------------------------------------
// PUT twitter
// {
//    "mappings": {
//        "properties": {
//           "counter": {
//              "type": "integer",
//              "store": false
//           },
//           "tags": {
//              "type": "keyword",
//              "store": true
//           }
//        }
//    }
// }
// --------------------------------------------------------------------------------

func Test_docs_get_913770050ebbf3b9b549a899bc11060a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:913770050ebbf3b9b549a899bc11060a[]
	res, err := es.Indices.Create(
		"twitter",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "counter": {
		        "type": "integer",
		        "store": false
		      },
		      "tags": {
		        "type": "keyword",
		        "store": true
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
	// end:913770050ebbf3b9b549a899bc11060a[]
}
