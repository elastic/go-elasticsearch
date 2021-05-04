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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/fields/id-field.asciidoc#L12>
//
// --------------------------------------------------------------------------------
// # Example documents
// PUT my-index-000001/_doc/1
// {
//   "text": "Document with ID 1"
// }
//
// PUT my-index-000001/_doc/2?refresh=true
// {
//   "text": "Document with ID 2"
// }
//
// GET my-index-000001/_search
// {
//   "query": {
//     "terms": {
//       "_id": [ "1", "2" ] <1>
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_fields_id_field_8c8b5224befab7804461c7e7b6086d9a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8c8b5224befab7804461c7e7b6086d9a[]
	{
		res, err := es.Index(
			"my-index-000001",
			strings.NewReader(`{
		  "text": "Document with ID 1"
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Index(
			"my-index-000001",
			strings.NewReader(`{
		  "text": "Document with ID 2"
		}`),
			es.Index.WithDocumentID("2"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithIndex("my-index-000001"),
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "terms": {
		      "_id": [
		        "1",
		        "2"
		      ]
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
	}
	// end:8c8b5224befab7804461c7e7b6086d9a[]
}
