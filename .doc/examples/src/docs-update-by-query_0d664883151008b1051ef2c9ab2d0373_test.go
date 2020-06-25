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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L530>
//
// --------------------------------------------------------------------------------
// POST twitter/_update_by_query
// {
//   "slice": {
//     "id": 0,
//     "max": 2
//   },
//   "script": {
//     "source": "ctx._source['extra'] = 'test'"
//   }
// }
// POST twitter/_update_by_query
// {
//   "slice": {
//     "id": 1,
//     "max": 2
//   },
//   "script": {
//     "source": "ctx._source['extra'] = 'test'"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_0d664883151008b1051ef2c9ab2d0373(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:0d664883151008b1051ef2c9ab2d0373[]
	{
		res, err := es.UpdateByQuery(
			[]string{"twitter"},
			es.UpdateByQuery.WithBody(strings.NewReader(`{
		  "slice": {
		    "id": 0,
		    "max": 2
		  },
		  "script": {
		    "source": "ctx._source['extra'] = 'test'"
		  }
		}`)),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.UpdateByQuery(
			[]string{"twitter"},
			es.UpdateByQuery.WithBody(strings.NewReader(`{
		  "slice": {
		    "id": 1,
		    "max": 2
		  },
		  "script": {
		    "source": "ctx._source['extra'] = 'test'"
		  }
		}`)),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:0d664883151008b1051ef2c9ab2d0373[]
}
