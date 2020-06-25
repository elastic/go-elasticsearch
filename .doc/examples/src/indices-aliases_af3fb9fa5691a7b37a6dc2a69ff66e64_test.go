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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/aliases.asciidoc#L200>
//
// --------------------------------------------------------------------------------
// POST /_aliases
// {
//     "actions" : [
//         { "remove" : { "index" : "test1", "alias" : "alias1" } },
//         { "add" : { "index" : "test1", "alias" : "alias2" } }
//     ]
// }
// --------------------------------------------------------------------------------

func Test_indices_aliases_af3fb9fa5691a7b37a6dc2a69ff66e64(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:af3fb9fa5691a7b37a6dc2a69ff66e64[]
	res, err := es.Indices.UpdateAliases(strings.NewReader(`{
		  "actions": [
		    {
		      "remove": {
		        "index": "test1",
		        "alias": "alias1"
		      }
		    },
		    {
		      "add": {
		        "index": "test1",
		        "alias": "alias2"
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
	// end:af3fb9fa5691a7b37a6dc2a69ff66e64[]
}
