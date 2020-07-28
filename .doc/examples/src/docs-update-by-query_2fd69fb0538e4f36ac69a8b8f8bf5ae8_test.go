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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L355>
//
// --------------------------------------------------------------------------------
// POST twitter/_update_by_query
// {
//   "script": {
//     "source": "ctx._source.likes++",
//     "lang": "painless"
//   },
//   "query": {
//     "term": {
//       "user": "kimchy"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_2fd69fb0538e4f36ac69a8b8f8bf5ae8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:2fd69fb0538e4f36ac69a8b8f8bf5ae8[]
	res, err := es.UpdateByQuery(
		[]string{"twitter"},
		es.UpdateByQuery.WithBody(strings.NewReader(`{
		  "script": {
		    "source": "ctx._source.likes++",
		    "lang": "painless"
		  },
		  "query": {
		    "term": {
		      "user": "kimchy"
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:2fd69fb0538e4f36ac69a8b8f8bf5ae8[]
}
