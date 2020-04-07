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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L10>
//
// --------------------------------------------------------------------------------
// GET twitter/_doc/0
// --------------------------------------------------------------------------------

func Test_docs_get_fbcf5078a6a9e09790553804054c36b3(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:fbcf5078a6a9e09790553804054c36b3[]
	res, err := es.Get("twitter", "0", es.Get.WithPretty())
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:fbcf5078a6a9e09790553804054c36b3[]
}
