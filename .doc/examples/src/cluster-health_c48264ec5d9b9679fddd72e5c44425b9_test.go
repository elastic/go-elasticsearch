// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/cluster/health.asciidoc#L186>
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
