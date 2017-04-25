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
	m := newIndexMethod(t)
	p, err := newAPIPackage(m)
	if err != nil {
		t.Fatal(err)
	}
	var writer bytes.Buffer
	err = p.generateOption(filepath.Join("..", templatesDir), &writer)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"
)

// OpType - explicit operation type.
type OpType int
const (
	// OpTypeIndex can be used to set OpType to "index"
	OpTypeIndex = iota
	// OpTypeCreate can be used to set OpType to "create"
	OpTypeCreate = iota
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option struct {
	name string
	apply func(r *http.Request)
}

// WithID document ID.
func WithID(id string) *Option {
	return &Option{
		name: "WithID",
		apply: func(r *http.Request) {
		},
	}
}

// WithOpType explicit operation type.
func WithOpType(opType OpType) *Option {
	return &Option{
		name: "WithOpType",
		apply: func(r *http.Request) {
		},
	}
}

// WithTimeout explicit operation timeout.
func WithTimeout(timeout time.Time) *Option {
	return &Option{
		name: "WithTimeout",
		apply: func(r *http.Request) {
		},
	}
}

// WithVersion explicit version number for concurrency control.
func WithVersion(version int) *Option {
	return &Option{
		name: "WithVersion",
		apply: func(r *http.Request) {
		},
	}
}

// WithWaitForActiveShards sets the number of shard copies that must be active before proceeding with the index operation. Defaults to 1, meaning the primary shard only. Set to "all" for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas + 1).
func WithWaitForActiveShards(waitForActiveShards string) *Option {
	return &Option{
		name: "WithWaitForActiveShards",
		apply: func(r *http.Request) {
		},
	}
}


`
	if d := diff(t, expectedCode, writer.String()); len(d) > 0 {
		t.Fail()
	}
}

func TestGenerateAPI(t *testing.T) {
	m := newIndexMethod(t)
	p, err := newAPIPackage(m)
	if err != nil {
		t.Fatal(err)
	}
	var writer bytes.Buffer
	err = p.generateAPI(filepath.Join("..", templatesDir), &writer)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

// Package api is the root API package.
package api

import (
	"github.com/elastic/goelasticsearch/transport"
)

// API is the root API client.
type API struct {
	// transport is the transport client.
	transport *transport.Transport
}

// New is the constructor for API.
func New(transport *transport.Transport) *API {
	return &API{
		
		transport: transport,
	}
}
`
	if d := diff(t, expectedCode, writer.String()); len(d) > 0 {
		t.Fail()
	}
}
