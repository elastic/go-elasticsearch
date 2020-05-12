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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/cluster/health.asciidoc#L179>
//
// --------------------------------------------------------------------------------
// GET /_cluster/health/twitter?level=shards
// --------------------------------------------------------------------------------

func Test_cluster_health_c48264ec5d9b9679fddd72e5c44425b9(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:c48264ec5d9b9679fddd72e5c44425b9[]
	res, err := es.Cluster.Health(
		es.Cluster.Health.WithIndex([]string{"twitter"}...),
		es.Cluster.Health.WithLevel("shards"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:c48264ec5d9b9679fddd72e5c44425b9[]
}
