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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/bulk.asciidoc#L545>
//
// --------------------------------------------------------------------------------
// POST _bulk
// { "update" : {"_id" : "1", "_index" : "index1", "retry_on_conflict" : 3} }
// { "doc" : {"field" : "value"} }
// { "update" : { "_id" : "0", "_index" : "index1", "retry_on_conflict" : 3} }
// { "script" : { "source": "ctx._source.counter += params.param1", "lang" : "painless", "params" : {"param1" : 1}}, "upsert" : {"counter" : 1}}
// { "update" : {"_id" : "2", "_index" : "index1", "retry_on_conflict" : 3} }
// { "doc" : {"field" : "value"}, "doc_as_upsert" : true }
// { "update" : {"_id" : "3", "_index" : "index1", "_source" : true} }
// { "doc" : {"field" : "value"} }
// { "update" : {"_id" : "4", "_index" : "index1"} }
// { "doc" : {"field" : "value"}, "_source": true}
// --------------------------------------------------------------------------------

func Test_docs_bulk_8cd00a3aba7c3c158277bc032aac2830(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8cd00a3aba7c3c158277bc032aac2830[]
	res, err := es.Bulk(
		strings.NewReader(`
{ "update" : {"_id" : "1", "_index" : "index1", "retry_on_conflict" : 3} }
{ "doc" : {"field" : "value"} }
{ "update" : { "_id" : "0", "_index" : "index1", "retry_on_conflict" : 3} }
{ "script" : { "source": "ctx._source.counter += params.param1", "lang" : "painless", "params" : {"param1" : 1}}, "upsert" : {"counter" : 1}}
{ "update" : {"_id" : "2", "_index" : "index1", "retry_on_conflict" : 3} }
{ "doc" : {"field" : "value"}, "doc_as_upsert" : true }
{ "update" : {"_id" : "3", "_index" : "index1", "_source" : true} }
{ "doc" : {"field" : "value"} }
{ "update" : {"_id" : "4", "_index" : "index1"} }
{ "doc" : {"field" : "value"}, "_source": true}
`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:8cd00a3aba7c3c158277bc032aac2830[]
}
