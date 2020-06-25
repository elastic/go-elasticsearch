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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/scroll.asciidoc#L161>
//
// --------------------------------------------------------------------------------
// DELETE /_search/scroll
// {
//     "scroll_id" : "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAD4WYm9laVYtZndUQlNsdDcwakFMNjU1QQ=="
// }
// --------------------------------------------------------------------------------

func Test_search_request_scroll_b0d64d0a554549e5b2808002a0725493(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b0d64d0a554549e5b2808002a0725493[]
	res, err := es.ClearScroll(
		es.ClearScroll.WithBody(strings.NewReader(`{
		  "scroll_id": "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAD4WYm9laVYtZndUQlNsdDcwakFMNjU1QQ=="
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:b0d64d0a554549e5b2808002a0725493[]
}
