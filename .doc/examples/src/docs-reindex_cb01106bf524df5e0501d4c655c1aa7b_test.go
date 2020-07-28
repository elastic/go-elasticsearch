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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L264>
//
// --------------------------------------------------------------------------------
// POST _reindex?slices=5&refresh
// {
//   "source": {
//     "index": "twitter"
//   },
//   "dest": {
//     "index": "new_twitter"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_cb01106bf524df5e0501d4c655c1aa7b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:cb01106bf524df5e0501d4c655c1aa7b[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "twitter"
		  },
		  "dest": {
		    "index": "new_twitter"
		  }
		}`),
		es.Reindex.WithRefresh(true),
		es.Reindex.WithSlices("5"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:cb01106bf524df5e0501d4c655c1aa7b[]
}
