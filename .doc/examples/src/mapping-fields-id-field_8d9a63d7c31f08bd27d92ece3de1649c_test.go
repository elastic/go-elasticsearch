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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/fields/id-field.asciidoc#L12>
//
// --------------------------------------------------------------------------------
// # Example documents
// PUT my_index/_doc/1
// {
//   "text": "Document with ID 1"
// }
//
// PUT my_index/_doc/2?refresh=true
// {
//   "text": "Document with ID 2"
// }
//
// GET my_index/_search
// {
//   "query": {
//     "terms": {
//       "_id": [ "1", "2" ] <1>
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_fields_id_field_8d9a63d7c31f08bd27d92ece3de1649c(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:8d9a63d7c31f08bd27d92ece3de1649c[]
	{
		res, err := es.Index(
			"my_index",
			strings.NewReader(`{
		  "text": "Document with ID 1"
		}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Index(
			"my_index",
			strings.NewReader(`{
		  "text": "Document with ID 2"
		}`),
			es.Index.WithDocumentID("2"),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithIndex("my_index"),
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "terms": {
		      "_id": [
		        "1",
		        "2"
		      ]
		    }
		  }
		}`)),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:8d9a63d7c31f08bd27d92ece3de1649c[]
}
