// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/search.asciidoc#L678>
//
// --------------------------------------------------------------------------------
// GET /_all/_search?q=user.id:kimchy
// --------------------------------------------------------------------------------

func Test_search_search_8eb4376d4ff2587d0f14e63a7fb29a2a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8eb4376d4ff2587d0f14e63a7fb29a2a[]
	res, err := es.Search(
		es.Search.WithIndex("_all"),
		es.Search.WithQuery("user.id:kimchy"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:8eb4376d4ff2587d0f14e63a7fb29a2a[]
}
