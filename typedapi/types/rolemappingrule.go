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

// RoleMappingRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/_types/RoleMappingRule.ts#L23-L31
type RoleMappingRule struct {
	All    []RoleMappingRule `json:"all,omitempty"`
	Any    []RoleMappingRule `json:"any,omitempty"`
	Except *RoleMappingRule  `json:"except,omitempty"`
	Field  *FieldRule        `json:"field,omitempty"`
}

// RoleMappingRuleBuilder holds RoleMappingRule struct and provides a builder API.
type RoleMappingRuleBuilder struct {
	v *RoleMappingRule
}

// NewRoleMappingRule provides a builder for the RoleMappingRule struct.
func NewRoleMappingRuleBuilder() *RoleMappingRuleBuilder {
	r := RoleMappingRuleBuilder{
		&RoleMappingRule{},
	}

	return &r
}

// Build finalize the chain and returns the RoleMappingRule struct
func (rb *RoleMappingRuleBuilder) Build() RoleMappingRule {
	return *rb.v
}

func (rb *RoleMappingRuleBuilder) All(all []RoleMappingRuleBuilder) *RoleMappingRuleBuilder {
	tmp := make([]RoleMappingRule, len(all))
	for _, value := range all {
		tmp = append(tmp, value.Build())
	}
	rb.v.All = tmp
	return rb
}

func (rb *RoleMappingRuleBuilder) Any(any []RoleMappingRuleBuilder) *RoleMappingRuleBuilder {
	tmp := make([]RoleMappingRule, len(any))
	for _, value := range any {
		tmp = append(tmp, value.Build())
	}
	rb.v.Any = tmp
	return rb
}

func (rb *RoleMappingRuleBuilder) Except(except *RoleMappingRuleBuilder) *RoleMappingRuleBuilder {
	v := except.Build()
	rb.v.Except = &v
	return rb
}

func (rb *RoleMappingRuleBuilder) Field(field *FieldRuleBuilder) *RoleMappingRuleBuilder {
	v := field.Build()
	rb.v.Field = &v
	return rb
}
