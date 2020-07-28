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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L339>
//
// --------------------------------------------------------------------------------
// GET twitter/_doc/1?stored_fields=tags,counter
// --------------------------------------------------------------------------------

func Test_docs_get_710c7871f20f176d51209b1574b0d61b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:710c7871f20f176d51209b1574b0d61b[]
	res, err := es.Get(
		"twitter",
		"1",
		es.Get.WithStoredFields("tags,counter"),
		es.Get.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:710c7871f20f176d51209b1574b0d61b[]
}
