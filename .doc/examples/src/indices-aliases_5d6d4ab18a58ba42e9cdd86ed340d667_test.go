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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/aliases.asciidoc#L348>
//
// --------------------------------------------------------------------------------
// POST /_aliases
// {
//   "actions": [
//     {
//       "add": {
//         "index": "test",
//         "alias": "alias1",
//         "routing": "1"
//       }
//     }
//   ]
// }
// --------------------------------------------------------------------------------

func Test_indices_aliases_5d6d4ab18a58ba42e9cdd86ed340d667(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5d6d4ab18a58ba42e9cdd86ed340d667[]
	res, err := es.Indices.UpdateAliases(strings.NewReader(`{
		  "actions": [
		    {
		      "add": {
		        "index": "test",
		        "alias": "alias1",
		        "routing": "1"
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
	// end:5d6d4ab18a58ba42e9cdd86ed340d667[]
}
