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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// ScriptedMetricAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/metric.ts#L254-L280
type ScriptedMetricAggregation struct {
	// CombineScript Runs once on each shard after document collection is complete.
	// Allows the aggregation to consolidate the state returned from each shard.
	CombineScript Script `json:"combine_script,omitempty"`
	// Field The field on which to run the aggregation.
	Field *string `json:"field,omitempty"`
	// InitScript Runs prior to any collection of documents.
	// Allows the aggregation to set up any initial state.
	InitScript Script `json:"init_script,omitempty"`
	// MapScript Run once per document collected.
	// If no `combine_script` is specified, the resulting state needs to be stored
	// in the `state` object.
	MapScript Script `json:"map_script,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	// Params A global object with script parameters for `init`, `map` and `combine`
	// scripts.
	// It is shared between the scripts.
	Params map[string]json.RawMessage `json:"params,omitempty"`
	// ReduceScript Runs once on the coordinating node after all shards have returned their
	// results.
	// The script is provided with access to a variable `states`, which is an array
	// of the result of the `combine_script` on each shard.
	ReduceScript Script `json:"reduce_script,omitempty"`
	Script       Script `json:"script,omitempty"`
}

func (s *ScriptedMetricAggregation) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "combine_script":
			if err := dec.Decode(&s.CombineScript); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "init_script":
			if err := dec.Decode(&s.InitScript); err != nil {
				return err
			}

		case "map_script":
			if err := dec.Decode(&s.MapScript); err != nil {
				return err
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return err
			}

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return err
			}

		case "reduce_script":
			if err := dec.Decode(&s.ReduceScript); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewScriptedMetricAggregation returns a ScriptedMetricAggregation.
func NewScriptedMetricAggregation() *ScriptedMetricAggregation {
	r := &ScriptedMetricAggregation{
		Params: make(map[string]json.RawMessage, 0),
	}

	return r
}
