/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Package client is the starting point to establish a connection to Elasticsearch, as well as to invoke any of the
// available APIs. Note that API is an embedded field in Client, hence its documentation can be found in the api
// package.
//
// Here's a quick example:
//	c, _ := client.New(client.WithHost("https://elasticsearch:9200"))
//	body := map[string]interface{}{
//		"query": map[string]interface{}{
//			"term": map[string]interface{}{
//				"user": "kimchy",
//			},
//		},
//	}
//	resp, err := c.Search(body)
//	// ...
package client

import (
	"github.com/elastic/go-elasticsearch/api"
	"github.com/elastic/go-elasticsearch/transport"
)

// TODO: we should find a trick to link to api here

// Client is the top-level client.
type Client struct {
	// API is the root of the API. Since it's embedded, it allows to call APIs directly on the client.
	//	client, _ := client.New()
	//	resp, err := client.Search(...)
	//	...
	*api.API

	transport *transport.Transport
}

// New is the constructor for the Client.
func New(options ...Option) (*Client, error) {
	t := transport.New()
	c := &Client{
		API:       api.New(t),
		transport: t,
	}
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
