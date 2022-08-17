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

// Acknowledgement type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/license/post/types.ts#L20-L23
type Acknowledgement struct {
	License []string `json:"license"`
	Message string   `json:"message"`
}

// AcknowledgementBuilder holds Acknowledgement struct and provides a builder API.
type AcknowledgementBuilder struct {
	v *Acknowledgement
}

// NewAcknowledgement provides a builder for the Acknowledgement struct.
func NewAcknowledgementBuilder() *AcknowledgementBuilder {
	r := AcknowledgementBuilder{
		&Acknowledgement{},
	}

	return &r
}

// Build finalize the chain and returns the Acknowledgement struct
func (rb *AcknowledgementBuilder) Build() Acknowledgement {
	return *rb.v
}

func (rb *AcknowledgementBuilder) License(license ...string) *AcknowledgementBuilder {
	rb.v.License = license
	return rb
}

func (rb *AcknowledgementBuilder) Message(message string) *AcknowledgementBuilder {
	rb.v.Message = message
	return rb
}
