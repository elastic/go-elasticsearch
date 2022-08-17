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

// ScriptCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Conditions.ts#L77-L85
type ScriptCondition struct {
	Id     *string                `json:"id,omitempty"`
	Lang   *string                `json:"lang,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	Source *string                `json:"source,omitempty"`
}

// ScriptConditionBuilder holds ScriptCondition struct and provides a builder API.
type ScriptConditionBuilder struct {
	v *ScriptCondition
}

// NewScriptCondition provides a builder for the ScriptCondition struct.
func NewScriptConditionBuilder() *ScriptConditionBuilder {
	r := ScriptConditionBuilder{
		&ScriptCondition{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ScriptCondition struct
func (rb *ScriptConditionBuilder) Build() ScriptCondition {
	return *rb.v
}

func (rb *ScriptConditionBuilder) Id(id string) *ScriptConditionBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *ScriptConditionBuilder) Lang(lang string) *ScriptConditionBuilder {
	rb.v.Lang = &lang
	return rb
}

func (rb *ScriptConditionBuilder) Params(value map[string]interface{}) *ScriptConditionBuilder {
	rb.v.Params = value
	return rb
}

func (rb *ScriptConditionBuilder) Source(source string) *ScriptConditionBuilder {
	rb.v.Source = &source
	return rb
}
