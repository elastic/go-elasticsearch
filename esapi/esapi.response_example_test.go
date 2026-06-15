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

package esapi_test

import (
	"log"

	"github.com/elastic/go-elasticsearch/v9"
)

func ExampleResponse_IsError() {
	es, _ := elasticsearch.New()

	res, err := es.Info()

	// Handle connection errors
	//
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	defer func() { _ = res.Body.Close() }()

	// Handle error response (3xx, 4xx, 5xx)
	//
	if res.IsError() {
		log.Printf("ERROR: %s", res.Status())
		return
	}

	// Handle successful response (2xx)
	//
	log.Println(res)
}

func ExampleResponse_Status() {
	es, _ := elasticsearch.New()

	res, _ := es.Info()
	log.Println(res.Status())

	// 200 OK
}

func ExampleResponse_String() {
	es, _ := elasticsearch.New()

	res, _ := es.Info()
	log.Println(res.String())

	// [200 OK] {
	// "name" : "es1",
	// "cluster_name" : "go-elasticsearch",
	// ...
	// }
}
