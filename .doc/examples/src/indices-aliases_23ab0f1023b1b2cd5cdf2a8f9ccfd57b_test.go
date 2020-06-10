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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/aliases.asciidoc#L294>
//
// --------------------------------------------------------------------------------
// PUT /test1
// {
//   "mappings": {
//     "properties": {
//       "user" : {
//         "type": "keyword"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_aliases_23ab0f1023b1b2cd5cdf2a8f9ccfd57b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:23ab0f1023b1b2cd5cdf2a8f9ccfd57b[]
	res, err := es.Indices.Create(
		"test1",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "user": {
		        "type": "keyword"
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
	// end:23ab0f1023b1b2cd5cdf2a8f9ccfd57b[]
}
