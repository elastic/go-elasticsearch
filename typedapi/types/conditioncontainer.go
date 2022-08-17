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

// ConditionContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Conditions.ts#L51-L60
type ConditionContainer struct {
	Always       *AlwaysCondition       `json:"always,omitempty"`
	ArrayCompare *ArrayCompareCondition `json:"array_compare,omitempty"`
	Compare      *CompareCondition      `json:"compare,omitempty"`
	Never        *NeverCondition        `json:"never,omitempty"`
	Script       *ScriptCondition       `json:"script,omitempty"`
}

// ConditionContainerBuilder holds ConditionContainer struct and provides a builder API.
type ConditionContainerBuilder struct {
	v *ConditionContainer
}

// NewConditionContainer provides a builder for the ConditionContainer struct.
func NewConditionContainerBuilder() *ConditionContainerBuilder {
	r := ConditionContainerBuilder{
		&ConditionContainer{},
	}

	return &r
}

// Build finalize the chain and returns the ConditionContainer struct
func (rb *ConditionContainerBuilder) Build() ConditionContainer {
	return *rb.v
}

func (rb *ConditionContainerBuilder) Always(always *AlwaysConditionBuilder) *ConditionContainerBuilder {
	v := always.Build()
	rb.v.Always = &v
	return rb
}

func (rb *ConditionContainerBuilder) ArrayCompare(arraycompare *ArrayCompareConditionBuilder) *ConditionContainerBuilder {
	v := arraycompare.Build()
	rb.v.ArrayCompare = &v
	return rb
}

func (rb *ConditionContainerBuilder) Compare(compare *CompareConditionBuilder) *ConditionContainerBuilder {
	v := compare.Build()
	rb.v.Compare = &v
	return rb
}

func (rb *ConditionContainerBuilder) Never(never *NeverConditionBuilder) *ConditionContainerBuilder {
	v := never.Build()
	rb.v.Never = &v
	return rb
}

func (rb *ConditionContainerBuilder) Script(script *ScriptConditionBuilder) *ConditionContainerBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
