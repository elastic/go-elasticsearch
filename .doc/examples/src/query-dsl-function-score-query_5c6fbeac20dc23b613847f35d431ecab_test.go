// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
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
//   "query": {
//     "function_score": {
//       "functions": [
//         {
//           "gauss": {
//             "price": {
//               "origin": "0",
//               "scale": "20"
//             }
//           }
//         },
//         {
//           "gauss": {
//             "location": {
//               "origin": "11, 12",
//               "scale": "2km"
//             }
//           }
//         }
//       ],
//       "query": {
//         "match": {
//           "properties": "balcony"
//         }
//       },
//       "score_mode": "multiply"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_function_score_query_5c6fbeac20dc23b613847f35d431ecab(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5c6fbeac20dc23b613847f35d431ecab[]
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
	// end:5c6fbeac20dc23b613847f35d431ecab[]
}
