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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L370>
//
// --------------------------------------------------------------------------------
// PUT twitter/_doc/2?routing=user1
// {
//   "counter" : 1,
//   "tags" : ["white"]
// }
// --------------------------------------------------------------------------------

func Test_docs_get_4339ba0872d8a9def4f80d7ebb2e560a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4339ba0872d8a9def4f80d7ebb2e560a[]
	res, err := es.Index(
		"twitter",
		strings.NewReader(`{
		  "counter": 1,
		  "tags": [
		    "white"
		  ]
		}`),
		es.Index.WithDocumentID("2"),
		es.Index.WithRouting("user1"),
		es.Index.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:4339ba0872d8a9def4f80d7ebb2e560a[]
}
