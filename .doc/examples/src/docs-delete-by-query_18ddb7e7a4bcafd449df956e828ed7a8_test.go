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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L667>
//
// --------------------------------------------------------------------------------
// POST _tasks/r1A2WoRbTwKZ516z6NEs5A:36619/_cancel
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_18ddb7e7a4bcafd449df956e828ed7a8(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:18ddb7e7a4bcafd449df956e828ed7a8[]
	res, err := es.Tasks.Cancel(
		es.Tasks.Cancel.WithTaskID("r1A2WoRbTwKZ516z6NEs5A:36619"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:18ddb7e7a4bcafd449df956e828ed7a8[]
}
