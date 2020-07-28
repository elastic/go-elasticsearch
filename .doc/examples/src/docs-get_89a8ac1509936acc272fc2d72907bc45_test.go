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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L273>
//
// --------------------------------------------------------------------------------
// GET twitter/_source/1
// --------------------------------------------------------------------------------

func Test_docs_get_89a8ac1509936acc272fc2d72907bc45(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:89a8ac1509936acc272fc2d72907bc45[]
	res, err := es.GetSource("twitter", "1", es.GetSource.WithPretty())
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:89a8ac1509936acc272fc2d72907bc45[]
}
