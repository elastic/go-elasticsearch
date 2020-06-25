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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/api-conventions.asciidoc#L317>
//
// --------------------------------------------------------------------------------
// GET /_cluster/state?filter_path=metadata.indices.*.state,-metadata.indices.logstash-*
// --------------------------------------------------------------------------------

func Test_api_conventions_1e18a67caf8f06ff2710ec4a8b30f625(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1e18a67caf8f06ff2710ec4a8b30f625[]
	res, err := es.Cluster.State(
		es.Cluster.State.WithFilterPath("metadata.indices.*.state,-metadata.indices.logstash-*"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1e18a67caf8f06ff2710ec4a8b30f625[]
}
