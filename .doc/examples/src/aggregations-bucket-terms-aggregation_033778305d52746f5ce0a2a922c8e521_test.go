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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L410>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "aggs" : {
//         "genres" : {
//             "terms" : {
//                 "script" : {
//                     "source": "doc['genre'].value",
//                     "lang": "painless"
//                 }
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_033778305d52746f5ce0a2a922c8e521(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:033778305d52746f5ce0a2a922c8e521[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "genres": {
		      "terms": {
		        "script": {
		          "source": "doc['genre'].value",
		          "lang": "painless"
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
	// end:033778305d52746f5ce0a2a922c8e521[]
}
