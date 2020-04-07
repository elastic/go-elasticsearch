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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L751>
//
// --------------------------------------------------------------------------------
// PUT metricbeat-2016.05.30/_doc/1?refresh
// {"system.cpu.idle.pct": 0.908}
// PUT metricbeat-2016.05.31/_doc/1?refresh
// {"system.cpu.idle.pct": 0.105}
// --------------------------------------------------------------------------------

func Test_docs_reindex_9a4d5e41c52c20635d1fd9c6e13f6c7a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:9a4d5e41c52c20635d1fd9c6e13f6c7a[]
	{
		res, err := es.Index(
			"metricbeat-2016.05.30",
			strings.NewReader(`{
		  "system.cpu.idle.pct": 0.908
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Index(
			"metricbeat-2016.05.31",
			strings.NewReader(`{
		  "system.cpu.idle.pct": 0.105
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:9a4d5e41c52c20635d1fd9c6e13f6c7a[]
}
