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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L95>
//
// --------------------------------------------------------------------------------
// PUT /publications
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_12433d2b637d002e8d5c9a1adce69d3b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:12433d2b637d002e8d5c9a1adce69d3b[]
	res, err := es.Indices.Create("publications")
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:12433d2b637d002e8d5c9a1adce69d3b[]
}
