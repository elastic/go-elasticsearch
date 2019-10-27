// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/create-index.asciidoc#L99>
//
// --------------------------------------------------------------------------------
// PUT /twitter
// {
//     "settings" : {
//         "number_of_shards" : 3,
//         "number_of_replicas" : 2
//     }
// }
// --------------------------------------------------------------------------------

func Test_indices_create_index_b9c5d7ca6ca9c6f747201f45337a4abf(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b9c5d7ca6ca9c6f747201f45337a4abf[]
	res, err := es.Indices.Create(
		"twitter",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "settings": {
		    "number_of_shards": 3,
		    "number_of_replicas": 2
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:b9c5d7ca6ca9c6f747201f45337a4abf[]
}
