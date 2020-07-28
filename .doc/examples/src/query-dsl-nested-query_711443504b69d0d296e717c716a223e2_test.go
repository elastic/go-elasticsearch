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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/nested-query.asciidoc#L206>
//
// --------------------------------------------------------------------------------
// GET /drivers/_search
// {
//   "query": {
//     "nested": {
//       "path": "driver",
//       "query": {
//         "nested": {
//           "path": "driver.vehicle",
//           "query": {
//             "bool": {
//               "must": [
//                 { "match": { "driver.vehicle.make": "Powell Motors" } },
//                 { "match": { "driver.vehicle.model": "Canyonero" } }
//               ]
//             }
//           }
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_nested_query_711443504b69d0d296e717c716a223e2(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:711443504b69d0d296e717c716a223e2[]
	res, err := es.Search(
		es.Search.WithIndex("drivers"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "nested": {
		      "path": "driver",
		      "query": {
		        "nested": {
		          "path": "driver.vehicle",
		          "query": {
		            "bool": {
		              "must": [
		                {
		                  "match": {
		                    "driver.vehicle.make": "Powell Motors"
		                  }
		                },
		                {
		                  "match": {
		                    "driver.vehicle.model": "Canyonero"
		                  }
		                }
		              ]
		            }
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
	// end:711443504b69d0d296e717c716a223e2[]
}
