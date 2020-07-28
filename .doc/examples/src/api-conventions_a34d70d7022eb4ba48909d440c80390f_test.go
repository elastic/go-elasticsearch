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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L142>
//
// --------------------------------------------------------------------------------
// # GET /<logstash-{now/d-2d}>,<logstash-{now/d-1d}>,<logstash-{now/d}>/_search
// GET /%3Clogstash-%7Bnow%2Fd-2d%7D%3E%2C%3Clogstash-%7Bnow%2Fd-1d%7D%3E%2C%3Clogstash-%7Bnow%2Fd%7D%3E/_search
// {
//   "query" : {
//     "match": {
//       "test": "data"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_api_conventions_a34d70d7022eb4ba48909d440c80390f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:a34d70d7022eb4ba48909d440c80390f[]
	res, err := es.Search(
		es.Search.WithIndex("%3Clogstash-%7Bnow%2Fd-2d%7D%3E%2C%3Clogstash-%7Bnow%2Fd-1d%7D%3E%2C%3Clogstash-%7Bnow%2Fd%7D%3E"),
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "test": "data"
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
	// end:a34d70d7022eb4ba48909d440c80390f[]
}
