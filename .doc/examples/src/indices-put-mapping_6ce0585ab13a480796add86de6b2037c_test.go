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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/put-mapping.asciidoc#L441>
//
// --------------------------------------------------------------------------------
// POST /users/_doc?refresh=wait_for
// {
//   "user_id" : 12345
// }
//
// POST /users/_doc?refresh=wait_for
// {
//   "user_id" : 12346
// }
// --------------------------------------------------------------------------------

func Test_indices_put_mapping_6ce0585ab13a480796add86de6b2037c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6ce0585ab13a480796add86de6b2037c[]
	{
		res, err := es.Index(
			"users",
			strings.NewReader(`{
		  "user_id": 12345
		}`),
			es.Index.WithRefresh("wait_for"),
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
			"users",
			strings.NewReader(`{
		  "user_id": 12346
		}`),
			es.Index.WithRefresh("wait_for"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:6ce0585ab13a480796add86de6b2037c[]
}
