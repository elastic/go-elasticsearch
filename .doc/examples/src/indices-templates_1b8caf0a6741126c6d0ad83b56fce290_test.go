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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/templates.asciidoc#L146>
//
// --------------------------------------------------------------------------------
// PUT _template/template_1
// {
//     "index_patterns" : ["te*"],
//     "settings" : {
//         "number_of_shards" : 1
//     },
//     "aliases" : {
//         "alias1" : {},
//         "alias2" : {
//             "filter" : {
//                 "term" : {"user" : "kimchy" }
//             },
//             "routing" : "kimchy"
//         },
//         "{index}-alias" : {} <1>
//     }
// }
// --------------------------------------------------------------------------------

func Test_indices_templates_1b8caf0a6741126c6d0ad83b56fce290(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:1b8caf0a6741126c6d0ad83b56fce290[]
	res, err := es.Indices.PutTemplate(
		"template_1",
		strings.NewReader(`{
		  "index_patterns": [
		    "te*"
		  ],
		  "settings": {
		    "number_of_shards": 1
		  },
		  "aliases": {
		    "alias1": {},
		    "alias2": {
		      "filter": {
		        "term": {
		          "user": "kimchy"
		        }
		      },
		      "routing": "kimchy"
		    },
		    "{index}-alias": {}
		  }
		}`),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:1b8caf0a6741126c6d0ad83b56fce290[]
}
