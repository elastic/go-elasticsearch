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
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search/request-body.asciidoc#L65>
//
// --------------------------------------------------------------------------------
// GET /_search?q=message:number&size=0&terminate_after=1
// --------------------------------------------------------------------------------

func Test_search_request_body_bfcd65ab85d684d36a8550080032958d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:bfcd65ab85d684d36a8550080032958d[]
	res, err := es.Search(
		es.Search.WithQuery("message:number"),
		es.Search.WithSize(0),
		es.Search.WithTerminateAfter(1),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:bfcd65ab85d684d36a8550080032958d[]
}
