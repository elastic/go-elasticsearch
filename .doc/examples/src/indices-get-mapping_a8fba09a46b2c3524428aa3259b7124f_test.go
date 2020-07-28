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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/get-mapping.asciidoc#L11>
//
// --------------------------------------------------------------------------------
// GET /twitter/_mapping
// --------------------------------------------------------------------------------

func Test_indices_get_mapping_a8fba09a46b2c3524428aa3259b7124f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:a8fba09a46b2c3524428aa3259b7124f[]
	res, err := es.Indices.GetMapping(es.Indices.GetMapping.WithIndex("twitter"))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:a8fba09a46b2c3524428aa3259b7124f[]
}
