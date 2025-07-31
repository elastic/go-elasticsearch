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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// LearningToRankConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/inference.ts#L87-L91
type LearningToRankConfig struct {
	DefaultParams                 map[string]json.RawMessage         `json:"default_params,omitempty"`
	FeatureExtractors             []map[string]QueryFeatureExtractor `json:"feature_extractors,omitempty"`
	NumTopFeatureImportanceValues int                                `json:"num_top_feature_importance_values"`
}

func (s *LearningToRankConfig) UnmarshalJSON(data []byte) error {

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

		case "default_params":
			if s.DefaultParams == nil {
				s.DefaultParams = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.DefaultParams); err != nil {
				return fmt.Errorf("%s | %w", "DefaultParams", err)
			}

		case "feature_extractors":
			if err := dec.Decode(&s.FeatureExtractors); err != nil {
				return fmt.Errorf("%s | %w", "FeatureExtractors", err)
			}

		case "num_top_feature_importance_values":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumTopFeatureImportanceValues", err)
				}
				s.NumTopFeatureImportanceValues = value
			case float64:
				f := int(v)
				s.NumTopFeatureImportanceValues = f
			}

		}
	}
	return nil
}

// NewLearningToRankConfig returns a LearningToRankConfig.
func NewLearningToRankConfig() *LearningToRankConfig {
	r := &LearningToRankConfig{
		DefaultParams: make(map[string]json.RawMessage),
	}

	return r
}
