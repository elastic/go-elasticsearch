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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L10>
//
// --------------------------------------------------------------------------------
// POST /twitter/_delete_by_query
// {
//   "query": {
//     "match": {
//       "message": "some message"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_ebb6b59fbc9325c17e45f524602d6be2(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:ebb6b59fbc9325c17e45f524602d6be2[]
	res, err := es.DeleteByQuery(
		[]string{"twitter"},
		strings.NewReader(`{
		  "query": {
		    "match": {
		      "message": "some message"
		    }
		  }
		}`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:ebb6b59fbc9325c17e45f524602d6be2[]
}
