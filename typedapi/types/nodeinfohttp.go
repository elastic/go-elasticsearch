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

// NodeInfoHttp type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L295-L300
type NodeInfoHttp struct {
	BoundAddress            []string  `json:"bound_address"`
	MaxContentLength        *ByteSize `json:"max_content_length,omitempty"`
	MaxContentLengthInBytes int64     `json:"max_content_length_in_bytes"`
	PublishAddress          string    `json:"publish_address"`
}

// NodeInfoHttpBuilder holds NodeInfoHttp struct and provides a builder API.
type NodeInfoHttpBuilder struct {
	v *NodeInfoHttp
}

// NewNodeInfoHttp provides a builder for the NodeInfoHttp struct.
func NewNodeInfoHttpBuilder() *NodeInfoHttpBuilder {
	r := NodeInfoHttpBuilder{
		&NodeInfoHttp{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoHttp struct
func (rb *NodeInfoHttpBuilder) Build() NodeInfoHttp {
	return *rb.v
}

func (rb *NodeInfoHttpBuilder) BoundAddress(bound_address ...string) *NodeInfoHttpBuilder {
	rb.v.BoundAddress = bound_address
	return rb
}

func (rb *NodeInfoHttpBuilder) MaxContentLength(maxcontentlength *ByteSizeBuilder) *NodeInfoHttpBuilder {
	v := maxcontentlength.Build()
	rb.v.MaxContentLength = &v
	return rb
}

func (rb *NodeInfoHttpBuilder) MaxContentLengthInBytes(maxcontentlengthinbytes int64) *NodeInfoHttpBuilder {
	rb.v.MaxContentLengthInBytes = maxcontentlengthinbytes
	return rb
}

func (rb *NodeInfoHttpBuilder) PublishAddress(publishaddress string) *NodeInfoHttpBuilder {
	rb.v.PublishAddress = publishaddress
	return rb
}
