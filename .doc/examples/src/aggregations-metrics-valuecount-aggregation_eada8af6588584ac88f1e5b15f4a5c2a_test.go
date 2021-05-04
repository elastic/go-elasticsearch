// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/metrics/valuecount-aggregation.asciidoc#L96>
//
// --------------------------------------------------------------------------------
// PUT metrics_index/_doc/1
// {
//   "network.name" : "net-1",
//   "latency_histo" : {
//       "values" : [0.1, 0.2, 0.3, 0.4, 0.5],
//       "counts" : [3, 7, 23, 12, 6] <1>
//    }
// }
//
// PUT metrics_index/_doc/2
// {
//   "network.name" : "net-2",
//   "latency_histo" : {
//       "values" :  [0.1, 0.2, 0.3, 0.4, 0.5],
//       "counts" : [8, 17, 8, 7, 6] <1>
//    }
// }
//
// POST /metrics_index/_search?size=0
// {
//   "aggs": {
//     "total_requests": {
//       "value_count": { "field": "latency_histo" }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_metrics_valuecount_aggregation_eada8af6588584ac88f1e5b15f4a5c2a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:eada8af6588584ac88f1e5b15f4a5c2a[]
	{
		res, err := es.Index(
			"metrics_index",
			strings.NewReader(`{
		  "network.name": "net-1",
		  "latency_histo": {
		    "values": [
		      0.1,
		      0.2,
		      0.3,
		      0.4,
		      0.5
		    ],
		    "counts": [
		      3,
		      7,
		      23,
		      12,
		      6
		    ]
		  }
		}`),
			es.Index.WithDocumentID("1"),
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
			"metrics_index",
			strings.NewReader(`{
		  "network.name": "net-2",
		  "latency_histo": {
		    "values": [
		      0.1,
		      0.2,
		      0.3,
		      0.4,
		      0.5
		    ],
		    "counts": [
		      8,
		      17,
		      8,
		      7,
		      6
		    ]
		  }
		}`),
			es.Index.WithDocumentID("2"),
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
			es.Search.WithIndex("metrics_index"),
			es.Search.WithBody(strings.NewReader(`{
		  "aggs": {
		    "total_requests": {
		      "value_count": {
		        "field": "latency_histo"
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
	// end:eada8af6588584ac88f1e5b15f4a5c2a[]
}
