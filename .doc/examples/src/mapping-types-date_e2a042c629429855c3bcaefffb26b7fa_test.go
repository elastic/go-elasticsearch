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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/types/date.asciidoc#L77>
//
// --------------------------------------------------------------------------------
// PUT my_index
// {
//   "mappings": {
//     "properties": {
//       "date": {
//         "type":   "date",
//         "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_types_date_e2a042c629429855c3bcaefffb26b7fa(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:e2a042c629429855c3bcaefffb26b7fa[]
	res, err := es.Indices.Create(
		"my_index",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "date": {
		        "type": "date",
		        "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
		      }
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:e2a042c629429855c3bcaefffb26b7fa[]
}
