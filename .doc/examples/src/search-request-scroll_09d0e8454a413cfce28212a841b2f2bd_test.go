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
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request/scroll.asciidoc#L206>
//
// --------------------------------------------------------------------------------
// GET /twitter/_search?scroll=1m
// {
//   "slice": {
//     "id": 0,                      <1>
//     "max": 2                      <2>
//   },
//   "query": {
//     "match": {
//       "title": "elasticsearch"
//     }
//   }
// }
// GET /twitter/_search?scroll=1m
// {
//   "slice": {
//     "id": 1,
//     "max": 2
//   },
//   "query": {
//     "match": {
//       "title": "elasticsearch"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_request_scroll_09d0e8454a413cfce28212a841b2f2bd(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:09d0e8454a413cfce28212a841b2f2bd[]
	{
		res, err := es.Search(
			es.Search.WithIndex("twitter"),
			es.Search.WithBody(strings.NewReader(`{
		  "slice": {
		    "id": 0,
		    "max": 2
		  },
		  "query": {
		    "match": {
		      "title": "elasticsearch"
		    }
		  }
		}`)),
			es.Search.WithScroll(time.Duration(60000000000)),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithIndex("twitter"),
			es.Search.WithBody(strings.NewReader(`{
		  "slice": {
		    "id": 1,
		    "max": 2
		  },
		  "query": {
		    "match": {
		      "title": "elasticsearch"
		    }
		  }
		}`)),
			es.Search.WithScroll(time.Duration(60000000000)),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:09d0e8454a413cfce28212a841b2f2bd[]
}
