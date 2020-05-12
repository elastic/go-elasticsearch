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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/nested-query.asciidoc#L133>
//
// --------------------------------------------------------------------------------
// PUT /drivers
// {
//     "mappings" : {
//         "properties" : {
//             "driver" : {
//                 "type" : "nested",
//                 "properties" : {
//                     "last_name" : {
//                         "type" : "text"
//                     },
//                     "vehicle" : {
//                         "type" : "nested",
//                         "properties" : {
//                             "make" : {
//                                 "type" : "text"
//                             },
//                             "model" : {
//                                 "type" : "text"
//                             }
//                         }
//                     }
//                 }
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_nested_query_54092c8c646133f5dbbc047990dd458d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:54092c8c646133f5dbbc047990dd458d[]
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
	// end:54092c8c646133f5dbbc047990dd458d[]
}
