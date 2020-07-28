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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/get-mapping.asciidoc#L78>
//
// --------------------------------------------------------------------------------
// GET /*/_mapping
//
// GET /_all/_mapping
//
// GET /_mapping
// --------------------------------------------------------------------------------

func Test_indices_get_mapping_5b7d6f1db88ca6f42c48fa3dbb4341e8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5b7d6f1db88ca6f42c48fa3dbb4341e8[]
	{
		res, err := es.Indices.GetMapping(es.Indices.GetMapping.WithIndex("*"))
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

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
	// end:5b7d6f1db88ca6f42c48fa3dbb4341e8[]
}
