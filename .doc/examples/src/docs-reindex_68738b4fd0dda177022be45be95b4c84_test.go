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
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L174>
//
// --------------------------------------------------------------------------------
// POST _reindex/r1A2WoRbTwKZ516z6NEs5A:36619/_rethrottle?requests_per_second=-1
// --------------------------------------------------------------------------------

func Test_docs_reindex_68738b4fd0dda177022be45be95b4c84(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:68738b4fd0dda177022be45be95b4c84[]
	res, err := es.ReindexRethrottle(
		"r1A2WoRbTwKZ516z6NEs5A:36619",
		esapi.IntPtr(-1),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:68738b4fd0dda177022be45be95b4c84[]
}
