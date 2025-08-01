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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MemMlStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/get_memory_stats/types.ts#L90-L111
type MemMlStats struct {
	// AnomalyDetectors Amount of native memory set aside for anomaly detection jobs.
	AnomalyDetectors ByteSize `json:"anomaly_detectors,omitempty"`
	// AnomalyDetectorsInBytes Amount of native memory, in bytes, set aside for anomaly detection jobs.
	AnomalyDetectorsInBytes int `json:"anomaly_detectors_in_bytes"`
	// DataFrameAnalytics Amount of native memory set aside for data frame analytics jobs.
	DataFrameAnalytics ByteSize `json:"data_frame_analytics,omitempty"`
	// DataFrameAnalyticsInBytes Amount of native memory, in bytes, set aside for data frame analytics jobs.
	DataFrameAnalyticsInBytes int `json:"data_frame_analytics_in_bytes"`
	// Max Maximum amount of native memory (separate to the JVM heap) that may be used
	// by machine learning native processes.
	Max ByteSize `json:"max,omitempty"`
	// MaxInBytes Maximum amount of native memory (separate to the JVM heap), in bytes, that
	// may be used by machine learning native processes.
	MaxInBytes int `json:"max_in_bytes"`
	// NativeCodeOverhead Amount of native memory set aside for loading machine learning native code
	// shared libraries.
	NativeCodeOverhead ByteSize `json:"native_code_overhead,omitempty"`
	// NativeCodeOverheadInBytes Amount of native memory, in bytes, set aside for loading machine learning
	// native code shared libraries.
	NativeCodeOverheadInBytes int `json:"native_code_overhead_in_bytes"`
	// NativeInference Amount of native memory set aside for trained models that have a PyTorch
	// model_type.
	NativeInference ByteSize `json:"native_inference,omitempty"`
	// NativeInferenceInBytes Amount of native memory, in bytes, set aside for trained models that have a
	// PyTorch model_type.
	NativeInferenceInBytes int `json:"native_inference_in_bytes"`
}

func (s *MemMlStats) UnmarshalJSON(data []byte) error {

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

		case "anomaly_detectors":
			if err := dec.Decode(&s.AnomalyDetectors); err != nil {
				return fmt.Errorf("%s | %w", "AnomalyDetectors", err)
			}

		case "anomaly_detectors_in_bytes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AnomalyDetectorsInBytes", err)
				}
				s.AnomalyDetectorsInBytes = value
			case float64:
				f := int(v)
				s.AnomalyDetectorsInBytes = f
			}

		case "data_frame_analytics":
			if err := dec.Decode(&s.DataFrameAnalytics); err != nil {
				return fmt.Errorf("%s | %w", "DataFrameAnalytics", err)
			}

		case "data_frame_analytics_in_bytes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataFrameAnalyticsInBytes", err)
				}
				s.DataFrameAnalyticsInBytes = value
			case float64:
				f := int(v)
				s.DataFrameAnalyticsInBytes = f
			}

		case "max":
			if err := dec.Decode(&s.Max); err != nil {
				return fmt.Errorf("%s | %w", "Max", err)
			}

		case "max_in_bytes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxInBytes", err)
				}
				s.MaxInBytes = value
			case float64:
				f := int(v)
				s.MaxInBytes = f
			}

		case "native_code_overhead":
			if err := dec.Decode(&s.NativeCodeOverhead); err != nil {
				return fmt.Errorf("%s | %w", "NativeCodeOverhead", err)
			}

		case "native_code_overhead_in_bytes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NativeCodeOverheadInBytes", err)
				}
				s.NativeCodeOverheadInBytes = value
			case float64:
				f := int(v)
				s.NativeCodeOverheadInBytes = f
			}

		case "native_inference":
			if err := dec.Decode(&s.NativeInference); err != nil {
				return fmt.Errorf("%s | %w", "NativeInference", err)
			}

		case "native_inference_in_bytes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NativeInferenceInBytes", err)
				}
				s.NativeInferenceInBytes = value
			case float64:
				f := int(v)
				s.NativeInferenceInBytes = f
			}

		}
	}
	return nil
}

// NewMemMlStats returns a MemMlStats.
func NewMemMlStats() *MemMlStats {
	r := &MemMlStats{}

	return r
}
