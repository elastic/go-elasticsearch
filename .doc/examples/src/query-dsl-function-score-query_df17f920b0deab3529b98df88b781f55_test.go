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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/function-score-query.asciidoc#L578>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "function_score": {
//           "functions": [
//             {
//               "gauss": {
//                 "price": {
//                   "origin": "0",
//                   "scale": "20"
//                 }
//               }
//             },
//             {
//               "gauss": {
//                 "location": {
//                   "origin": "11, 12",
//                   "scale": "2km"
//                 }
//               }
//             }
//           ],
//           "query": {
//             "match": {
//               "properties": "balcony"
//             }
//           },
//           "score_mode": "multiply"
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_function_score_query_df17f920b0deab3529b98df88b781f55(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:df17f920b0deab3529b98df88b781f55[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "function_score": {
		      "functions": [
		        {
		          "gauss": {
		            "price": {
		              "origin": "0",
		              "scale": "20"
		            }
		          }
		        },
		        {
		          "gauss": {
		            "location": {
		              "origin": "11, 12",
		              "scale": "2km"
		            }
		          }
		        }
		      ],
		      "query": {
		        "match": {
		          "properties": "balcony"
		        }
		      },
		      "score_mode": "multiply"
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
	// end:df17f920b0deab3529b98df88b781f55[]
}
