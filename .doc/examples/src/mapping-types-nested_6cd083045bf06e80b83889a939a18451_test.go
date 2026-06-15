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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/nested.asciidoc#L85>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "user": {
//         "type": "nested" <1>
//       }
//     }
//   }
// }
//
// PUT my-index-000001/_doc/1
// {
//   "group" : "fans",
//   "user" : [
//     {
//       "first" : "John",
//       "last" :  "Smith"
//     },
//     {
//       "first" : "Alice",
//       "last" :  "White"
//     }
//   ]
// }
//
// GET my-index-000001/_search
// {
//   "query": {
//     "nested": {
//       "path": "user",
//       "query": {
//         "bool": {
//           "must": [
//             { "match": { "user.first": "Alice" }},
//             { "match": { "user.last":  "Smith" }} <2>
//           ]
//         }
//       }
//     }
//   }
// }
//
// GET my-index-000001/_search
// {
//   "query": {
//     "nested": {
//       "path": "user",
//       "query": {
//         "bool": {
//           "must": [
//             { "match": { "user.first": "Alice" }},
//             { "match": { "user.last":  "White" }} <3>
//           ]
//         }
//       },
//       "inner_hits": { <4>
//         "highlight": {
//           "fields": {
//             "user.first": {}
//           }
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_nested_6cd083045bf06e80b83889a939a18451(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6cd083045bf06e80b83889a939a18451[]
	{
		res, err := es.Indices.Create(
			"my-index-000001",
			es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "user": {
		        "type": "nested"
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
			"my-index-000001",
			strings.NewReader(`{
		  "group": "fans",
		  "user": [
		    {
		      "first": "John",
		      "last": "Smith"
		    },
		    {
		      "first": "Alice",
		      "last": "White"
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
		res, err := es.Search(
			es.Search.WithIndex("my-index-000001"),
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "nested": {
		      "path": "user",
		      "query": {
		        "bool": {
		          "must": [
		            {
		              "match": {
		                "user.first": "Alice"
		              }
		            },
		            {
		              "match": {
		                "user.last": "Smith"
		              }
		            }
		          ]
		        }
		      }
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

	{
		res, err := es.Search(
			es.Search.WithIndex("my-index-000001"),
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "nested": {
		      "path": "user",
		      "query": {
		        "bool": {
		          "must": [
		            {
		              "match": {
		                "user.first": "Alice"
		              }
		            },
		            {
		              "match": {
		                "user.last": "White"
		              }
		            }
		          ]
		        }
		      },
		      "inner_hits": {
		        "highlight": {
		          "fields": {
		            "user.first": {}
		          }
		        }
		      }
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
	// end:6cd083045bf06e80b83889a939a18451[]
}
