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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/setup/install/check-running.asciidoc#L7>
//
// --------------------------------------------------------------------------------
// GET /
// --------------------------------------------------------------------------------

func Test_setup_install_check_running_3d1ff6097e2359f927c88c2ccdb36252(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3d1ff6097e2359f927c88c2ccdb36252[]
	res, err := es.Info()
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:3d1ff6097e2359f927c88c2ccdb36252[]
}
