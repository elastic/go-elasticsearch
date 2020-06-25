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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L700>
//
// --------------------------------------------------------------------------------
// POST test/_doc/1?refresh
// {
//   "text": "words words",
//   "flag": "foo"
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_1577e6e806b3283c9e99f1596d310754(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1577e6e806b3283c9e99f1596d310754[]
	res, err := es.Index(
		"test",
		strings.NewReader(`{
		  "text": "words words",
		  "flag": "foo"
		}`),
		es.Index.WithDocumentID("1"),
		es.Index.WithRefresh("true"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1577e6e806b3283c9e99f1596d310754[]
}
