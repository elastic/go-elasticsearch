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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping.asciidoc#L250>
//
// --------------------------------------------------------------------------------
// GET /my-index/_mapping/field/employee-id
// --------------------------------------------------------------------------------

func Test_mapping_99a52be903945b17e734a1d02a57e958(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:99a52be903945b17e734a1d02a57e958[]
	res, err := es.Indices.GetMapping(es.Indices.GetMapping.WithIndex("my-index"))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:99a52be903945b17e734a1d02a57e958[]
}
