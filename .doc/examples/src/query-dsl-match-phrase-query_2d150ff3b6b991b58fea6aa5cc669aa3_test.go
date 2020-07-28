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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/match-phrase-query.asciidoc#L30>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "match_phrase": {
//       "message": {
//         "query": "this is a test",
//         "analyzer": "my_analyzer"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_match_phrase_query_2d150ff3b6b991b58fea6aa5cc669aa3(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:2d150ff3b6b991b58fea6aa5cc669aa3[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match_phrase": {
		      "message": {
		        "query": "this is a test",
		        "analyzer": "my_analyzer"
		      }
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
	// end:2d150ff3b6b991b58fea6aa5cc669aa3[]
}
