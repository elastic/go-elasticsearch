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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/params/fielddata.asciidoc#L56>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "my_field": { <1>
//         "type": "text",
//         "fields": {
//           "keyword": { <2>
//             "type": "keyword"
//           }
//         }
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_params_fielddata_afef5cac988592b97ae289ab39c2f437(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:afef5cac988592b97ae289ab39c2f437[]
	res, err := es.Indices.Create(
		"my-index-000001",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "my_field": {
		        "type": "text",
		        "fields": {
		          "keyword": {
		            "type": "keyword"
		          }
		        }
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
	// end:afef5cac988592b97ae289ab39c2f437[]
}
