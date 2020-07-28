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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping.asciidoc#L217>
//
// --------------------------------------------------------------------------------
// GET /my-index-000001/_mapping
// --------------------------------------------------------------------------------

func Test_mapping_a6fdd0100cd362df54af6c95d1055c96(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:a6fdd0100cd362df54af6c95d1055c96[]
	res, err := es.Indices.GetMapping(es.Indices.GetMapping.WithIndex("my-index-000001"))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:a6fdd0100cd362df54af6c95d1055c96[]
}
