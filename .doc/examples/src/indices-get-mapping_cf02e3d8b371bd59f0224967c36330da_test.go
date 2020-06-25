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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/get-mapping.asciidoc#L60>
//
// --------------------------------------------------------------------------------
// GET /twitter,kimchy/_mapping
// --------------------------------------------------------------------------------

func Test_indices_get_mapping_cf02e3d8b371bd59f0224967c36330da(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:cf02e3d8b371bd59f0224967c36330da[]
	res, err := es.Indices.GetMapping(es.Indices.GetMapping.WithIndex("twitter,kimchy"))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:cf02e3d8b371bd59f0224967c36330da[]
}
