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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/get-index.asciidoc#L10>
//
// --------------------------------------------------------------------------------
// GET /twitter
// --------------------------------------------------------------------------------

func Test_indices_get_index_be8f28f31207b173de61be032fcf239c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:be8f28f31207b173de61be032fcf239c[]
	res, err := es.Indices.Get([]string{"twitter"})
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:be8f28f31207b173de61be032fcf239c[]
}
