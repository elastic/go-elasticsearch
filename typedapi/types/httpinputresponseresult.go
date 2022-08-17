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


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// HttpInputResponseResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L302-L306
type HttpInputResponseResult struct {
	Body    string      `json:"body"`
	Headers HttpHeaders `json:"headers"`
	Status  int         `json:"status"`
}

// HttpInputResponseResultBuilder holds HttpInputResponseResult struct and provides a builder API.
type HttpInputResponseResultBuilder struct {
	v *HttpInputResponseResult
}

// NewHttpInputResponseResult provides a builder for the HttpInputResponseResult struct.
func NewHttpInputResponseResultBuilder() *HttpInputResponseResultBuilder {
	r := HttpInputResponseResultBuilder{
		&HttpInputResponseResult{},
	}

	return &r
}

// Build finalize the chain and returns the HttpInputResponseResult struct
func (rb *HttpInputResponseResultBuilder) Build() HttpInputResponseResult {
	return *rb.v
}

func (rb *HttpInputResponseResultBuilder) Body(body string) *HttpInputResponseResultBuilder {
	rb.v.Body = body
	return rb
}

func (rb *HttpInputResponseResultBuilder) Headers(headers *HttpHeadersBuilder) *HttpInputResponseResultBuilder {
	v := headers.Build()
	rb.v.Headers = v
	return rb
}

func (rb *HttpInputResponseResultBuilder) Status(status int) *HttpInputResponseResultBuilder {
	rb.v.Status = status
	return rb
}
