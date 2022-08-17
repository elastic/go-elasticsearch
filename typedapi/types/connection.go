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

// Connection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/graph/_types/Connection.ts#L22-L27
type Connection struct {
	DocCount int64   `json:"doc_count"`
	Source   int64   `json:"source"`
	Target   int64   `json:"target"`
	Weight   float64 `json:"weight"`
}

// ConnectionBuilder holds Connection struct and provides a builder API.
type ConnectionBuilder struct {
	v *Connection
}

// NewConnection provides a builder for the Connection struct.
func NewConnectionBuilder() *ConnectionBuilder {
	r := ConnectionBuilder{
		&Connection{},
	}

	return &r
}

// Build finalize the chain and returns the Connection struct
func (rb *ConnectionBuilder) Build() Connection {
	return *rb.v
}

func (rb *ConnectionBuilder) DocCount(doccount int64) *ConnectionBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *ConnectionBuilder) Source(source int64) *ConnectionBuilder {
	rb.v.Source = source
	return rb
}

func (rb *ConnectionBuilder) Target(target int64) *ConnectionBuilder {
	rb.v.Target = target
	return rb
}

func (rb *ConnectionBuilder) Weight(weight float64) *ConnectionBuilder {
	rb.v.Weight = weight
	return rb
}
