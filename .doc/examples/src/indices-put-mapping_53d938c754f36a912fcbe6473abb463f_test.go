// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L463>
//
// --------------------------------------------------------------------------------
// POST /_reindex
// {
//   "source": {
//     "index": "users"
//   },
//   "dest": {
//     "index": "new_users"
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_53d938c754f36a912fcbe6473abb463f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:53d938c754f36a912fcbe6473abb463f[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "users"
		  },
		  "dest": {
		    "index": "new_users"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:53d938c754f36a912fcbe6473abb463f[]
}
