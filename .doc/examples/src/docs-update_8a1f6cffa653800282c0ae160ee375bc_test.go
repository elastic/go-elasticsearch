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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update.asciidoc#L137>
//
// --------------------------------------------------------------------------------
// POST test/_update/1
// {
//   "script": {
//     "source": "if (ctx._source.tags.contains(params.tag)) { ctx._source.tags.remove(ctx._source.tags.indexOf(params.tag)) }",
//     "lang": "painless",
//     "params": {
//       "tag": "blue"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_update_8a1f6cffa653800282c0ae160ee375bc(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8a1f6cffa653800282c0ae160ee375bc[]
	res, err := es.Update(
		"test",
		"1",
		strings.NewReader(`{
		  "script": {
		    "source": "if (ctx._source.tags.contains(params.tag)) { ctx._source.tags.remove(ctx._source.tags.indexOf(params.tag)) }",
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
	// end:8a1f6cffa653800282c0ae160ee375bc[]
}
