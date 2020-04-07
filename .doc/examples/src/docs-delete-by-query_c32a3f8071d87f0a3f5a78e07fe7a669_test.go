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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L376>
//
// --------------------------------------------------------------------------------
// POST twitter/_delete_by_query?routing=1
// {
//   "query": {
//     "range" : {
//         "age" : {
//            "gte" : 10
//         }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_c32a3f8071d87f0a3f5a78e07fe7a669(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:c32a3f8071d87f0a3f5a78e07fe7a669[]
	res, err := es.DeleteByQuery(
		[]string{"twitter"},
		strings.NewReader(`{
		  "query": {
		    "range": {
		      "age": {
		        "gte": 10
		      }
		    }
		  }
		}`),
		es.DeleteByQuery.WithRouting("1"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:c32a3f8071d87f0a3f5a78e07fe7a669[]
}
