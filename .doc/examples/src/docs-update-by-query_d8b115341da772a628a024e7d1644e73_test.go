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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L334>
//
// --------------------------------------------------------------------------------
// POST twitter/_update_by_query?routing=1
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_d8b115341da772a628a024e7d1644e73(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d8b115341da772a628a024e7d1644e73[]
	res, err := es.UpdateByQuery(
		[]string{"twitter"},
		es.UpdateByQuery.WithRouting("1"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:d8b115341da772a628a024e7d1644e73[]
}
