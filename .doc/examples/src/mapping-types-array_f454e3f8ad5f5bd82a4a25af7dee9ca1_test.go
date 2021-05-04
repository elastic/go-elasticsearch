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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/array.asciidoc#L42>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001/_doc/1
// {
//   "message": "some arrays in this document...",
//   "tags":  [ "elasticsearch", "wow" ], <1>
//   "lists": [ <2>
//     {
//       "name": "prog_list",
//       "description": "programming list"
//     },
//     {
//       "name": "cool_list",
//       "description": "cool stuff list"
//     }
//   ]
// }
//
// PUT my-index-000001/_doc/2 <3>
// {
//   "message": "no arrays in this document...",
//   "tags":  "elasticsearch",
//   "lists": {
//     "name": "prog_list",
//     "description": "programming list"
//   }
// }
//
// GET my-index-000001/_search
// {
//   "query": {
//     "match": {
//       "tags": "elasticsearch" <4>
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_array_f454e3f8ad5f5bd82a4a25af7dee9ca1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:f454e3f8ad5f5bd82a4a25af7dee9ca1[]
	{
		res, err := es.Index(
			"my-index-000001",
			strings.NewReader(`{
		  "message": "some arrays in this document...",
		  "tags": [
		    "elasticsearch",
		    "wow"
		  ],
		  "lists": [
		    {
		      "name": "prog_list",
		      "description": "programming list"
		    },
		    {
		      "name": "cool_list",
		      "description": "cool stuff list"
		    }
		  ]
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
			"my-index-000001",
			strings.NewReader(`{
		  "message": "no arrays in this document...",
		  "tags": "elasticsearch",
		  "lists": {
		    "name": "prog_list",
		    "description": "programming list"
		  }
		}`),
			es.Index.WithDocumentID("2"),
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
			es.Search.WithIndex("my-index-000001"),
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "tags": "elasticsearch"
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
	}
	// end:f454e3f8ad5f5bd82a4a25af7dee9ca1[]
}
