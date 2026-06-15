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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/aggregations/bucket/terms-aggregation.asciidoc#L549>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//    "size": 0,
//    "aggs": {
//       "expired_sessions": {
//          "terms": {
//             "field": "account_id",
//             "include": {
//                "partition": 0,
//                "num_partitions": 20
//             },
//             "size": 10000,
//             "order": {
//                "last_access": "asc"
//             }
//          },
//          "aggs": {
//             "last_access": {
//                "max": {
//                   "field": "access_date"
//                }
//             }
//          }
//       }
//    }
// }
// --------------------------------------------------------------------------------

func Test_aggregations_bucket_terms_aggregation_5d9d7b84e2fec7ecd832145cbb951cf1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:5d9d7b84e2fec7ecd832145cbb951cf1[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "size": 0,
		  "aggs": {
		    "expired_sessions": {
		      "terms": {
		        "field": "account_id",
		        "include": {
		          "partition": 0,
		          "num_partitions": 20
		        },
		        "size": 10000,
		        "order": {
		          "last_access": "asc"
		        }
		      },
		      "aggs": {
		        "last_access": {
		          "max": {
		            "field": "access_date"
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
	// end:5d9d7b84e2fec7ecd832145cbb951cf1[]
}
