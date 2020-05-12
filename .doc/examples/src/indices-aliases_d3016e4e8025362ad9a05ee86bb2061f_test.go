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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/aliases.asciidoc#L12>
//
// --------------------------------------------------------------------------------
// POST /_aliases
// {
//     "actions" : [
//         { "add" : { "index" : "twitter", "alias" : "alias1" } }
//     ]
// }
// --------------------------------------------------------------------------------

func Test_indices_aliases_d3016e4e8025362ad9a05ee86bb2061f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d3016e4e8025362ad9a05ee86bb2061f[]
	res, err := es.Indices.UpdateAliases(strings.NewReader(`{
		  "actions": [
		    {
		      "add": {
		        "index": "twitter",
		        "alias": "alias1"
		      }
		    }
		  ]
		}`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:d3016e4e8025362ad9a05ee86bb2061f[]
}
