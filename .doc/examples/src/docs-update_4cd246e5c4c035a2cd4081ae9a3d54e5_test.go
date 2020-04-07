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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update.asciidoc#L114>
//
// --------------------------------------------------------------------------------
// POST test/_update/1
// {
//     "script" : {
//         "source": "ctx._source.tags.add(params.tag)",
//         "lang": "painless",
//         "params" : {
//             "tag" : "blue"
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_docs_update_4cd246e5c4c035a2cd4081ae9a3d54e5(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4cd246e5c4c035a2cd4081ae9a3d54e5[]
	res, err := es.Update(
		"test",
		"1",
		strings.NewReader(`{
		  "script": {
		    "source": "ctx._source.tags.add(params.tag)",
		    "lang": "painless",
		    "params": {
		      "tag": "blue"
		    }
		  }
		}`),
		es.Update.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:4cd246e5c4c035a2cd4081ae9a3d54e5[]
}
