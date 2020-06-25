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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search.asciidoc#L96>
//
// --------------------------------------------------------------------------------
// POST /_search
// {
//     "query" : {
//         "match_all" : {}
//     },
//     "stats" : ["group1", "group2"]
// }
// --------------------------------------------------------------------------------

func Test_search_189a921df2f5b1fe580937210ce9c1c2(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:189a921df2f5b1fe580937210ce9c1c2[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match_all": {}
		  },
		  "stats": [
		    "group1",
		    "group2"
		  ]
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:189a921df2f5b1fe580937210ce9c1c2[]
}
