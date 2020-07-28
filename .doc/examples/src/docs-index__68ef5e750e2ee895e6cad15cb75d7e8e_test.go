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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/index_.asciidoc#L534>
//
// --------------------------------------------------------------------------------
// PUT twitter/_create/1
// {
//   "user" : "kimchy",
//   "post_date" : "2009-11-15T14:12:12",
//   "message" : "trying out Elasticsearch"
// }
// --------------------------------------------------------------------------------

func Test_docs_index__68ef5e750e2ee895e6cad15cb75d7e8e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:68ef5e750e2ee895e6cad15cb75d7e8e[]
	res, err := es.Create(
		"twitter",
		"1",
		strings.NewReader(`{
		  "user": "kimchy",
		  "post_date": "2009-11-15T14:12:12",
		  "message": "trying out Elasticsearch"
		}`),
		es.Create.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:68ef5e750e2ee895e6cad15cb75d7e8e[]
}
