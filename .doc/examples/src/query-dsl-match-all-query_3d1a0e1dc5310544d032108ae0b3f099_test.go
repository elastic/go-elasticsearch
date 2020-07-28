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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/match-all-query.asciidoc#L23>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "match_all": { "boost" : 1.2 }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_match_all_query_3d1a0e1dc5310544d032108ae0b3f099(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3d1a0e1dc5310544d032108ae0b3f099[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match_all": {
		      "boost": 1.2
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
	// end:3d1a0e1dc5310544d032108ae0b3f099[]
}
