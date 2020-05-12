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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/from-size.asciidoc#L14>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "from" : 0, "size" : 10,
//     "query" : {
//         "term" : { "user" : "kimchy" }
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_request_from_size_9a26759ccbd338224ecaacf7c49ab08e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:9a26759ccbd338224ecaacf7c49ab08e[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "from": 0,
		  "size": 10,
		  "query": {
		    "term": {
		      "user": "kimchy"
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
	// end:9a26759ccbd338224ecaacf7c49ab08e[]
}
