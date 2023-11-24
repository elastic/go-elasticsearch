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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TrainedModelInferenceStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/TrainedModel.ts#L104-L124
type TrainedModelInferenceStats struct {
	// CacheMissCount The number of times the model was loaded for inference and was not retrieved
	// from the cache.
	// If this number is close to the `inference_count`, the cache is not being
	// appropriately used.
	// This can be solved by increasing the cache size or its time-to-live (TTL).
	// Refer to general machine learning settings for the appropriate settings.
	CacheMissCount int `json:"cache_miss_count"`
	// FailureCount The number of failures when using the model for inference.
	FailureCount int `json:"failure_count"`
	// InferenceCount The total number of times the model has been called for inference.
	// This is across all inference contexts, including all pipelines.
	InferenceCount int `json:"inference_count"`
	// MissingAllFieldsCount The number of inference calls where all the training features for the model
	// were missing.
	MissingAllFieldsCount int `json:"missing_all_fields_count"`
	// Timestamp The time when the statistics were last updated.
	Timestamp DateTime `json:"timestamp"`
}

func (s *TrainedModelInferenceStats) UnmarshalJSON(data []byte) error {

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

		case "cache_miss_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CacheMissCount = value
			case float64:
				f := int(v)
				s.CacheMissCount = f
			}

		case "failure_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FailureCount = value
			case float64:
				f := int(v)
				s.FailureCount = f
			}

		case "inference_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.InferenceCount = value
			case float64:
				f := int(v)
				s.InferenceCount = f
			}

		case "missing_all_fields_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MissingAllFieldsCount = value
			case float64:
				f := int(v)
				s.MissingAllFieldsCount = f
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTrainedModelInferenceStats returns a TrainedModelInferenceStats.
func NewTrainedModelInferenceStats() *TrainedModelInferenceStats {
	r := &TrainedModelInferenceStats{}

	return r
}
