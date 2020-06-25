// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/aliases.asciidoc#L388>
//
// --------------------------------------------------------------------------------
// GET /alias2/_search?q=user:kimchy&routing=2,3
// --------------------------------------------------------------------------------

func Test_indices_aliases_427f6b5c5376cbf0f71f242a60ca3d9e(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:427f6b5c5376cbf0f71f242a60ca3d9e[]
	res, err := es.Search(
		es.Search.WithIndex("alias2"),
		es.Search.WithQuery("user:kimchy"),
		es.Search.WithRouting("2,3"),
		es.Search.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:427f6b5c5376cbf0f71f242a60ca3d9e[]
}
