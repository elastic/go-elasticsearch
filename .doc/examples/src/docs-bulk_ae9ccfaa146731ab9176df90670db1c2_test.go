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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/bulk.asciidoc#L11>
//
// --------------------------------------------------------------------------------
// POST _bulk
// { "index" : { "_index" : "test", "_id" : "1" } }
// { "field1" : "value1" }
// { "delete" : { "_index" : "test", "_id" : "2" } }
// { "create" : { "_index" : "test", "_id" : "3" } }
// { "field1" : "value3" }
// { "update" : {"_id" : "1", "_index" : "test"} }
// { "doc" : {"field2" : "value2"} }
// --------------------------------------------------------------------------------

func Test_docs_bulk_ae9ccfaa146731ab9176df90670db1c2(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:ae9ccfaa146731ab9176df90670db1c2[]
	res, err := es.Bulk(
		strings.NewReader(`
{ "index" : { "_index" : "test", "_id" : "1" } }
{ "field1" : "value1" }
{ "delete" : { "_index" : "test", "_id" : "2" } }
{ "create" : { "_index" : "test", "_id" : "3" } }
{ "field1" : "value3" }
{ "update" : {"_id" : "1", "_index" : "test"} }
{ "doc" : {"field2" : "value2"} }
`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:ae9ccfaa146731ab9176df90670db1c2[]
}
