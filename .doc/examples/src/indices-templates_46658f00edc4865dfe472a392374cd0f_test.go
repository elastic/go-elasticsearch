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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/templates.asciidoc#L239>
//
// --------------------------------------------------------------------------------
// GET /_template/template_1?filter_path=*.version
// --------------------------------------------------------------------------------

func Test_indices_templates_46658f00edc4865dfe472a392374cd0f(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:46658f00edc4865dfe472a392374cd0f[]
	res, err := es.Indices.GetTemplate(
		es.Indices.GetTemplate.WithName("template_1"),
		es.Indices.GetTemplate.WithFilterPath("*.version"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:46658f00edc4865dfe472a392374cd0f[]
}
