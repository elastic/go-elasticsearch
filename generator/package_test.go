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

package generator

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestGenerateOption(t *testing.T) {
	m, err := newMethod(filepath.Join("..", DefaultSpecDir, "index.json"))
	if err != nil {
		t.Fatal(err)
	}
	p, err := newAPIPackage(m)
	if err != nil {
		t.Fatal(err)
	}
	var writer bytes.Buffer
	err = p.generateOption(filepath.Join("..", templatesDir), &writer)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `// generated by github.com/elastic/elasticsearch-go/cmd/generator; DO NOT EDIT

package client

import (
	"net/http"
)

// Option is a non-required API option that gets applied to an HTTP request
type Option func(r *http.Request)

// WithID document ID
func WithID(id string) Option {
	return func(r *http.Request) {
	}
}

// WithOpType explicit operation type
func WithOpType(opType struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithParent ID of the parent document
func WithParent(parent string) Option {
	return func(r *http.Request) {
	}
}

// WithPipeline the pipeline id to preprocess incoming documents with
func WithPipeline(pipeline string) Option {
	return func(r *http.Request) {
	}
}

// WithRefresh if true then refresh the affected shards to make this operation visible to search, if wait_for then wait for a refresh to make this operation visible to search, if false (the default) then do nothing with refreshes.
func WithRefresh(refresh struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithRouting specific routing value
func WithRouting(routing string) Option {
	return func(r *http.Request) {
	}
}

// WithTimeout explicit operation timeout
func WithTimeout(timeout time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithTimestamp explicit timestamp for the document
func WithTimestamp(timestamp time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithTTL expiration time for the document
func WithTTL(ttl time.Time) Option {
	return func(r *http.Request) {
	}
}

// WithVersion explicit version number for concurrency control
func WithVersion(version int) Option {
	return func(r *http.Request) {
	}
}

// WithVersionType specific version type
func WithVersionType(versionType struct{}) Option {
	return func(r *http.Request) {
	}
}

// WithWaitForActiveShards sets the number of shard copies that must be active before proceeding with the index operation. Defaults to 1, meaning the primary shard only. Set to all for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas &#43; 1)
func WithWaitForActiveShards(waitForActiveShards string) Option {
	return func(r *http.Request) {
	}
}


`
	if d := diff(t, expectedCode, writer.String()); len(d) > 0 {
		t.Fail()
	}
}

func TestGenerateAPI(t *testing.T) {
	m, err := newMethod(filepath.Join("..", DefaultSpecDir, "index.json"))
	if err != nil {
		t.Fatal(err)
	}
	p, err := newAPIPackage(m)
	if err != nil {
		t.Fatal(err)
	}
	var writer bytes.Buffer
	err = p.generateAPI(filepath.Join("..", templatesDir), &writer)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `// generated by github.com/elastic/elasticsearch-go/cmd/generator; DO NOT EDIT

package client

import (
	"net/http"
)

// Client is the transport client.
type Client struct {

	// client is the transport client
	client *http.Client
}

// addClients adds the subclients
func (c *Client)addClients() {
}
`
	if d := diff(t, expectedCode, writer.String()); len(d) > 0 {
		t.Fail()
	}
}
