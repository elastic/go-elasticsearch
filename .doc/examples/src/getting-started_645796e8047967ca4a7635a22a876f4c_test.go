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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/getting-started.asciidoc#L691>
//
// --------------------------------------------------------------------------------
// GET /bank/_search
// {
//   "size": 0,
//   "aggs": {
//     "group_by_state": {
//       "terms": {
//         "field": "state.keyword",
//         "order": {
//           "average_balance": "desc"
//         }
//       },
//       "aggs": {
//         "average_balance": {
//           "avg": {
//             "field": "balance"
//           }
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_getting_started_645796e8047967ca4a7635a22a876f4c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:645796e8047967ca4a7635a22a876f4c[]
	res, err := es.Search(
		es.Search.WithIndex("bank"),
		es.Search.WithBody(strings.NewReader(`{
		  "size": 0,
		  "aggs": {
		    "group_by_state": {
		      "terms": {
		        "field": "state.keyword",
		        "order": {
		          "average_balance": "desc"
		        }
		      },
		      "aggs": {
		        "average_balance": {
		          "avg": {
		            "field": "balance"
		          }
		        }
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
	// end:645796e8047967ca4a7635a22a876f4c[]
}
