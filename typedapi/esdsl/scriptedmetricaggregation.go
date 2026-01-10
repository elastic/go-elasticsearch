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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _scriptedMetricAggregation struct {
	v *types.ScriptedMetricAggregation
}

// A metric aggregation that uses scripts to provide a metric output.
func NewScriptedMetricAggregation() *_scriptedMetricAggregation {

	return &_scriptedMetricAggregation{v: types.NewScriptedMetricAggregation()}

}

func (s *_scriptedMetricAggregation) CombineScript(combinescript types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.CombineScript = combinescript.ScriptCaster()

	return s
}

func (s *_scriptedMetricAggregation) InitScript(initscript types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.InitScript = initscript.ScriptCaster()

	return s
}

func (s *_scriptedMetricAggregation) MapScript(mapscript types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.MapScript = mapscript.ScriptCaster()

	return s
}

func (s *_scriptedMetricAggregation) Params(params map[string]json.RawMessage) *_scriptedMetricAggregation {

	s.v.Params = params
	return s
}

func (s *_scriptedMetricAggregation) AddParam(key string, value json.RawMessage) *_scriptedMetricAggregation {

	var tmp map[string]json.RawMessage
	if s.v.Params == nil {
		s.v.Params = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

func (s *_scriptedMetricAggregation) ReduceScript(reducescript types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.ReduceScript = reducescript.ScriptCaster()

	return s
}

func (s *_scriptedMetricAggregation) Field(field string) *_scriptedMetricAggregation {

	s.v.Field = &field

	return s
}

func (s *_scriptedMetricAggregation) Missing(missing types.MissingVariant) *_scriptedMetricAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_scriptedMetricAggregation) Script(script types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_scriptedMetricAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.ScriptedMetric = s.v

	return container
}

func (s *_scriptedMetricAggregation) ScriptedMetricAggregationCaster() *types.ScriptedMetricAggregation {
	return s.v
}
