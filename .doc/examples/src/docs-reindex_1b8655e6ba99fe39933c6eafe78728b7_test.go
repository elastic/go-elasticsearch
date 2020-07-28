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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L204>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": "twitter",
//     "slice": {
//       "id": 0,
//       "max": 2
//     }
//   },
//   "dest": {
//     "index": "new_twitter"
//   }
// }
// POST _reindex
// {
//   "source": {
//     "index": "twitter",
//     "slice": {
//       "id": 1,
//       "max": 2
//     }
//   },
//   "dest": {
//     "index": "new_twitter"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_1b8655e6ba99fe39933c6eafe78728b7(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1b8655e6ba99fe39933c6eafe78728b7[]
	{
		res, err := es.Reindex(
			strings.NewReader(`{
		  "source": {
		    "index": "twitter",
		    "slice": {
		      "id": 0,
		      "max": 2
		    }
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
	}

	{
		res, err := es.Reindex(
			strings.NewReader(`{
		  "source": {
		    "index": "twitter",
		    "slice": {
		      "id": 1,
		      "max": 2
		    }
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
	}
	// end:1b8655e6ba99fe39933c6eafe78728b7[]
}
