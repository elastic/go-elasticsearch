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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/sort.asciidoc#L153>
//
// --------------------------------------------------------------------------------
// POST /index_long,index_double/_search
// {
//    "sort" : [
//       {
//         "field" : {
//             "numeric_type" : "double"
//         }
//       }
//    ]
// }
// --------------------------------------------------------------------------------

func Test_search_request_sort_2891aa10ee9d474780adf94d5607f2db(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:2891aa10ee9d474780adf94d5607f2db[]
	res, err := es.Search(
		es.Search.WithIndex("index_long,index_double"),
		es.Search.WithBody(strings.NewReader(`{
		  "sort": [
		    {
		      "field": {
		        "numeric_type": "double"
		      }
		    }
		  ]
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:2891aa10ee9d474780adf94d5607f2db[]
}
