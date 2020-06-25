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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete.asciidoc#L75>
//
// --------------------------------------------------------------------------------
// DELETE /twitter/_doc/1?routing=kimchy
// --------------------------------------------------------------------------------

func Test_docs_delete_47b5ff897f26e9c943cee5c06034181d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:47b5ff897f26e9c943cee5c06034181d[]
	res, err := es.Delete("twitter",
		"1",
		es.Delete.WithRouting("kimchy"),
		es.Delete.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:47b5ff897f26e9c943cee5c06034181d[]
}
