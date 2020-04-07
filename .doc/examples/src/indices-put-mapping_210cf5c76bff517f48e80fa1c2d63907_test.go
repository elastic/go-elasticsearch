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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L186>
//
// --------------------------------------------------------------------------------
// GET /my_index/_mapping
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_210cf5c76bff517f48e80fa1c2d63907(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:210cf5c76bff517f48e80fa1c2d63907[]
	res, err := es.Indices.GetMapping(es.Indices.GetMapping.WithIndex("my_index"))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:210cf5c76bff517f48e80fa1c2d63907[]
}
