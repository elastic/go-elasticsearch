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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/getting-started.asciidoc#L544>
//
// --------------------------------------------------------------------------------
// GET /bank/_search
// {
//   "query": {
//     "bool": {
//       "must": { "match_all": {} },
//       "filter": {
//         "range": {
//           "balance": {
//             "gte": 20000,
//             "lte": 30000
//           }
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_getting_started_251ea12c1248385ab409906ac64d9ee9(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:251ea12c1248385ab409906ac64d9ee9[]
	res, err := es.Search(
		es.Search.WithIndex("bank/"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "bool": {
		      "must": {
		        "match_all": {}
		      },
		      "filter": {
		        "range": {
		          "balance": {
		            "gte": 20000,
		            "lte": 30000
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
	// end:251ea12c1248385ab409906ac64d9ee9[]
}
