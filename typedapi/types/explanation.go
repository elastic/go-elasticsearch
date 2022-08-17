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

// Explanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/explain/types.ts#L22-L26
type Explanation struct {
	Description string              `json:"description"`
	Details     []ExplanationDetail `json:"details"`
	Value       float32             `json:"value"`
}

// ExplanationBuilder holds Explanation struct and provides a builder API.
type ExplanationBuilder struct {
	v *Explanation
}

// NewExplanation provides a builder for the Explanation struct.
func NewExplanationBuilder() *ExplanationBuilder {
	r := ExplanationBuilder{
		&Explanation{},
	}

	return &r
}

// Build finalize the chain and returns the Explanation struct
func (rb *ExplanationBuilder) Build() Explanation {
	return *rb.v
}

func (rb *ExplanationBuilder) Description(description string) *ExplanationBuilder {
	rb.v.Description = description
	return rb
}

func (rb *ExplanationBuilder) Details(details []ExplanationDetailBuilder) *ExplanationBuilder {
	tmp := make([]ExplanationDetail, len(details))
	for _, value := range details {
		tmp = append(tmp, value.Build())
	}
	rb.v.Details = tmp
	return rb
}

func (rb *ExplanationBuilder) Value(value float32) *ExplanationBuilder {
	rb.v.Value = value
	return rb
}
