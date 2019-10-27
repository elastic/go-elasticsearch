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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/index_.asciidoc#L146>
//
// --------------------------------------------------------------------------------
// PUT _cluster/settings
// {
//     "persistent": {
//         "action.auto_create_index": "twitter,index10,-index1*,+ind*" <1>
//     }
// }
//
// PUT _cluster/settings
// {
//     "persistent": {
//         "action.auto_create_index": "false" <2>
//     }
// }
//
// PUT _cluster/settings
// {
//     "persistent": {
//         "action.auto_create_index": "true" <3>
//     }
// }
// --------------------------------------------------------------------------------

func Test_docs_index__804a97ff4d0613e6568e4efb19c52021(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:804a97ff4d0613e6568e4efb19c52021[]
	{
		res, err := es.Cluster.PutSettings(
			strings.NewReader(`{
		  "persistent": {
		    "action.auto_create_index": "twitter,index10,-index1*,+ind*"
		  }
		}`))
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Cluster.PutSettings(
			strings.NewReader(`{
		  "persistent": {
		    "action.auto_create_index": "false"
		  }
		}`))
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Cluster.PutSettings(
			strings.NewReader(`{
		  "persistent": {
		    "action.auto_create_index": "true"
		  }
		}`))
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:804a97ff4d0613e6568e4efb19c52021[]
}
