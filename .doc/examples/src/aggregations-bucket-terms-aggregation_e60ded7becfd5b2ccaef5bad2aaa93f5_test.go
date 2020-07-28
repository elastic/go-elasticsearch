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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L162>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "aggs": {
//     "products": {
//       "terms": {
//         "field": "product",
//         "size": 5,
//         "show_term_doc_count_error": true
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_e60ded7becfd5b2ccaef5bad2aaa93f5(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:e60ded7becfd5b2ccaef5bad2aaa93f5[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "products": {
		      "terms": {
		        "field": "product",
		        "size": 5,
		        "show_term_doc_count_error": true
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
	// end:e60ded7becfd5b2ccaef5bad2aaa93f5[]
}
