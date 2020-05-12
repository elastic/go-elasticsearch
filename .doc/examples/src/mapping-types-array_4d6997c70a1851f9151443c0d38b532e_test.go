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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/array.asciidoc#L42>
//
// --------------------------------------------------------------------------------
// PUT my_index/_doc/1
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
// PUT my_index/_doc/2 <3>
// {
//   "message": "no arrays in this document...",
//   "tags":  "elasticsearch",
//   "lists": {
//     "name": "prog_list",
//     "description": "programming list"
//   }
// }
//
// GET my_index/_search
// {
//   "query": {
//     "match": {
//       "tags": "elasticsearch" <4>
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_array_4d6997c70a1851f9151443c0d38b532e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4d6997c70a1851f9151443c0d38b532e[]
	{
		res, err := es.Index(
			"my_index",
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
			"my_index",
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
			es.Search.WithIndex("my_index"),
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
	// end:4d6997c70a1851f9151443c0d38b532e[]
}
