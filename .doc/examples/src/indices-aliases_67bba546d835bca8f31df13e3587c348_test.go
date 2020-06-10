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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/aliases.asciidoc#L443>
//
// --------------------------------------------------------------------------------
// GET /test/_doc/1
// --------------------------------------------------------------------------------

func Test_indices_aliases_67bba546d835bca8f31df13e3587c348(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:67bba546d835bca8f31df13e3587c348[]
	res, err := es.Get("test", "1", es.Get.WithPretty())
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:67bba546d835bca8f31df13e3587c348[]
}
