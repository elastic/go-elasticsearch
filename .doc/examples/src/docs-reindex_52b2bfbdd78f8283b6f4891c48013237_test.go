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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L635>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "max_docs": 1,
//   "source": {
//     "index": "twitter"
//   },
//   "dest": {
//     "index": "new_twitter"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_52b2bfbdd78f8283b6f4891c48013237(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:52b2bfbdd78f8283b6f4891c48013237[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "max_docs": 1,
		  "source": {
		    "index": "twitter"
		  },
		  "dest": {
		    "index": "new_twitter"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:52b2bfbdd78f8283b6f4891c48013237[]
}
