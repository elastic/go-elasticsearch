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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L557>
//
// --------------------------------------------------------------------------------
// GET _refresh
// POST twitter/_search?size=0&q=extra:test&filter_path=hits.total
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_4acf902c2598b2558f34f20c1744c433(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4acf902c2598b2558f34f20c1744c433[]
	{
		res, err := es.Indices.Refresh()
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithIndex("twitter"),
			es.Search.WithFilterPath("hits.total"),
			es.Search.WithQuery("extra:test"),
			es.Search.WithSize(0),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:4acf902c2598b2558f34f20c1744c433[]
}
