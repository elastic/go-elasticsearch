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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/from-size.asciidoc#L19>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "from": 5,
//   "size": 20,
//   "query": {
//     "match": {
//       "user.id": "kimchy"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_request_from_size_820f689eaaef15fc07abd1073fa880f8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:820f689eaaef15fc07abd1073fa880f8[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "from": 5,
		  "size": 20,
		  "query": {
		    "match": {
		      "user.id": "kimchy"
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
	// end:820f689eaaef15fc07abd1073fa880f8[]
}
