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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/count.asciidoc#L99>
//
// --------------------------------------------------------------------------------
// PUT /twitter/_doc/1?refresh
// {
//   "user": "kimchy"
// }
//
// GET /twitter/_count?q=user:kimchy
//
// GET /twitter/_count
// {
//   "query" : {
//     "term" : { "user" : "kimchy" }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_count_0e44c19403fbd14f97b84c8afffbe4c1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0e44c19403fbd14f97b84c8afffbe4c1[]
	{
		res, err := es.Index(
			"twitter",
			strings.NewReader(`{
		  "user": "kimchy"
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Count(
			es.Count.WithIndex("?q=user:kimchy"),
			es.Count.WithQuery("user:kimchy"),
			es.Count.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Count(
			es.Count.WithIndex(""),
			es.Count.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "user": "kimchy"
		    }
		  }
		}`)),
			es.Count.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:0e44c19403fbd14f97b84c8afffbe4c1[]
}
