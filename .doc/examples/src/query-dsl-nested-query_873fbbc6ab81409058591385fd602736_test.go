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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/nested-query.asciidoc#L165>
//
// --------------------------------------------------------------------------------
// PUT /drivers/_doc/1
// {
//   "driver" : {
//         "last_name" : "McQueen",
//         "vehicle" : [
//             {
//                 "make" : "Powell Motors",
//                 "model" : "Canyonero"
//             },
//             {
//                 "make" : "Miller-Meteor",
//                 "model" : "Ecto-1"
//             }
//         ]
//     }
// }
//
// PUT /drivers/_doc/2?refresh
// {
//   "driver" : {
//         "last_name" : "Hudson",
//         "vehicle" : [
//             {
//                 "make" : "Mifune",
//                 "model" : "Mach Five"
//             },
//             {
//                 "make" : "Miller-Meteor",
//                 "model" : "Ecto-1"
//             }
//         ]
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_nested_query_873fbbc6ab81409058591385fd602736(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:873fbbc6ab81409058591385fd602736[]
	{
		res, err := es.Index(
			"drivers",
			strings.NewReader(`{
		  "driver": {
		    "last_name": "McQueen",
		    "vehicle": [
		      {
		        "make": "Powell Motors",
		        "model": "Canyonero"
		      },
		      {
		        "make": "Miller-Meteor",
		        "model": "Ecto-1"
		      }
		    ]
		  }
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Index(
			"drivers",
			strings.NewReader(`{
		  "driver": {
		    "last_name": "Hudson",
		    "vehicle": [
		      {
		        "make": "Mifune",
		        "model": "Mach Five"
		      },
		      {
		        "make": "Miller-Meteor",
		        "model": "Ecto-1"
		      }
		    ]
		  }
		}`),
			es.Index.WithDocumentID("2"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:873fbbc6ab81409058591385fd602736[]
}
