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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L288>
//
// --------------------------------------------------------------------------------
// HEAD twitter/_source/1
// --------------------------------------------------------------------------------

func Test_docs_get_2468ab381257d759d8a88af1141f6f9c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:2468ab381257d759d8a88af1141f6f9c[]
	res, err := es.ExistsSource("twitter", "_source", es.ExistsSource.WithPretty())
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:2468ab381257d759d8a88af1141f6f9c[]
}
