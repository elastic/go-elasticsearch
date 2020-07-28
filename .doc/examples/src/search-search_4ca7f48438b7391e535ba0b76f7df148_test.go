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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/search.asciidoc#L684>
//
// --------------------------------------------------------------------------------
// GET /*/_search?q=user.id:kimchy
// --------------------------------------------------------------------------------

func Test_search_search_4ca7f48438b7391e535ba0b76f7df148(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4ca7f48438b7391e535ba0b76f7df148[]
	res, err := es.Search(
		es.Search.WithIndex("*"),
		es.Search.WithQuery("user.id:kimchy"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:4ca7f48438b7391e535ba0b76f7df148[]
}
