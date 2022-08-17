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

// AutoscalingDecider type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/autoscaling/get_autoscaling_capacity/GetAutoscalingCapacityResponse.ts#L52-L56
type AutoscalingDecider struct {
	ReasonDetails    interface{}         `json:"reason_details,omitempty"`
	ReasonSummary    *string             `json:"reason_summary,omitempty"`
	RequiredCapacity AutoscalingCapacity `json:"required_capacity"`
}

// AutoscalingDeciderBuilder holds AutoscalingDecider struct and provides a builder API.
type AutoscalingDeciderBuilder struct {
	v *AutoscalingDecider
}

// NewAutoscalingDecider provides a builder for the AutoscalingDecider struct.
func NewAutoscalingDeciderBuilder() *AutoscalingDeciderBuilder {
	r := AutoscalingDeciderBuilder{
		&AutoscalingDecider{},
	}

	return &r
}

// Build finalize the chain and returns the AutoscalingDecider struct
func (rb *AutoscalingDeciderBuilder) Build() AutoscalingDecider {
	return *rb.v
}

func (rb *AutoscalingDeciderBuilder) ReasonDetails(reasondetails interface{}) *AutoscalingDeciderBuilder {
	rb.v.ReasonDetails = reasondetails
	return rb
}

func (rb *AutoscalingDeciderBuilder) ReasonSummary(reasonsummary string) *AutoscalingDeciderBuilder {
	rb.v.ReasonSummary = &reasonsummary
	return rb
}

func (rb *AutoscalingDeciderBuilder) RequiredCapacity(requiredcapacity *AutoscalingCapacityBuilder) *AutoscalingDeciderBuilder {
	v := requiredcapacity.Build()
	rb.v.RequiredCapacity = v
	return rb
}
