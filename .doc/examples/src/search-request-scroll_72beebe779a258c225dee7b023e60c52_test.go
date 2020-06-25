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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/scroll.asciidoc#L148>
//
// --------------------------------------------------------------------------------
// GET /_nodes/stats/indices/search
// --------------------------------------------------------------------------------

func Test_search_request_scroll_72beebe779a258c225dee7b023e60c52(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:72beebe779a258c225dee7b023e60c52[]
	res, err := es.Nodes.Stats(
		es.Nodes.Stats.WithMetric([]string{"indices"}...),
		es.Nodes.Stats.WithIndexMetric([]string{"search"}...),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:72beebe779a258c225dee7b023e60c52[]
}
