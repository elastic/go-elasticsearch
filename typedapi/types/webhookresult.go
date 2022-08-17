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

// WebhookResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L295-L298
type WebhookResult struct {
	Request  HttpInputRequestResult   `json:"request"`
	Response *HttpInputResponseResult `json:"response,omitempty"`
}

// WebhookResultBuilder holds WebhookResult struct and provides a builder API.
type WebhookResultBuilder struct {
	v *WebhookResult
}

// NewWebhookResult provides a builder for the WebhookResult struct.
func NewWebhookResultBuilder() *WebhookResultBuilder {
	r := WebhookResultBuilder{
		&WebhookResult{},
	}

	return &r
}

// Build finalize the chain and returns the WebhookResult struct
func (rb *WebhookResultBuilder) Build() WebhookResult {
	return *rb.v
}

func (rb *WebhookResultBuilder) Request(request *HttpInputRequestResultBuilder) *WebhookResultBuilder {
	v := request.Build()
	rb.v.Request = v
	return rb
}

func (rb *WebhookResultBuilder) Response(response *HttpInputResponseResultBuilder) *WebhookResultBuilder {
	v := response.Build()
	rb.v.Response = &v
	return rb
}
