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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/cluster/health.asciidoc#L150>
//
// --------------------------------------------------------------------------------
// GET _cluster/health
// --------------------------------------------------------------------------------

func Test_cluster_health_b02e4907c9936c1adc16ccce9d49900d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b02e4907c9936c1adc16ccce9d49900d[]
	res, err := es.Cluster.Health()
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:b02e4907c9936c1adc16ccce9d49900d[]
}
