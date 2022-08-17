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

// Invocation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/slm/_types/SnapshotLifecycle.ts#L138-L141
type Invocation struct {
	SnapshotName Name     `json:"snapshot_name"`
	Time         DateTime `json:"time"`
}

// InvocationBuilder holds Invocation struct and provides a builder API.
type InvocationBuilder struct {
	v *Invocation
}

// NewInvocation provides a builder for the Invocation struct.
func NewInvocationBuilder() *InvocationBuilder {
	r := InvocationBuilder{
		&Invocation{},
	}

	return &r
}

// Build finalize the chain and returns the Invocation struct
func (rb *InvocationBuilder) Build() Invocation {
	return *rb.v
}

func (rb *InvocationBuilder) SnapshotName(snapshotname Name) *InvocationBuilder {
	rb.v.SnapshotName = snapshotname
	return rb
}

func (rb *InvocationBuilder) Time(time *DateTimeBuilder) *InvocationBuilder {
	v := time.Build()
	rb.v.Time = v
	return rb
}
