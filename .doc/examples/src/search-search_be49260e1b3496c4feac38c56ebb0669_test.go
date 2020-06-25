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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/search.asciidoc#L596>
//
// --------------------------------------------------------------------------------
// GET /twitter/_search?q=user:kimchy
// --------------------------------------------------------------------------------

func Test_search_search_be49260e1b3496c4feac38c56ebb0669(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:be49260e1b3496c4feac38c56ebb0669[]
	res, err := es.Search(
		es.Search.WithIndex("twitter"),
		es.Search.WithQuery("user:kimchy"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:be49260e1b3496c4feac38c56ebb0669[]
}
