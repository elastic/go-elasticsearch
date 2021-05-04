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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L662>
//
// --------------------------------------------------------------------------------
// PUT test
// {
//   "mappings": {
//     "dynamic": false,   <1>
//     "properties": {
//       "text": {"type": "text"}
//     }
//   }
// }
//
// POST test/_doc?refresh
// {
//   "text": "words words",
//   "flag": "bar"
// }
// POST test/_doc?refresh
// {
//   "text": "words words",
//   "flag": "foo"
// }
// PUT test/_mapping   <2>
// {
//   "properties": {
//     "text": {"type": "text"},
//     "flag": {"type": "text", "analyzer": "keyword"}
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_2fe28d9a91b3081a9ec4601af8fb7b1c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:2fe28d9a91b3081a9ec4601af8fb7b1c[]
	{
		res, err := es.Indices.Create(
			"test",
			es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "dynamic": false,
		    "properties": {
		      "text": {
		        "type": "text"
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
	}

	{
		res, err := es.Index(
			"test",
			strings.NewReader(`{
		  "text": "words words",
		  "flag": "bar"
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
			"test",
			strings.NewReader(`{
		  "text": "words words",
		  "flag": "foo"
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
		res, err := es.Indices.PutMapping(
			[]string{"test"},
			strings.NewReader(`{
		  "properties": {
		    "text": {
		      "type": "text"
		    },
		    "flag": {
		      "type": "text",
		      "analyzer": "keyword"
		    }
		  }
		}`),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:2fe28d9a91b3081a9ec4601af8fb7b1c[]
}
