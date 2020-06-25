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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/bulk.asciidoc#L564>
//
// --------------------------------------------------------------------------------
// POST /_bulk
// { "update": {"_id": "5", "_index": "index1"} }
// { "doc": {"my_field": "foo"} }
// { "update": {"_id": "6", "_index": "index1"} }
// { "doc": {"my_field": "foo"} }
// { "create": {"_id": "7", "_index": "index1"} }
// { "my_field": "foo" }
// --------------------------------------------------------------------------------

func Test_docs_bulk_1aa91d3d48140d6367b6cabca8737b8f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1aa91d3d48140d6367b6cabca8737b8f[]
	res, err := es.Bulk(
		strings.NewReader(`
{ "update": {"_id": "5", "_index": "index1"} }
{ "doc": {"my_field": "foo"} }
{ "update": {"_id": "6", "_index": "index1"} }
{ "doc": {"my_field": "foo"} }
{ "create": {"_id": "7", "_index": "index1"} }
{ "my_field": "foo" }
`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1aa91d3d48140d6367b6cabca8737b8f[]
}
