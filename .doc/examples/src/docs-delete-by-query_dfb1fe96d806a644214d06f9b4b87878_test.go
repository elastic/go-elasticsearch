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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L401>
//
// --------------------------------------------------------------------------------
// POST twitter/_delete_by_query?scroll_size=5000
// {
//   "query": {
//     "term": {
//       "user": "kimchy"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_dfb1fe96d806a644214d06f9b4b87878(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:dfb1fe96d806a644214d06f9b4b87878[]
	res, err := es.DeleteByQuery(
		[]string{"twitter"},
		strings.NewReader(`{
		  "query": {
		    "term": {
		      "user": "kimchy"
		    }
		  }
		}`),
		es.DeleteByQuery.WithScrollSize(5000),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:dfb1fe96d806a644214d06f9b4b87878[]
}
