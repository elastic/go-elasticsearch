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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L599>
//
// --------------------------------------------------------------------------------
// POST twitter/_search?size=0&q=extra:test&filter_path=hits.total
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_025b54db0edc50c24ea48a2bd94366ad(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:025b54db0edc50c24ea48a2bd94366ad[]
	res, err := es.Search(
		es.Search.WithIndex("twitter"),
		es.Search.WithFilterPath("hits.total"),
		es.Search.WithQuery("extra:test"),
		es.Search.WithSize(0),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:025b54db0edc50c24ea48a2bd94366ad[]
}
