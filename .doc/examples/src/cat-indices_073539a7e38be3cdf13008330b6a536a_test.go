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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/cat/indices.asciidoc#L100>
//
// --------------------------------------------------------------------------------
// GET /_cat/indices/twi*?v&s=index
// --------------------------------------------------------------------------------

func Test_cat_indices_073539a7e38be3cdf13008330b6a536a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:073539a7e38be3cdf13008330b6a536a[]
	res, err := es.Cat.Indices(
		es.Cat.Indices.WithIndex([]string{"twi*"}...),
		es.Cat.Indices.WithS("index"),
		es.Cat.Indices.WithV(true),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:073539a7e38be3cdf13008330b6a536a[]
}
