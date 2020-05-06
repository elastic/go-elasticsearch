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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L309>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "aggs" : {
//         "countries" : {
//             "terms" : {
//                 "field" : "artist.country",
//                 "order" : { "rock>playback_stats.avg" : "desc" }
//             },
//             "aggs" : {
//                 "rock" : {
//                     "filter" : { "term" : { "genre" :  "rock" }},
//                     "aggs" : {
//                         "playback_stats" : { "stats" : { "field" : "play_count" }}
//                     }
//                 }
//             }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_dc15e2373e5ecbe09b4ea0858eb63d47(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:dc15e2373e5ecbe09b4ea0858eb63d47[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "countries": {
		      "terms": {
		        "field": "artist.country",
		        "order": {
		          "rock>playback_stats.avg": "desc"
		        }
		      },
		      "aggs": {
		        "rock": {
		          "filter": {
		            "term": {
		              "genre": "rock"
		            }
		          },
		          "aggs": {
		            "playback_stats": {
		              "stats": {
		                "field": "play_count"
		              }
		            }
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
	// end:dc15e2373e5ecbe09b4ea0858eb63d47[]
}
