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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/scroll.asciidoc#L63>
//
// --------------------------------------------------------------------------------
// POST /_search/scroll                                                               <1>
// {
//   "scroll" : "1m",                                                                 <2>
//   "scroll_id" : "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAD4WYm9laVYtZndUQlNsdDcwakFMNjU1QQ==" <3>
// }
// --------------------------------------------------------------------------------

func Test_search_request_scroll_add240aa149d8b11139947502b279ee0(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:add240aa149d8b11139947502b279ee0[]
	res, err := es.Scroll(
		es.Scroll.WithBody(strings.NewReader(`{
		  "scroll": "1m",
		  "scroll_id": "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAD4WYm9laVYtZndUQlNsdDcwakFMNjU1QQ=="
		}`)),
		es.Scroll.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:add240aa149d8b11139947502b279ee0[]
}
