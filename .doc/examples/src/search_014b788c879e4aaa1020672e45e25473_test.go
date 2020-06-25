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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/search.asciidoc#L72>
//
// --------------------------------------------------------------------------------
// PUT /_cluster/settings
// {
//     "transient": {
//         "cluster.routing.use_adaptive_replica_selection": false
//     }
// }
// --------------------------------------------------------------------------------

func Test_search_014b788c879e4aaa1020672e45e25473(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:014b788c879e4aaa1020672e45e25473[]
	res, err := es.Cluster.PutSettings(
		strings.NewReader(`{
		  "transient": {
		    "cluster.routing.use_adaptive_replica_selection": false
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:014b788c879e4aaa1020672e45e25473[]
}
