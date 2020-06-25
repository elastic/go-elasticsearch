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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L413>
//
// --------------------------------------------------------------------------------
// GET _tasks?detailed=true&actions=*byquery
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_7df191cc7f814e410a4ac7261065e6ef(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:7df191cc7f814e410a4ac7261065e6ef[]
	res, err := es.Tasks.List(
		es.Tasks.List.WithActions("*byquery"),
		es.Tasks.List.WithDetailed(true),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:7df191cc7f814e410a4ac7261065e6ef[]
}
