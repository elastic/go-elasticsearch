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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update.asciidoc#L325>
//
// --------------------------------------------------------------------------------
// POST test/_update/1
// {
//     "doc" : {
//         "name" : "new_name"
//     },
//     "doc_as_upsert" : true
// }
// --------------------------------------------------------------------------------

func Test_docs_update_7cac05cb589f1614fd5b8589153bef06(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:7cac05cb589f1614fd5b8589153bef06[]
	res, err := es.Update(
		"test",
		"1",
		strings.NewReader(`{
		  "doc": {
		    "name": "new_name"
		  },
		  "doc_as_upsert": true
		}`),
		es.Update.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:7cac05cb589f1614fd5b8589153bef06[]
}
