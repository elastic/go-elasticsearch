// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/bulk.asciidoc#L646>
//
// --------------------------------------------------------------------------------
// POST /_bulk?filter_path=items.*.error
// { "update": {"_id": "5", "_index": "index1"} }
// { "doc": {"my_field": "baz"} }
// { "update": {"_id": "6", "_index": "index1"} }
// { "doc": {"my_field": "baz"} }
// { "update": {"_id": "7", "_index": "index1"} }
// { "doc": {"my_field": "baz"} }
// --------------------------------------------------------------------------------

func Test_docs_bulk_bfdad8a928ea30d7cf60d0a0a6bc6e2e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:bfdad8a928ea30d7cf60d0a0a6bc6e2e[]
	res, err := es.Bulk(
		strings.NewReader(`
{ "update": {"_id": "5", "_index": "index1"} }
{ "doc": {"my_field": "baz"} }
{ "update": {"_id": "6", "_index": "index1"} }
{ "doc": {"my_field": "baz"} }
{ "update": {"_id": "7", "_index": "index1"} }
{ "doc": {"my_field": "baz"} }
`),
		es.Bulk.WithFilterPath("items.*.error"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:bfdad8a928ea30d7cf60d0a0a6bc6e2e[]
}
