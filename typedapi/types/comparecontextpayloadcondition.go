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

// CompareContextPayloadCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Conditions.ts#L41-L47
type CompareContextPayloadCondition struct {
	Eq  interface{} `json:"eq,omitempty"`
	Gt  interface{} `json:"gt,omitempty"`
	Gte interface{} `json:"gte,omitempty"`
	Lt  interface{} `json:"lt,omitempty"`
	Lte interface{} `json:"lte,omitempty"`
}

// CompareContextPayloadConditionBuilder holds CompareContextPayloadCondition struct and provides a builder API.
type CompareContextPayloadConditionBuilder struct {
	v *CompareContextPayloadCondition
}

// NewCompareContextPayloadCondition provides a builder for the CompareContextPayloadCondition struct.
func NewCompareContextPayloadConditionBuilder() *CompareContextPayloadConditionBuilder {
	r := CompareContextPayloadConditionBuilder{
		&CompareContextPayloadCondition{},
	}

	return &r
}

// Build finalize the chain and returns the CompareContextPayloadCondition struct
func (rb *CompareContextPayloadConditionBuilder) Build() CompareContextPayloadCondition {
	return *rb.v
}

func (rb *CompareContextPayloadConditionBuilder) Eq(eq interface{}) *CompareContextPayloadConditionBuilder {
	rb.v.Eq = eq
	return rb
}

func (rb *CompareContextPayloadConditionBuilder) Gt(gt interface{}) *CompareContextPayloadConditionBuilder {
	rb.v.Gt = gt
	return rb
}

func (rb *CompareContextPayloadConditionBuilder) Gte(gte interface{}) *CompareContextPayloadConditionBuilder {
	rb.v.Gte = gte
	return rb
}

func (rb *CompareContextPayloadConditionBuilder) Lt(lt interface{}) *CompareContextPayloadConditionBuilder {
	rb.v.Lt = lt
	return rb
}

func (rb *CompareContextPayloadConditionBuilder) Lte(lte interface{}) *CompareContextPayloadConditionBuilder {
	rb.v.Lte = lte
	return rb
}
