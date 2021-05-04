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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search.asciidoc#L72>
//
// --------------------------------------------------------------------------------
// PUT /_cluster/settings
// {
//   "transient": {
//     "cluster.routing.use_adaptive_replica_selection": false
//   }
// }
// --------------------------------------------------------------------------------

func Test_search_3227b17f70d3cb3a5f2b296d6943848a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3227b17f70d3cb3a5f2b296d6943848a[]
	res, err := es.Cluster.PutSettings(
		strings.NewReader(`{
		  "transient": {
		    "cluster.routing.use_adaptive_replica_selection": false
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:3227b17f70d3cb3a5f2b296d6943848a[]
}
