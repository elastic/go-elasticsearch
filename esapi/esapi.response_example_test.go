// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package esapi_test

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func ExampleResponse_IsError() {
	es, _ := elasticsearch.NewDefaultClient()

	res, err := es.Info()

	// Handle connection errors
	//
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	defer res.Body.Close()

	// Handle error response (4xx, 5xx)
	//
	if res.IsError() {
		log.Fatalf("ERROR: %s", res.Status())
	}

	// Handle successful response (2xx)
	//
	log.Println(res)
}

func ExampleResponse_Status() {
	es, _ := elasticsearch.NewDefaultClient()

	res, _ := es.Info()
	log.Println(res.Status())

	// 200 OK
}

func ExampleResponse_String() {
	es, _ := elasticsearch.NewDefaultClient()

	res, _ := es.Info()
	log.Println(res.String())

	// [200 OK] {
	// "name" : "es1",
	// "cluster_name" : "go-elasticsearch",
	// ...
	// }
}
