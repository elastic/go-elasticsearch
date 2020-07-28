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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L683>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": "twitter",
//     "_source": ["user", "_doc"]
//   },
//   "dest": {
//     "index": "new_twitter"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_e9c2e15b36372d5281c879d336322b6c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:e9c2e15b36372d5281c879d336322b6c[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "twitter",
		    "_source": [
		      "user",
		      "_doc"
		    ]
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
	// end:e9c2e15b36372d5281c879d336322b6c[]
}
