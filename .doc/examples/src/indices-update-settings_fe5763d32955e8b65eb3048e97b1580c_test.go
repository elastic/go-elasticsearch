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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/update-settings.asciidoc#L127>
//
// --------------------------------------------------------------------------------
// POST /twitter/_forcemerge?max_num_segments=5
// --------------------------------------------------------------------------------

func Test_indices_update_settings_fe5763d32955e8b65eb3048e97b1580c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:fe5763d32955e8b65eb3048e97b1580c[]
	res, err := es.Indices.Forcemerge(es.Indices.Forcemerge.WithIndex([]string{"twitter"}...),
		es.Indices.Forcemerge.WithMaxNumSegments(5),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:fe5763d32955e8b65eb3048e97b1580c[]
}
