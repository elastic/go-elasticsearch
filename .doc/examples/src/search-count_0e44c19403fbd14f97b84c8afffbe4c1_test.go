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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/count.asciidoc#L99>
//
// --------------------------------------------------------------------------------
// PUT /twitter/_doc/1?refresh
// {
//   "user": "kimchy"
// }
//
// GET /twitter/_count?q=user:kimchy
//
// GET /twitter/_count
// {
//   "query" : {
//     "term" : { "user" : "kimchy" }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_count_0e44c19403fbd14f97b84c8afffbe4c1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0e44c19403fbd14f97b84c8afffbe4c1[]
	{
		res, err := es.Index(
			"twitter",
			strings.NewReader(`{
		  "user": "kimchy"
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
		res, err := es.Count(
			es.Count.WithIndex("?q=user:kimchy"),
			es.Count.WithQuery("user:kimchy"),
			es.Count.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Count(
			es.Count.WithIndex(""),
			es.Count.WithBody(strings.NewReader(`{
		  "query": {
		    "term": {
		      "user": "kimchy"
		    }
		  }
		}`)),
			es.Count.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:0e44c19403fbd14f97b84c8afffbe4c1[]
}
