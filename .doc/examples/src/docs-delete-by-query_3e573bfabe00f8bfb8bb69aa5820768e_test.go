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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L449>
//
// --------------------------------------------------------------------------------
// GET _refresh
// POST twitter/_search?size=0&filter_path=hits.total
// {
//   "query": {
//     "range": {
//       "likes": {
//         "lt": 10
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_3e573bfabe00f8bfb8bb69aa5820768e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3e573bfabe00f8bfb8bb69aa5820768e[]
	{
		res, err := es.Indices.Refresh()
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithIndex("twitter"),
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "range": {
		      "likes": {
		        "lt": 10
		      }
		    }
		  }
		}`)),
			es.Search.WithFilterPath("hits.total"),
			es.Search.WithSize(0),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:3e573bfabe00f8bfb8bb69aa5820768e[]
}
