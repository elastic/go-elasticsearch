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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/update-settings.asciidoc#L79>
//
// --------------------------------------------------------------------------------
// PUT /twitter/_settings
// {
//   "index" : {
//     "refresh_interval" : null
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_update_settings_a05e022599d1789e90ff3a41469b0b6d(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:a05e022599d1789e90ff3a41469b0b6d[]
	res, err := es.Indices.PutSettings(strings.NewReader(`{
		  "index": {
		    "refresh_interval": null
		  }
		}`),

		es.Indices.PutSettings.WithIndex([]string{"twitter"}...),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:a05e022599d1789e90ff3a41469b0b6d[]
}
