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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update.asciidoc#L298>
//
// --------------------------------------------------------------------------------
// POST sessions/_update/dh3sgudg8gsrgl
// {
//   "scripted_upsert": true,
//   "script": {
//     "id": "my_web_session_summariser",
//     "params": {
//       "pageViewEvent": {
//         "url": "foo.com/bar",
//         "response": 404,
//         "time": "2014-01-01 12:32"
//       }
//     }
//   },
//   "upsert": {}
// }
// --------------------------------------------------------------------------------

func Test_docs_update_4832108e01f5efac5e6e40ac405998ca(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4832108e01f5efac5e6e40ac405998ca[]
	res, err := es.Update(
		"sessions",
		"dh3sgudg8gsrgl",
		strings.NewReader(`{
		  "scripted_upsert": true,
		  "script": {
		    "id": "my_web_session_summariser",
		    "params": {
		      "pageViewEvent": {
		        "url": "foo.com/bar",
		        "response": 404,
		        "time": "2014-01-01 12:32"
		      }
		    }
		  },
		  "upsert": {}
		}`),
		es.Update.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:4832108e01f5efac5e6e40ac405998ca[]
}
