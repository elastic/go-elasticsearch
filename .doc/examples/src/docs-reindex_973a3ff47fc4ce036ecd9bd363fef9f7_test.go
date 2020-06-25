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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L780>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": "metricbeat-*"
//   },
//   "dest": {
//     "index": "metricbeat"
//   },
//   "script": {
//     "lang": "painless",
//     "source": "ctx._index = 'metricbeat-' + (ctx._index.substring('metricbeat-'.length(), ctx._index.length())) + '-1'"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_973a3ff47fc4ce036ecd9bd363fef9f7(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:973a3ff47fc4ce036ecd9bd363fef9f7[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "metricbeat-*"
		  },
		  "dest": {
		    "index": "metricbeat"
		  },
		  "script": {
		    "lang": "painless",
		    "source": "ctx._index = 'metricbeat-' + (ctx._index.substring('metricbeat-'.length(), ctx._index.length())) + '-1'"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:973a3ff47fc4ce036ecd9bd363fef9f7[]
}
