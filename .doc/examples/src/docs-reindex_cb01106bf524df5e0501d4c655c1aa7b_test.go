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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L264>
//
// --------------------------------------------------------------------------------
// POST _reindex?slices=5&refresh
// {
//   "source": {
//     "index": "twitter"
//   },
//   "dest": {
//     "index": "new_twitter"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_cb01106bf524df5e0501d4c655c1aa7b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:cb01106bf524df5e0501d4c655c1aa7b[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "twitter"
		  },
		  "dest": {
		    "index": "new_twitter"
		  }
		}`),
		es.Reindex.WithRefresh(true),
		es.Reindex.WithSlices("5"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:cb01106bf524df5e0501d4c655c1aa7b[]
}
