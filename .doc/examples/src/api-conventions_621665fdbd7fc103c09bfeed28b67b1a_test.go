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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L298>
//
// --------------------------------------------------------------------------------
// GET /_count?filter_path=-_shards
// --------------------------------------------------------------------------------

func Test_api_conventions_621665fdbd7fc103c09bfeed28b67b1a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:621665fdbd7fc103c09bfeed28b67b1a[]
	res, err := es.Count(
		es.Count.WithFilterPath("-_shards"),
		es.Count.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:621665fdbd7fc103c09bfeed28b67b1a[]
}
