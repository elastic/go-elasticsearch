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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L386>
//
// --------------------------------------------------------------------------------
// GET twitter/_settings?flat_settings=true
// --------------------------------------------------------------------------------

func Test_api_conventions_b9a153725b28fdd0a5aabd7f17a8c2d7(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b9a153725b28fdd0a5aabd7f17a8c2d7[]
	res, err := es.Indices.GetSettings(
		es.Indices.GetSettings.WithIndex([]string{"twitter"}...),
		es.Indices.GetSettings.WithFlatSettings(true),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:b9a153725b28fdd0a5aabd7f17a8c2d7[]
}
