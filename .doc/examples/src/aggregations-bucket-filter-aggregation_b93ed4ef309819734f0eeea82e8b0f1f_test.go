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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/filter-aggregation.asciidoc#L9>
//
// --------------------------------------------------------------------------------
// POST /sales/_search?size=0
// {
//     "aggs" : {
//         "t_shirts" : {
//             "filter" : { "term": { "type": "t-shirt" } },
//             "aggs" : {
//                 "avg_price" : { "avg" : { "field" : "price" } }
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_filter_aggregation_b93ed4ef309819734f0eeea82e8b0f1f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b93ed4ef309819734f0eeea82e8b0f1f[]
	res, err := es.Search(
		es.Search.WithIndex("sales"),
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "t_shirts": {
		      "filter": {
		        "term": {
		          "type": "t-shirt"
		        }
		      },
		      "aggs": {
		        "avg_price": {
		          "avg": {
		            "field": "price"
		          }
		        }
		      }
		    }
		  }
		}`)),
		es.Search.WithSize(0),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:b93ed4ef309819734f0eeea82e8b0f1f[]
}
