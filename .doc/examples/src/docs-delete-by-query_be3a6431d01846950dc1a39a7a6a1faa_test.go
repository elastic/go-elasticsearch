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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L640>
//
// --------------------------------------------------------------------------------
// GET /_tasks/r1A2WoRbTwKZ516z6NEs5A:36619
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_be3a6431d01846950dc1a39a7a6a1faa(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:be3a6431d01846950dc1a39a7a6a1faa[]
	res, err := es.Tasks.Get(
		"r1A2WoRbTwKZ516z6NEs5A:36619",
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:be3a6431d01846950dc1a39a7a6a1faa[]
}
