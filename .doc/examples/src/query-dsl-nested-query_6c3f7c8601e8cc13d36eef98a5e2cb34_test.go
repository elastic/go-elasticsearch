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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/nested-query.asciidoc#L133>
//
// --------------------------------------------------------------------------------
// PUT /drivers
// {
//   "mappings": {
//     "properties": {
//       "driver": {
//         "type": "nested",
//         "properties": {
//           "last_name": {
//             "type": "text"
//           },
//           "vehicle": {
//             "type": "nested",
//             "properties": {
//               "make": {
//                 "type": "text"
//               },
//               "model": {
//                 "type": "text"
//               }
//             }
//           }
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_nested_query_6c3f7c8601e8cc13d36eef98a5e2cb34(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6c3f7c8601e8cc13d36eef98a5e2cb34[]
	res, err := es.Indices.Create(
		"drivers",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "driver": {
		        "type": "nested",
		        "properties": {
		          "last_name": {
		            "type": "text"
		          },
		          "vehicle": {
		            "type": "nested",
		            "properties": {
		              "make": {
		                "type": "text"
		              },
		              "model": {
		                "type": "text"
		              }
		            }
		          }
		        }
		      }
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:6c3f7c8601e8cc13d36eef98a5e2cb34[]
}
