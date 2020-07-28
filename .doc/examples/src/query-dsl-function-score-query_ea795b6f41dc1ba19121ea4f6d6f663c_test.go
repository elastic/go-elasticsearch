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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/function-score-query.asciidoc#L175>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "function_score": {
//       "query": {
//         "match": { "message": "elasticsearch" }
//       },
//       "script_score": {
//         "script": {
//           "params": {
//             "a": 5,
//             "b": 1.2
//           },
//           "source": "params.a / Math.pow(params.b, doc['likes'].value)"
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_function_score_query_ea795b6f41dc1ba19121ea4f6d6f663c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:ea795b6f41dc1ba19121ea4f6d6f663c[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "function_score": {
		      "query": {
		        "match": {
		          "message": "elasticsearch"
		        }
		      },
		      "script_score": {
		        "script": {
		          "params": {
		            "a": 5,
		            "b": 1.2
		          },
		          "source": "params.a / Math.pow(params.b, doc['likes'].value)"
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
	// end:ea795b6f41dc1ba19121ea4f6d6f663c[]
}
