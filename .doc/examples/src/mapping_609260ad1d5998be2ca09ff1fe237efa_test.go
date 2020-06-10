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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping.asciidoc#L214>
//
// --------------------------------------------------------------------------------
// GET /my-index/_mapping
// --------------------------------------------------------------------------------

func Test_mapping_609260ad1d5998be2ca09ff1fe237efa(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:609260ad1d5998be2ca09ff1fe237efa[]
	res, err := es.Indices.GetMapping(es.Indices.GetMapping.WithIndex("my-index"))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:609260ad1d5998be2ca09ff1fe237efa[]
}
