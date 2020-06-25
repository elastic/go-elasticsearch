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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/get-mapping.asciidoc#L70>
//
// --------------------------------------------------------------------------------
// GET /_all/_mapping
//
// GET /_mapping
// --------------------------------------------------------------------------------

func Test_indices_get_mapping_09cdd5ae8114c49886026fef8d00a19c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:09cdd5ae8114c49886026fef8d00a19c[]
	{
		res, err := es.Indices.GetMapping(es.Indices.GetMapping.WithIndex("_all"))
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Indices.GetMapping()
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:09cdd5ae8114c49886026fef8d00a19c[]
}
