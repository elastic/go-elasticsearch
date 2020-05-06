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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/query-string-query.asciidoc#L528>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "query_string": {
//             "fields": [
//                 "title",
//                 "content"
//             ],
//             "query": "this OR that OR thus",
//             "type": "cross_fields",
//             "minimum_should_match": 2
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_query_string_query_6f21a878fee3b43c5332b81aaddbeac7(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6f21a878fee3b43c5332b81aaddbeac7[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "query_string": {
		      "fields": [
		        "title",
		        "content"
		      ],
		      "query": "this OR that OR thus",
		      "type": "cross_fields",
		      "minimum_should_match": 2
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:6f21a878fee3b43c5332b81aaddbeac7[]
}
