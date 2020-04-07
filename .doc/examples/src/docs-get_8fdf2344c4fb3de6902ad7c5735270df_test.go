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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L65>
//
// --------------------------------------------------------------------------------
// GET twitter/_doc/0?_source_includes=*.id&_source_excludes=entities
// --------------------------------------------------------------------------------

func Test_docs_get_8fdf2344c4fb3de6902ad7c5735270df(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8fdf2344c4fb3de6902ad7c5735270df[]
	res, err := es.Get(
		"twitter",
		"0",
		es.Get.WithSourceExcludes("entities"),
		es.Get.WithSourceIncludes("*.id"),
		es.Get.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:8fdf2344c4fb3de6902ad7c5735270df[]
}
