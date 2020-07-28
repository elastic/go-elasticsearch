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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/update-settings.asciidoc#L103>
//
// --------------------------------------------------------------------------------
// PUT /twitter/_settings
// {
//   "index" : {
//     "refresh_interval" : "-1"
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_update_settings_f1fafd45a69147795642f974e37e7ad0(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:f1fafd45a69147795642f974e37e7ad0[]
	res, err := es.Indices.PutSettings(strings.NewReader(`{
		  "index": {
		    "refresh_interval": "-1"
		  }
		}`),

		es.Indices.PutSettings.WithIndex([]string{"twitter"}...),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:f1fafd45a69147795642f974e37e7ad0[]
}
