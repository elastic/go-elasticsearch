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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L343>
//
// --------------------------------------------------------------------------------
// POST twitter/_update_by_query?scroll_size=100
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_54a770f053f3225ea0d1e34334232411(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:54a770f053f3225ea0d1e34334232411[]
	res, err := es.UpdateByQuery(
		[]string{"twitter"},
		es.UpdateByQuery.WithScrollSize(100),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:54a770f053f3225ea0d1e34334232411[]
}
