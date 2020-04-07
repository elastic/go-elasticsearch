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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L640>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": ["twitter", "blog"]
//   },
//   "dest": {
//     "index": "all_together"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_6f097c298a7abf4c032c4314920c49c8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6f097c298a7abf4c032c4314920c49c8[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": [
		      "twitter",
		      "blog"
		    ]
		  },
		  "dest": {
		    "index": "all_together"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:6f097c298a7abf4c032c4314920c49c8[]
}
