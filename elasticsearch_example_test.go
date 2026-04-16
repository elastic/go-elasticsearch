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

//go:build !integration
// +build !integration

package elasticsearch_test

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
)

func init() {
	log.SetFlags(0)
}

func ExampleNew() {
	es, err := elasticsearch.New()
	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	log.Print(es.Transport.(*elastictransport.Client).URLs())
}

func ExampleNew_config() {
	es, err := elasticsearch.New(
		elasticsearch.WithAddresses("http://localhost:9200"),
		elasticsearch.WithBasicAuth("foo", "bar"),
		elasticsearch.WithTransportOptions(
			elastictransport.WithTransport(&http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Second,
				DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			}),
		),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
	}

	log.Print(es.Transport.(*elastictransport.Client).URLs())
}

func ExampleNew_logger() {
	// Use one of the bundled loggers:
	//
	// * elastictransport.TextLogger
	// * elastictransport.ColorLogger
	// * elastictransport.CurlLogger
	// * elastictransport.JSONLogger

	_, _ = elasticsearch.New(
		elasticsearch.WithLogger(&elastictransport.ColorLogger{Output: os.Stdout}),
	)
}
