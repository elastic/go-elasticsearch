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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L700>
//
// --------------------------------------------------------------------------------
// POST test/_search?filter_path=hits.total
// {
//   "query": {
//     "match": {
//       "flag": "foo"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_abd4fc3ce7784413a56fe2dcfe2809b5(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:abd4fc3ce7784413a56fe2dcfe2809b5[]
	res, err := es.Search(
		es.Search.WithIndex("test"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "flag": "foo"
		    }
		  }
		}`)),
		es.Search.WithFilterPath("hits.total"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:abd4fc3ce7784413a56fe2dcfe2809b5[]
}
