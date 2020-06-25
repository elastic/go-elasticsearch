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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L79>
//
// --------------------------------------------------------------------------------
// # GET /<logstash-{now/d}>/_search
// GET /%3Clogstash-%7Bnow%2Fd%7D%3E/_search
// {
//   "query" : {
//     "match": {
//       "test": "data"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_api_conventions_978088f989d45dd09339582e9cbc60e0(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:978088f989d45dd09339582e9cbc60e0[]
	res, err := es.Search(
		es.Search.WithIndex("%3Clogstash-%7Bnow%2Fd%7D%3E"),
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
	// end:978088f989d45dd09339582e9cbc60e0[]
}
