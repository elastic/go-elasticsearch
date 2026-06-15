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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L120>
//
// --------------------------------------------------------------------------------
// # Create the two indices
// PUT /twitter-1
// PUT /twitter-2
//
// # Update both mappings
// PUT /twitter-1,twitter-2/_mapping <1>
// {
//   "properties": {
//     "user_name": {
//       "type": "text"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_1da77e114459e0b77d78a3dcc8fae429(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1da77e114459e0b77d78a3dcc8fae429[]
	{
		res, err := es.Indices.Create("twitter-1")
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Indices.Create(
			"twitter-2",
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Indices.PutMapping(
			[]string{"twitter-1,twitter-2"},
			strings.NewReader(`{
		  "properties": {
		    "user_name": {
		      "type": "text"
		    }
		  }
		}`),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:1da77e114459e0b77d78a3dcc8fae429[]
}
