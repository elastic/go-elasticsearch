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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/nested.asciidoc#L80>
//
// --------------------------------------------------------------------------------
// PUT my_index
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
// PUT my_index/_doc/1
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
// GET my_index/_search
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
// GET my_index/_search
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

func Test_mapping_types_nested_b919f88e6f47a40d5793479440a90ba6(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b919f88e6f47a40d5793479440a90ba6[]
	{
		res, err := es.Indices.Create(
			"my_index",
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
			"my_index",
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
			es.Search.WithIndex("my_index"),
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
			es.Search.WithIndex("my_index"),
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
	// end:b919f88e6f47a40d5793479440a90ba6[]
}
