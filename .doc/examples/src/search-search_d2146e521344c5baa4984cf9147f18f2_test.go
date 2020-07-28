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
// GET /my-index-000001/_search?q=user.id:kimchy
// --------------------------------------------------------------------------------

func Test_search_search_d2146e521344c5baa4984cf9147f18f2(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d2146e521344c5baa4984cf9147f18f2[]
	res, err := es.Search(
		es.Search.WithIndex("my-index-000001"),
		es.Search.WithQuery("user.id:kimchy"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:d2146e521344c5baa4984cf9147f18f2[]
}
