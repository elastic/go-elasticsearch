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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L802>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "max_docs": 10,
//   "source": {
//     "index": "twitter",
//     "query": {
//       "function_score" : {
//         "random_score" : {},
//         "min_score" : 0.9    <1>
//       }
//     }
//   },
//   "dest": {
//     "index": "random_twitter"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_1bc731a4df952228af6dfa6b48627332(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1bc731a4df952228af6dfa6b48627332[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "max_docs": 10,
		  "source": {
		    "index": "twitter",
		    "query": {
		      "function_score": {
		        "random_score": {},
		        "min_score": 0.9
		      }
		    }
		  },
		  "dest": {
		    "index": "random_twitter"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1bc731a4df952228af6dfa6b48627332[]
}
