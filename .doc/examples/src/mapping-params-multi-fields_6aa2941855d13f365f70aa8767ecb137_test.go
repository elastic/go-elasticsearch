// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/mapping/params/multi-fields.asciidoc#L10>
//
// --------------------------------------------------------------------------------
// PUT my-index-000001
// {
//   "mappings": {
//     "properties": {
//       "city": {
//         "type": "text",
//         "fields": {
//           "raw": { <1>
//             "type":  "keyword"
//           }
//         }
//       }
//     }
//   }
// }
//
// PUT my-index-000001/_doc/1
// {
//   "city": "New York"
// }
//
// PUT my-index-000001/_doc/2
// {
//   "city": "York"
// }
//
// GET my-index-000001/_search
// {
//   "query": {
//     "match": {
//       "city": "york" <2>
//     }
//   },
//   "sort": {
//     "city.raw": "asc" <3>
//   },
//   "aggs": {
//     "Cities": {
//       "terms": {
//         "field": "city.raw" <3>
//       }
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_mapping_params_multi_fields_6aa2941855d13f365f70aa8767ecb137(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:6aa2941855d13f365f70aa8767ecb137[]
	{
		res, err := es.Indices.Create(
			"my-index-000001",
			es.Indices.Create.WithBody(strings.NewReader(`{
		  "mappings": {
		    "properties": {
		      "city": {
		        "type": "text",
		        "fields": {
		          "raw": {
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
	}

	{
		res, err := es.Index(
			"my-index-000001",
			strings.NewReader(`{
		  "city": "New York"
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
			"my-index-000001",
			strings.NewReader(`{
		  "city": "York"
		}`),
			es.Index.WithDocumentID("2"),
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
			es.Search.WithIndex("my-index-000001"),
			es.Search.WithBody(strings.NewReader(`{
		  "query": {
		    "match": {
		      "city": "york"
		    }
		  },
		  "sort": {
		    "city.raw": "asc"
		  },
		  "aggs": {
		    "Cities": {
		      "terms": {
		        "field": "city.raw"
		      }
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
	// end:6aa2941855d13f365f70aa8767ecb137[]
}
