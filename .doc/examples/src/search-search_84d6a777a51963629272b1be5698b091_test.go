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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/search.asciidoc#L623>
//
// --------------------------------------------------------------------------------
// GET /twitter/_search
// {
//   "query": {
//     "term": {
//       "user": "kimchy"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_search_84d6a777a51963629272b1be5698b091(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:84d6a777a51963629272b1be5698b091[]
	res, err := es.Search(
		es.Search.WithIndex("twitter"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "user": "kimchy"
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:84d6a777a51963629272b1be5698b091[]
}
