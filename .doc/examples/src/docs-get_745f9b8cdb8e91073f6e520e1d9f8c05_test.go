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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L73>
//
// --------------------------------------------------------------------------------
// GET twitter/_doc/0?_source=*.id,retweeted
// --------------------------------------------------------------------------------

func Test_docs_get_745f9b8cdb8e91073f6e520e1d9f8c05(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:745f9b8cdb8e91073f6e520e1d9f8c05[]
	res, err := es.Get(
		"twitter",
		"0",
		es.Get.WithSource("*.id,retweeted"),
		es.Get.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:745f9b8cdb8e91073f6e520e1d9f8c05[]
}
