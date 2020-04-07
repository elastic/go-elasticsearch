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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/match-all-query.asciidoc#L39>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//     "query": {
//         "match_none": {}
//     }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_match_all_query_81c9aa2678d6166a9662ddf2c011a6a5(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:81c9aa2678d6166a9662ddf2c011a6a5[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match_none": {}
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:81c9aa2678d6166a9662ddf2c011a6a5[]
}
