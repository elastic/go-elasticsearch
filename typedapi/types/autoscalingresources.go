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

// AutoscalingResources type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/autoscaling/get_autoscaling_capacity/GetAutoscalingCapacityResponse.ts#L43-L46
type AutoscalingResources struct {
	Memory  int `json:"memory"`
	Storage int `json:"storage"`
}

// AutoscalingResourcesBuilder holds AutoscalingResources struct and provides a builder API.
type AutoscalingResourcesBuilder struct {
	v *AutoscalingResources
}

// NewAutoscalingResources provides a builder for the AutoscalingResources struct.
func NewAutoscalingResourcesBuilder() *AutoscalingResourcesBuilder {
	r := AutoscalingResourcesBuilder{
		&AutoscalingResources{},
	}

	return &r
}

// Build finalize the chain and returns the AutoscalingResources struct
func (rb *AutoscalingResourcesBuilder) Build() AutoscalingResources {
	return *rb.v
}

func (rb *AutoscalingResourcesBuilder) Memory(memory int) *AutoscalingResourcesBuilder {
	rb.v.Memory = memory
	return rb
}

func (rb *AutoscalingResourcesBuilder) Storage(storage int) *AutoscalingResourcesBuilder {
	rb.v.Storage = storage
	return rb
}
