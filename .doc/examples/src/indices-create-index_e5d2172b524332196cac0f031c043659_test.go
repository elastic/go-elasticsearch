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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/create-index.asciidoc#L81>
//
// --------------------------------------------------------------------------------
// PUT /twitter
// {
//     "settings" : {
//         "index" : {
//             "number_of_shards" : 3, <1>
//             "number_of_replicas" : 2 <2>
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_indices_create_index_e5d2172b524332196cac0f031c043659(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:e5d2172b524332196cac0f031c043659[]
	res, err := es.Indices.Create(
		"twitter",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "settings": {
		    "index": {
		      "number_of_shards": 3,
		      "number_of_replicas": 2
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:e5d2172b524332196cac0f031c043659[]
}
