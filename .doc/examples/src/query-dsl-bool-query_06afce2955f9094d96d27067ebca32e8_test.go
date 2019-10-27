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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/bool-query.asciidoc#L36>
//
// --------------------------------------------------------------------------------
// POST _search
// {
//   "query": {
//     "bool" : {
//       "must" : {
//         "term" : { "user" : "kimchy" }
//       },
//       "filter": {
//         "term" : { "tag" : "tech" }
//       },
//       "must_not" : {
//         "range" : {
//           "age" : { "gte" : 10, "lte" : 20 }
//         }
//       },
//       "should" : [
//         { "term" : { "tag" : "wow" } },
//         { "term" : { "tag" : "elasticsearch" } }
//       ],
//       "minimum_should_match" : 1,
//       "boost" : 1.0
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_bool_query_06afce2955f9094d96d27067ebca32e8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:06afce2955f9094d96d27067ebca32e8[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "bool": {
		      "must": {
		        "term": {
		          "user": "kimchy"
		        }
		      },
		      "filter": {
		        "term": {
		          "tag": "tech"
		        }
		      },
		      "must_not": {
		        "range": {
		          "age": {
		            "gte": 10,
		            "lte": 20
		          }
		        }
		      },
		      "should": [
		        {
		          "term": {
		            "tag": "wow"
		          }
		        },
		        {
		          "term": {
		            "tag": "elasticsearch"
		          }
		        }
		      ],
		      "minimum_should_match": 1,
		      "boost": 1.0
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
	// end:06afce2955f9094d96d27067ebca32e8[]
}
