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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/search.asciidoc#L413>
//
// --------------------------------------------------------------------------------
// GET /_all/_search?q=user:kimchy
// --------------------------------------------------------------------------------

func Test_search_search_8022e6a690344035b6472a43a9d122e0(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8022e6a690344035b6472a43a9d122e0[]
	res, err := es.Search(
		es.Search.WithIndex("_all/"),
		es.Search.WithQuery("user:kimchy"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:8022e6a690344035b6472a43a9d122e0[]
}
