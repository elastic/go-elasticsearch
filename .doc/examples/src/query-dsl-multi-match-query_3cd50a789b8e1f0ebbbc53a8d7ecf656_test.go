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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/multi-match-query.asciidoc#L438>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "bool": {
//       "should": [
//         {
//           "multi_match" : {
//             "query":      "Will Smith",
//             "type":       "cross_fields",
//             "fields":     [ "first", "last" ],
//             "minimum_should_match": "50%" <1>
//           }
//         },
//         {
//           "multi_match" : {
//             "query":      "Will Smith",
//             "type":       "cross_fields",
//             "fields":     [ "*.edge" ]
//           }
//         }
//       ]
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_multi_match_query_3cd50a789b8e1f0ebbbc53a8d7ecf656(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3cd50a789b8e1f0ebbbc53a8d7ecf656[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "bool": {
		      "should": [
		        {
		          "multi_match": {
		            "query": "Will Smith",
		            "type": "cross_fields",
		            "fields": [
		              "first",
		              "last"
		            ],
		            "minimum_should_match": "50%"
		          }
		        },
		        {
		          "multi_match": {
		            "query": "Will Smith",
		            "type": "cross_fields",
		            "fields": [
		              "*.edge"
		            ]
		          }
		        }
		      ]
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
	// end:3cd50a789b8e1f0ebbbc53a8d7ecf656[]
}
