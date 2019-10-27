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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/search.asciidoc#L7>
//
// --------------------------------------------------------------------------------
// GET /twitter/_search?q=tag:wow
// --------------------------------------------------------------------------------

func Test_search_search_9bdd3c0d47e60c8cfafc8109f9369922(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:9bdd3c0d47e60c8cfafc8109f9369922[]
	res, err := es.Search(
		es.Search.WithIndex("twitter/"),
		es.Search.WithQuery("tag:wow"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:9bdd3c0d47e60c8cfafc8109f9369922[]
}
