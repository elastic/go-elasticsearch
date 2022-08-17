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

// PointInTimeReference type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/PointInTimeReference.ts#L23-L26
type PointInTimeReference struct {
	Id        Id        `json:"id"`
	KeepAlive *Duration `json:"keep_alive,omitempty"`
}

// PointInTimeReferenceBuilder holds PointInTimeReference struct and provides a builder API.
type PointInTimeReferenceBuilder struct {
	v *PointInTimeReference
}

// NewPointInTimeReference provides a builder for the PointInTimeReference struct.
func NewPointInTimeReferenceBuilder() *PointInTimeReferenceBuilder {
	r := PointInTimeReferenceBuilder{
		&PointInTimeReference{},
	}

	return &r
}

// Build finalize the chain and returns the PointInTimeReference struct
func (rb *PointInTimeReferenceBuilder) Build() PointInTimeReference {
	return *rb.v
}

func (rb *PointInTimeReferenceBuilder) Id(id Id) *PointInTimeReferenceBuilder {
	rb.v.Id = id
	return rb
}

func (rb *PointInTimeReferenceBuilder) KeepAlive(keepalive *DurationBuilder) *PointInTimeReferenceBuilder {
	v := keepalive.Build()
	rb.v.KeepAlive = &v
	return rb
}
