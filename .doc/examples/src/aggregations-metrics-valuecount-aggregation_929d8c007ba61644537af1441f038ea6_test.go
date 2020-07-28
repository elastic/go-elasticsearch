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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/metrics/valuecount-aggregation.asciidoc#L65>
//
// --------------------------------------------------------------------------------
// POST /sales/_search?size=0
// {
//   "aggs": {
//     "types_count": {
//       "value_count": {
//         "script": {
//           "id": "my_script",
//           "params": {
//             "field": "type"
//           }
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_metrics_valuecount_aggregation_929d8c007ba61644537af1441f038ea6(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:929d8c007ba61644537af1441f038ea6[]
	res, err := es.Search(
		es.Search.WithIndex("sales"),
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "types_count": {
		      "value_count": {
		        "script": {
		          "id": "my_script",
		          "params": {
		            "field": "type"
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
	// end:929d8c007ba61644537af1441f038ea6[]
}
