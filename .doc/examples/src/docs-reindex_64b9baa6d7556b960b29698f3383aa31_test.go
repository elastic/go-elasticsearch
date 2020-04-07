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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L955>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "remote": {
//       "host": "http://otherhost:9200"
//     },
//     "index": "source",
//     "size": 10,
//     "query": {
//       "match": {
//         "test": "data"
//       }
//     }
//   },
//   "dest": {
//     "index": "dest"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_64b9baa6d7556b960b29698f3383aa31(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:64b9baa6d7556b960b29698f3383aa31[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "remote": {
		      "host": "http://otherhost:9200"
		    },
		    "index": "source",
		    "size": 10,
		    "query": {
		      "match": {
		        "test": "data"
		      }
		    }
		  },
		  "dest": {
		    "index": "dest"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:64b9baa6d7556b960b29698f3383aa31[]
}
