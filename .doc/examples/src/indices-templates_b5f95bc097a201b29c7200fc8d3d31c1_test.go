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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/templates.asciidoc#L170>
//
// --------------------------------------------------------------------------------
// PUT /_template/template_1
// {
//     "index_patterns" : ["*"],
//     "order" : 0,
//     "settings" : {
//         "number_of_shards" : 1
//     },
//     "mappings" : {
//         "_source" : { "enabled" : false }
//     }
// }
//
// PUT /_template/template_2
// {
//     "index_patterns" : ["te*"],
//     "order" : 1,
//     "settings" : {
//         "number_of_shards" : 1
//     },
//     "mappings" : {
//         "_source" : { "enabled" : true }
//     }
// }
// --------------------------------------------------------------------------------

func Test_indices_templates_b5f95bc097a201b29c7200fc8d3d31c1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:b5f95bc097a201b29c7200fc8d3d31c1[]
	{
		res, err := es.Indices.PutTemplate(
			"template_1",
			strings.NewReader(`{
		  "index_patterns": [
		    "*"
		  ],
		  "order": 0,
		  "settings": {
		    "number_of_shards": 1
		  },
		  "mappings": {
		    "_source": {
		      "enabled": false
		    }
		  }
		}`),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Indices.PutTemplate(
			"template_2",
			strings.NewReader(`{
		  "index_patterns": [
		    "te*"
		  ],
		  "order": 1,
		  "settings": {
		    "number_of_shards": 1
		  },
		  "mappings": {
		    "_source": {
		      "enabled": true
		    }
		  }
		}`),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:b5f95bc097a201b29c7200fc8d3d31c1[]
}
