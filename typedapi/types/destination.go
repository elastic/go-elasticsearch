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

// Destination type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/_types/Transform.ts#L34-L45
type Destination struct {
	// Index The destination index for the transform. The mappings of the destination
	// index are deduced based on the source
	// fields when possible. If alternate mappings are required, use the create
	// index API prior to starting the
	// transform.
	Index *IndexName `json:"index,omitempty"`
	// Pipeline The unique identifier for an ingest pipeline.
	Pipeline *string `json:"pipeline,omitempty"`
}

// DestinationBuilder holds Destination struct and provides a builder API.
type DestinationBuilder struct {
	v *Destination
}

// NewDestination provides a builder for the Destination struct.
func NewDestinationBuilder() *DestinationBuilder {
	r := DestinationBuilder{
		&Destination{},
	}

	return &r
}

// Build finalize the chain and returns the Destination struct
func (rb *DestinationBuilder) Build() Destination {
	return *rb.v
}

// Index The destination index for the transform. The mappings of the destination
// index are deduced based on the source
// fields when possible. If alternate mappings are required, use the create
// index API prior to starting the
// transform.

func (rb *DestinationBuilder) Index(index IndexName) *DestinationBuilder {
	rb.v.Index = &index
	return rb
}

// Pipeline The unique identifier for an ingest pipeline.

func (rb *DestinationBuilder) Pipeline(pipeline string) *DestinationBuilder {
	rb.v.Pipeline = &pipeline
	return rb
}
