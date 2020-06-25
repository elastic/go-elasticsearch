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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L800>
//
// --------------------------------------------------------------------------------
// GET metricbeat-2016.05.30-1/_doc/1
// GET metricbeat-2016.05.31-1/_doc/1
// --------------------------------------------------------------------------------

func Test_docs_reindex_3b04cc894e6a47d57983484010feac0c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3b04cc894e6a47d57983484010feac0c[]
	{
		res, err := es.Get("metricbeat-2016.05.30-1", "1", es.Get.WithPretty())
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Get("metricbeat-2016.05.31-1", "1", es.Get.WithPretty())
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:3b04cc894e6a47d57983484010feac0c[]
}
