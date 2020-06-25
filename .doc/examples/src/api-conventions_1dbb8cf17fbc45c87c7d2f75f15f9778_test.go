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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L250>
//
// --------------------------------------------------------------------------------
// GET /_cluster/state?filter_path=metadata.indices.*.stat*
// --------------------------------------------------------------------------------

func Test_api_conventions_1dbb8cf17fbc45c87c7d2f75f15f9778(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1dbb8cf17fbc45c87c7d2f75f15f9778[]
	res, err := es.Cluster.State(
		es.Cluster.State.WithFilterPath("metadata.indices.*.stat*"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1dbb8cf17fbc45c87c7d2f75f15f9778[]
}
