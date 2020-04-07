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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/exists-query.asciidoc#L20>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "exists": {
//             "field": "user"
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_exists_query_3342c69b2c2303247217532956fcce85(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3342c69b2c2303247217532956fcce85[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "exists": {
		      "field": "user"
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
	// end:3342c69b2c2303247217532956fcce85[]
}
