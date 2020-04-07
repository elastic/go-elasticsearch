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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L986>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "remote": {
//       "host": "http://otherhost:9200",
//       "socket_timeout": "1m",
//       "connect_timeout": "10s"
//     },
//     "index": "source",
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

func Test_docs_reindex_7f697eb436dfa3c30dfe610d8c32d132(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:7f697eb436dfa3c30dfe610d8c32d132[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "remote": {
		      "host": "http://otherhost:9200",
		      "socket_timeout": "1m",
		      "connect_timeout": "10s"
		    },
		    "index": "source",
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
	// end:7f697eb436dfa3c30dfe610d8c32d132[]
}
