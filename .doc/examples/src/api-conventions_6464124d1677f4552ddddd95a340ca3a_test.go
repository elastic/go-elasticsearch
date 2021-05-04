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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L353>
//
// --------------------------------------------------------------------------------
// POST /library/_doc?refresh
// {"title": "Book #1", "rating": 200.1}
// POST /library/_doc?refresh
// {"title": "Book #2", "rating": 1.7}
// POST /library/_doc?refresh
// {"title": "Book #3", "rating": 0.1}
// GET /_search?filter_path=hits.hits._source&_source=title&sort=rating:desc
// --------------------------------------------------------------------------------

func Test_api_conventions_6464124d1677f4552ddddd95a340ca3a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6464124d1677f4552ddddd95a340ca3a[]
	{
		res, err := es.Index(
			"library",
			strings.NewReader(`{
		  "title": "Book #1",
		  "rating": 200.1
		}`),
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
		res, err := es.Index(
			"library",
			strings.NewReader(`{
		  "title": "Book #2",
		  "rating": 1.7
		}`),
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
		res, err := es.Index(
			"library",
			strings.NewReader(`{
		  "title": "Book #3",
		  "rating": 0.1
		}`),
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
			es.Search.WithSource("title"),
			es.Search.WithFilterPath("hits.hits._source"),
			es.Search.WithSort("rating:desc"),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:6464124d1677f4552ddddd95a340ca3a[]
}
