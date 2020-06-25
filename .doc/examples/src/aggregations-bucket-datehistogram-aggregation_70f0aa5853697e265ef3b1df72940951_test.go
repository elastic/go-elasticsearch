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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/datehistogram-aggregation.asciidoc#L380>
//
// --------------------------------------------------------------------------------
// PUT my_index/_doc/1?refresh
// {
//   "date": "2015-10-01T00:30:00Z"
// }
//
// PUT my_index/_doc/2?refresh
// {
//   "date": "2015-10-01T01:30:00Z"
// }
//
// GET my_index/_search?size=0
// {
//   "aggs": {
//     "by_day": {
//       "date_histogram": {
//         "field":     "date",
//         "calendar_interval":  "day"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_datehistogram_aggregation_70f0aa5853697e265ef3b1df72940951(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:70f0aa5853697e265ef3b1df72940951[]
	{
		res, err := es.Index(
			"my_index",
			strings.NewReader(`{
		  "date": "2015-10-01T00:30:00Z"
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Index(
			"my_index",
			strings.NewReader(`{
		  "date": "2015-10-01T01:30:00Z"
		}`),
			es.Index.WithDocumentID("2"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithIndex("my_index"),
			es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "by_day": {
		      "date_histogram": {
		        "field": "date",
		        "calendar_interval": "day"
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
	}
	// end:70f0aa5853697e265ef3b1df72940951[]
}
