// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/scroll.asciidoc#L186>
//
// --------------------------------------------------------------------------------
// DELETE /_search/scroll/_all
// --------------------------------------------------------------------------------

func Test_search_request_scroll_c2c21e2824fbf6b7198ede30419da82b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:c2c21e2824fbf6b7198ede30419da82b[]
	res, err := es.ClearScroll(
		es.ClearScroll.WithScrollID("_all"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:c2c21e2824fbf6b7198ede30419da82b[]
}
