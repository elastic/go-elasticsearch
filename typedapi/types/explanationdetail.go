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

// ExplanationDetail type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/explain/types.ts#L28-L32
type ExplanationDetail struct {
	Description string              `json:"description"`
	Details     []ExplanationDetail `json:"details,omitempty"`
	Value       float32             `json:"value"`
}

// ExplanationDetailBuilder holds ExplanationDetail struct and provides a builder API.
type ExplanationDetailBuilder struct {
	v *ExplanationDetail
}

// NewExplanationDetail provides a builder for the ExplanationDetail struct.
func NewExplanationDetailBuilder() *ExplanationDetailBuilder {
	r := ExplanationDetailBuilder{
		&ExplanationDetail{},
	}

	return &r
}

// Build finalize the chain and returns the ExplanationDetail struct
func (rb *ExplanationDetailBuilder) Build() ExplanationDetail {
	return *rb.v
}

func (rb *ExplanationDetailBuilder) Description(description string) *ExplanationDetailBuilder {
	rb.v.Description = description
	return rb
}

func (rb *ExplanationDetailBuilder) Details(details []ExplanationDetailBuilder) *ExplanationDetailBuilder {
	tmp := make([]ExplanationDetail, len(details))
	for _, value := range details {
		tmp = append(tmp, value.Build())
	}
	rb.v.Details = tmp
	return rb
}

func (rb *ExplanationDetailBuilder) Value(value float32) *ExplanationDetailBuilder {
	rb.v.Value = value
	return rb
}
