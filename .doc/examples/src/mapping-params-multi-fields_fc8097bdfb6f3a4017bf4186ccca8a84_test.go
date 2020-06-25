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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/params/multi-fields.asciidoc#L75>
//
// --------------------------------------------------------------------------------
// PUT my_index
// {
//   "mappings": {
//     "properties": {
//       "text": { <1>
//         "type": "text",
//         "fields": {
//           "english": { <2>
//             "type":     "text",
//             "analyzer": "english"
//           }
//         }
//       }
//     }
//   }
// }
//
// PUT my_index/_doc/1
// { "text": "quick brown fox" } <3>
//
// PUT my_index/_doc/2
// { "text": "quick brown foxes" } <3>
//
// GET my_index/_search
// {
//   "query": {
//     "multi_match": {
//       "query": "quick brown foxes",
//       "fields": [ <4>
//         "text",
//         "text.english"
//       ],
//       "type": "most_fields" <4>
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_params_multi_fields_fc8097bdfb6f3a4017bf4186ccca8a84(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:fc8097bdfb6f3a4017bf4186ccca8a84[]
	{
		res, err := es.Indices.Create(
			"my_index",
			es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "text": {
		        "type": "text",
		        "fields": {
		          "english": {
		            "type": "text",
		            "analyzer": "english"
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
	}

	{
		res, err := es.Index(
			"my_index",
			strings.NewReader(`{
		  "text": "quick brown fox"
		} `),
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
		  "text": "quick brown foxes"
		} `),
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
		    "multi_match": {
		      "query": "quick brown foxes",
		      "fields": [
		        "text",
		        "text.english"
		      ],
		      "type": "most_fields"
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
	// end:fc8097bdfb6f3a4017bf4186ccca8a84[]
}
