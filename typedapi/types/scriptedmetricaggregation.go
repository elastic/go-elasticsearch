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
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


package types

// ScriptedMetricAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_types/aggregations/metric.ts#L137-L143
type ScriptedMetricAggregation struct {
	CombineScript *Script                `json:"combine_script,omitempty"`
	Field         *string                `json:"field,omitempty"`
	InitScript    *Script                `json:"init_script,omitempty"`
	MapScript     *Script                `json:"map_script,omitempty"`
	Missing       *Missing               `json:"missing,omitempty"`
	Params        map[string]interface{} `json:"params,omitempty"`
	ReduceScript  *Script                `json:"reduce_script,omitempty"`
	Script        *Script                `json:"script,omitempty"`
}

// NewScriptedMetricAggregation returns a ScriptedMetricAggregation.
func NewScriptedMetricAggregation() *ScriptedMetricAggregation {
	r := &ScriptedMetricAggregation{
		Params: make(map[string]interface{}, 0),
	}

	return r
}
