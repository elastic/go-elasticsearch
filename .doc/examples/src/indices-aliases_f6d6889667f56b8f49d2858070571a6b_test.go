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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/aliases.asciidoc#L409>
//
// --------------------------------------------------------------------------------
// POST /_aliases
// {
//     "actions" : [
//         {
//             "add" : {
//                  "index" : "test",
//                  "alias" : "alias1",
//                  "is_write_index" : true
//             }
//         },
//         {
//             "add" : {
//                  "index" : "test2",
//                  "alias" : "alias1"
//             }
//         }
//     ]
// }
// --------------------------------------------------------------------------------

func Test_indices_aliases_f6d6889667f56b8f49d2858070571a6b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:f6d6889667f56b8f49d2858070571a6b[]
	res, err := es.Indices.UpdateAliases(strings.NewReader(`{
		  "actions": [
		    {
		      "add": {
		        "index": "test",
		        "alias": "alias1",
		        "is_write_index": true
		      }
		    },
		    {
		      "add": {
		        "index": "test2",
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
	// end:f6d6889667f56b8f49d2858070571a6b[]
}
