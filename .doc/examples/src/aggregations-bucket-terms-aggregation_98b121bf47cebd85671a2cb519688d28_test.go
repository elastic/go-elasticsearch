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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L654>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "aggs" : {
//         "JapaneseCars" : {
//              "terms" : {
//                  "field" : "make",
//                  "include" : ["mazda", "honda"]
//              }
//          },
//         "ActiveCarManufacturers" : {
//              "terms" : {
//                  "field" : "make",
//                  "exclude" : ["rover", "jensen"]
//              }
//          }
//     }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_98b121bf47cebd85671a2cb519688d28(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:98b121bf47cebd85671a2cb519688d28[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "JapaneseCars": {
		      "terms": {
		        "field": "make",
		        "include": [
		          "mazda",
		          "honda"
		        ]
		      }
		    },
		    "ActiveCarManufacturers": {
		      "terms": {
		        "field": "make",
		        "exclude": [
		          "rover",
		          "jensen"
		        ]
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
	// end:98b121bf47cebd85671a2cb519688d28[]
}
