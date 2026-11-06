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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// LearningToRankRescoreProfileBreakdown type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/search/_types/profile.ts#L221-L231
type LearningToRankRescoreProfileBreakdown struct {
	// BuildFeatureExtractors The time spent turning the feature queries of the model into queries that can
	// score documents.
	BuildFeatureExtractors      int64 `json:"build_feature_extractors"`
	BuildFeatureExtractorsCount int64 `json:"build_feature_extractors_count"`
	// FeatureExtraction The time spent extracting the features of the top documents.
	FeatureExtraction      int64 `json:"feature_extraction"`
	FeatureExtractionCount int64 `json:"feature_extraction_count"`
	// Inference The time spent running the model on the extracted features.
	Inference      int64 `json:"inference"`
	InferenceCount int64 `json:"inference_count"`
}

func (s *LearningToRankRescoreProfileBreakdown) UnmarshalJSON(data []byte) error {

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

		case "build_feature_extractors":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildFeatureExtractors", err)
				}
				s.BuildFeatureExtractors = value
			case float64:
				f := int64(v)
				s.BuildFeatureExtractors = f
			}

		case "build_feature_extractors_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildFeatureExtractorsCount", err)
				}
				s.BuildFeatureExtractorsCount = value
			case float64:
				f := int64(v)
				s.BuildFeatureExtractorsCount = f
			}

		case "feature_extraction":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FeatureExtraction", err)
				}
				s.FeatureExtraction = value
			case float64:
				f := int64(v)
				s.FeatureExtraction = f
			}

		case "feature_extraction_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FeatureExtractionCount", err)
				}
				s.FeatureExtractionCount = value
			case float64:
				f := int64(v)
				s.FeatureExtractionCount = f
			}

		case "inference":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Inference", err)
				}
				s.Inference = value
			case float64:
				f := int64(v)
				s.Inference = f
			}

		case "inference_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "InferenceCount", err)
				}
				s.InferenceCount = value
			case float64:
				f := int64(v)
				s.InferenceCount = f
			}

		}
	}
	return nil
}

// NewLearningToRankRescoreProfileBreakdown returns a LearningToRankRescoreProfileBreakdown.
func NewLearningToRankRescoreProfileBreakdown() *LearningToRankRescoreProfileBreakdown {
	r := &LearningToRankRescoreProfileBreakdown{}

	return r
}
