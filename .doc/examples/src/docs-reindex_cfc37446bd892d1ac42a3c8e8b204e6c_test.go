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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L718>
//
// --------------------------------------------------------------------------------
// GET test2/_doc/1
// --------------------------------------------------------------------------------

func Test_docs_reindex_cfc37446bd892d1ac42a3c8e8b204e6c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:cfc37446bd892d1ac42a3c8e8b204e6c[]
	res, err := es.Get("test2", "1", es.Get.WithPretty())
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:cfc37446bd892d1ac42a3c8e8b204e6c[]
}
