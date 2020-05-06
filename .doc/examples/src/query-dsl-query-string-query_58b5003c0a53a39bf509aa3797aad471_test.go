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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/query-string-query.asciidoc#L352>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "query_string" : {
//             "fields" : ["content", "name.*^5"],
//             "query" : "this AND that OR thus"
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_query_string_query_58b5003c0a53a39bf509aa3797aad471(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:58b5003c0a53a39bf509aa3797aad471[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "query_string": {
		      "fields": [
		        "content",
		        "name.*^5"
		      ],
		      "query": "this AND that OR thus"
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
	// end:58b5003c0a53a39bf509aa3797aad471[]
}
