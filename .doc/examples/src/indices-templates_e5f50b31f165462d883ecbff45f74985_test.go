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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/templates.asciidoc#L10>
//
// --------------------------------------------------------------------------------
// PUT _template/template_1
// {
//   "index_patterns": ["te*", "bar*"],
//   "settings": {
//     "number_of_shards": 1
//   },
//   "mappings": {
//     "_source": {
//       "enabled": false
//     },
//     "properties": {
//       "host_name": {
//         "type": "keyword"
//       },
//       "created_at": {
//         "type": "date",
//         "format": "EEE MMM dd HH:mm:ss Z yyyy"
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_indices_templates_e5f50b31f165462d883ecbff45f74985(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:e5f50b31f165462d883ecbff45f74985[]
	res, err := es.Indices.PutTemplate(
		"template_1",
		strings.NewReader(`{
		  "index_patterns": [
		    "te*",
		    "bar*"
		  ],
		  "settings": {
		    "number_of_shards": 1
		  },
		  "mappings": {
		    "_source": {
		      "enabled": false
		    },
		    "properties": {
		      "host_name": {
		        "type": "keyword"
		      },
		      "created_at": {
		        "type": "date",
		        "format": "EEE MMM dd HH:mm:ss Z yyyy"
		      }
		    }
		  }
		}`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:e5f50b31f165462d883ecbff45f74985[]
}
