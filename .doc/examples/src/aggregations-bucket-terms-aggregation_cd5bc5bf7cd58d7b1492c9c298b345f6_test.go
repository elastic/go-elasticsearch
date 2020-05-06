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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L672>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "aggs" : {
//         "actors" : {
//              "terms" : {
//                  "field" : "actors",
//                  "size" : 10,
//                  "collect_mode" : "breadth_first" <1>
//              },
//             "aggs" : {
//                 "costars" : {
//                      "terms" : {
//                          "field" : "actors",
//                          "size" : 5
//                      }
//                  }
//             }
//          }
//     }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_cd5bc5bf7cd58d7b1492c9c298b345f6(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:cd5bc5bf7cd58d7b1492c9c298b345f6[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "actors": {
		      "terms": {
		        "field": "actors",
		        "size": 10,
		        "collect_mode": "breadth_first"
		      },
		      "aggs": {
		        "costars": {
		          "terms": {
		            "field": "actors",
		            "size": 5
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
	// end:cd5bc5bf7cd58d7b1492c9c298b345f6[]
}
