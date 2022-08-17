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

// ScriptedMetricAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L137-L143
type ScriptedMetricAggregation struct {
	CombineScript *Script                `json:"combine_script,omitempty"`
	Field         *Field                 `json:"field,omitempty"`
	InitScript    *Script                `json:"init_script,omitempty"`
	MapScript     *Script                `json:"map_script,omitempty"`
	Missing       *Missing               `json:"missing,omitempty"`
	Params        map[string]interface{} `json:"params,omitempty"`
	ReduceScript  *Script                `json:"reduce_script,omitempty"`
	Script        *Script                `json:"script,omitempty"`
}

// ScriptedMetricAggregationBuilder holds ScriptedMetricAggregation struct and provides a builder API.
type ScriptedMetricAggregationBuilder struct {
	v *ScriptedMetricAggregation
}

// NewScriptedMetricAggregation provides a builder for the ScriptedMetricAggregation struct.
func NewScriptedMetricAggregationBuilder() *ScriptedMetricAggregationBuilder {
	r := ScriptedMetricAggregationBuilder{
		&ScriptedMetricAggregation{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ScriptedMetricAggregation struct
func (rb *ScriptedMetricAggregationBuilder) Build() ScriptedMetricAggregation {
	return *rb.v
}

func (rb *ScriptedMetricAggregationBuilder) CombineScript(combinescript *ScriptBuilder) *ScriptedMetricAggregationBuilder {
	v := combinescript.Build()
	rb.v.CombineScript = &v
	return rb
}

func (rb *ScriptedMetricAggregationBuilder) Field(field Field) *ScriptedMetricAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *ScriptedMetricAggregationBuilder) InitScript(initscript *ScriptBuilder) *ScriptedMetricAggregationBuilder {
	v := initscript.Build()
	rb.v.InitScript = &v
	return rb
}

func (rb *ScriptedMetricAggregationBuilder) MapScript(mapscript *ScriptBuilder) *ScriptedMetricAggregationBuilder {
	v := mapscript.Build()
	rb.v.MapScript = &v
	return rb
}

func (rb *ScriptedMetricAggregationBuilder) Missing(missing *MissingBuilder) *ScriptedMetricAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *ScriptedMetricAggregationBuilder) Params(value map[string]interface{}) *ScriptedMetricAggregationBuilder {
	rb.v.Params = value
	return rb
}

func (rb *ScriptedMetricAggregationBuilder) ReduceScript(reducescript *ScriptBuilder) *ScriptedMetricAggregationBuilder {
	v := reducescript.Build()
	rb.v.ReduceScript = &v
	return rb
}

func (rb *ScriptedMetricAggregationBuilder) Script(script *ScriptBuilder) *ScriptedMetricAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
