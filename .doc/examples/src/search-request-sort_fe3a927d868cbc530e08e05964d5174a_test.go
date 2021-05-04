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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L94>
//
// --------------------------------------------------------------------------------
// PUT /my-index-000001/_doc/1?refresh
// {
//    "product": "chocolate",
//    "price": [20, 4]
// }
//
// POST /_search
// {
//    "query" : {
//       "term" : { "product" : "chocolate" }
//    },
//    "sort" : [
//       {"price" : {"order" : "asc", "mode" : "avg"}}
//    ]
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_fe3a927d868cbc530e08e05964d5174a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:fe3a927d868cbc530e08e05964d5174a[]
	{
		res, err := es.Index(
			"my-index-000001",
			strings.NewReader(`{
		  "product": "chocolate",
		  "price": [
		    20,
		    4
		  ]
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "product": "chocolate"
		    }
		  },
		  "sort": [
		    {
		      "price": {
		        "order": "asc",
		        "mode": "avg"
		      }
		    }
		  ]
		}`)),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:fe3a927d868cbc530e08e05964d5174a[]
}
