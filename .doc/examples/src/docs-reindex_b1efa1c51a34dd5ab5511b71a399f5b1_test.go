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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L403>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": "source"
//   },
//   "dest": {
//     "index": "dest",
//     "pipeline": "some_ingest_pipeline"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_b1efa1c51a34dd5ab5511b71a399f5b1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b1efa1c51a34dd5ab5511b71a399f5b1[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "source"
		  },
		  "dest": {
		    "index": "dest",
		    "pipeline": "some_ingest_pipeline"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:b1efa1c51a34dd5ab5511b71a399f5b1[]
}
