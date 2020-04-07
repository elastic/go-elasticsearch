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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L349>
//
// --------------------------------------------------------------------------------
// POST twitter/_delete_by_query?conflicts=proceed
// {
//   "query": {
//     "match_all": {}
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_e21e1c26dc8687e7bf7bd2bf019a6698(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:e21e1c26dc8687e7bf7bd2bf019a6698[]
	res, err := es.DeleteByQuery(
		[]string{"twitter"},
		strings.NewReader(`{
		  "query": {
		    "match_all": {}
		  }
		}`),
		es.DeleteByQuery.WithConflicts("proceed"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:e21e1c26dc8687e7bf7bd2bf019a6698[]
}
