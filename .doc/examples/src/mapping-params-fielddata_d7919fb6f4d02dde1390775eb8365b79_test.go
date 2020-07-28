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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/params/fielddata.asciidoc#L84>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001/_mapping
// {
//   "properties": {
//     "my_field": { <1>
//       "type":     "text",
//       "fielddata": true
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_params_fielddata_d7919fb6f4d02dde1390775eb8365b79(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:d7919fb6f4d02dde1390775eb8365b79[]
	res, err := es.Indices.PutMapping(
		[]string{"my-index-000001"},
		strings.NewReader(`{
		  "properties": {
		    "my_field": {
		      "type": "text",
		      "fielddata": true
		    }
		  }
		}`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:d7919fb6f4d02dde1390775eb8365b79[]
}
