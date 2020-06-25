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
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete.asciidoc#L131>
//
// --------------------------------------------------------------------------------
// DELETE /twitter/_doc/1?timeout=5m
// --------------------------------------------------------------------------------

func Test_docs_delete_d90a84a24a407731dfc1929ac8327746(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d90a84a24a407731dfc1929ac8327746[]
	res, err := es.Delete("twitter",
		"1",
		es.Delete.WithTimeout(time.Duration(300000000000)),
		es.Delete.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:d90a84a24a407731dfc1929ac8327746[]
}
