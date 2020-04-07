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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L384>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": "source",
//     "size": 100
//   },
//   "dest": {
//     "index": "dest",
//     "routing": "=cat"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_400e89eb46ead8e9c9e40f123fd5e590(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:400e89eb46ead8e9c9e40f123fd5e590[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "source",
		    "size": 100
		  },
		  "dest": {
		    "index": "dest",
		    "routing": "=cat"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:400e89eb46ead8e9c9e40f123fd5e590[]
}
