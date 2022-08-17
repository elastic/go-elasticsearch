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

// AutoscalingDeciders type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/autoscaling/get_autoscaling_capacity/GetAutoscalingCapacityResponse.ts#L31-L36
type AutoscalingDeciders struct {
	CurrentCapacity  AutoscalingCapacity           `json:"current_capacity"`
	CurrentNodes     []AutoscalingNode             `json:"current_nodes"`
	Deciders         map[string]AutoscalingDecider `json:"deciders"`
	RequiredCapacity AutoscalingCapacity           `json:"required_capacity"`
}

// AutoscalingDecidersBuilder holds AutoscalingDeciders struct and provides a builder API.
type AutoscalingDecidersBuilder struct {
	v *AutoscalingDeciders
}

// NewAutoscalingDeciders provides a builder for the AutoscalingDeciders struct.
func NewAutoscalingDecidersBuilder() *AutoscalingDecidersBuilder {
	r := AutoscalingDecidersBuilder{
		&AutoscalingDeciders{
			Deciders: make(map[string]AutoscalingDecider, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AutoscalingDeciders struct
func (rb *AutoscalingDecidersBuilder) Build() AutoscalingDeciders {
	return *rb.v
}

func (rb *AutoscalingDecidersBuilder) CurrentCapacity(currentcapacity *AutoscalingCapacityBuilder) *AutoscalingDecidersBuilder {
	v := currentcapacity.Build()
	rb.v.CurrentCapacity = v
	return rb
}

func (rb *AutoscalingDecidersBuilder) CurrentNodes(current_nodes []AutoscalingNodeBuilder) *AutoscalingDecidersBuilder {
	tmp := make([]AutoscalingNode, len(current_nodes))
	for _, value := range current_nodes {
		tmp = append(tmp, value.Build())
	}
	rb.v.CurrentNodes = tmp
	return rb
}

func (rb *AutoscalingDecidersBuilder) Deciders(values map[string]*AutoscalingDeciderBuilder) *AutoscalingDecidersBuilder {
	tmp := make(map[string]AutoscalingDecider, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Deciders = tmp
	return rb
}

func (rb *AutoscalingDecidersBuilder) RequiredCapacity(requiredcapacity *AutoscalingCapacityBuilder) *AutoscalingDecidersBuilder {
	v := requiredcapacity.Build()
	rb.v.RequiredCapacity = v
	return rb
}
