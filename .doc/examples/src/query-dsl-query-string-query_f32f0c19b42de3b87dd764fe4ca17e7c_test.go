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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/query-dsl/query-string-query.asciidoc#L418>
//
// --------------------------------------------------------------------------------
// GET /_search
// {
//    "query": {
//        "query_string" : {
//            "default_field": "title",
//            "query" : "ny city",
//            "auto_generate_synonyms_phrase_query" : false
//        }
//    }
// }
// --------------------------------------------------------------------------------

func Test_query_dsl_query_string_query_f32f0c19b42de3b87dd764fe4ca17e7c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:f32f0c19b42de3b87dd764fe4ca17e7c[]
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "query_string": {
		      "default_field": "title",
		      "query": "ny city",
		      "auto_generate_synonyms_phrase_query": false
		    }
		  }
		}`)),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:f32f0c19b42de3b87dd764fe4ca17e7c[]
}
