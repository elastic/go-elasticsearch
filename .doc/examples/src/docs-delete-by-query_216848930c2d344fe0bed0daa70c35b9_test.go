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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L586>
//
// --------------------------------------------------------------------------------
// GET _tasks?detailed=true&actions=*/delete/byquery
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_216848930c2d344fe0bed0daa70c35b9(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:216848930c2d344fe0bed0daa70c35b9[]
	res, err := es.Tasks.Get(
		"?detailed=true&actions=*/delete/byquery",
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:216848930c2d344fe0bed0daa70c35b9[]
}
