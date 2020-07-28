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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/terms-query.asciidoc#L19>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//   "query": {
//     "terms": {
//       "user": [ "kimchy", "elasticsearch" ],
//       "boost": 1.0
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_terms_query_ae563e7e5a445d0aa817cae7927b1c73(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:ae563e7e5a445d0aa817cae7927b1c73[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "terms": {
		      "user": [
		        "kimchy",
		        "elasticsearch"
		      ],
		      "boost": 1.0
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
	// end:ae563e7e5a445d0aa817cae7927b1c73[]
}
