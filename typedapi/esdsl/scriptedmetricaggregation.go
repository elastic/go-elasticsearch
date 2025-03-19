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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _scriptedMetricAggregation struct {
	v *types.ScriptedMetricAggregation
}

// A metric aggregation that uses scripts to provide a metric output.
func NewScriptedMetricAggregation() *_scriptedMetricAggregation {

	return &_scriptedMetricAggregation{v: types.NewScriptedMetricAggregation()}

}

// Runs once on each shard after document collection is complete.
// Allows the aggregation to consolidate the state returned from each shard.
func (s *_scriptedMetricAggregation) CombineScript(combinescript types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.CombineScript = combinescript.ScriptCaster()

	return s
}

// The field on which to run the aggregation.
func (s *_scriptedMetricAggregation) Field(field string) *_scriptedMetricAggregation {

	s.v.Field = &field

	return s
}

// Runs prior to any collection of documents.
// Allows the aggregation to set up any initial state.
func (s *_scriptedMetricAggregation) InitScript(initscript types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.InitScript = initscript.ScriptCaster()

	return s
}

// Run once per document collected.
// If no `combine_script` is specified, the resulting state needs to be stored
// in the `state` object.
func (s *_scriptedMetricAggregation) MapScript(mapscript types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.MapScript = mapscript.ScriptCaster()

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_scriptedMetricAggregation) Missing(missing types.MissingVariant) *_scriptedMetricAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

// A global object with script parameters for `init`, `map` and `combine`
// scripts.
// It is shared between the scripts.
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

// Runs once on the coordinating node after all shards have returned their
// results.
// The script is provided with access to a variable `states`, which is an array
// of the result of the `combine_script` on each shard.
func (s *_scriptedMetricAggregation) ReduceScript(reducescript types.ScriptVariant) *_scriptedMetricAggregation {

	s.v.ReduceScript = reducescript.ScriptCaster()

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
