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

// CompareCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Conditions.ts#L33-L39
type CompareCondition struct {
	Comparison      *string                         `json:"comparison,omitempty"`
	CtxPayloadMatch *CompareContextPayloadCondition `json:"ctx.payload.match,omitempty"`
	CtxPayloadValue *CompareContextPayloadCondition `json:"ctx.payload.value,omitempty"`
	Path            *string                         `json:"path,omitempty"`
	Value           interface{}                     `json:"value,omitempty"`
}

// CompareConditionBuilder holds CompareCondition struct and provides a builder API.
type CompareConditionBuilder struct {
	v *CompareCondition
}

// NewCompareCondition provides a builder for the CompareCondition struct.
func NewCompareConditionBuilder() *CompareConditionBuilder {
	r := CompareConditionBuilder{
		&CompareCondition{},
	}

	return &r
}

// Build finalize the chain and returns the CompareCondition struct
func (rb *CompareConditionBuilder) Build() CompareCondition {
	return *rb.v
}

func (rb *CompareConditionBuilder) Comparison(comparison string) *CompareConditionBuilder {
	rb.v.Comparison = &comparison
	return rb
}

func (rb *CompareConditionBuilder) CtxPayloadMatch(ctxpayloadmatch *CompareContextPayloadConditionBuilder) *CompareConditionBuilder {
	v := ctxpayloadmatch.Build()
	rb.v.CtxPayloadMatch = &v
	return rb
}

func (rb *CompareConditionBuilder) CtxPayloadValue(ctxpayloadvalue *CompareContextPayloadConditionBuilder) *CompareConditionBuilder {
	v := ctxpayloadvalue.Build()
	rb.v.CtxPayloadValue = &v
	return rb
}

func (rb *CompareConditionBuilder) Path(path string) *CompareConditionBuilder {
	rb.v.Path = &path
	return rb
}

func (rb *CompareConditionBuilder) Value(value interface{}) *CompareConditionBuilder {
	rb.v.Value = value
	return rb
}
