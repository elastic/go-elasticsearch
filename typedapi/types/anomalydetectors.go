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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// AnomalyDetectors type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/info/types.ts#L44-L50
type AnomalyDetectors struct {
	CategorizationAnalyzer               CategorizationAnalyzer `json:"categorization_analyzer"`
	CategorizationExamplesLimit          int                    `json:"categorization_examples_limit"`
	DailyModelSnapshotRetentionAfterDays int                    `json:"daily_model_snapshot_retention_after_days"`
	ModelMemoryLimit                     string                 `json:"model_memory_limit"`
	ModelSnapshotRetentionDays           int                    `json:"model_snapshot_retention_days"`
}

func (s *AnomalyDetectors) UnmarshalJSON(data []byte) error {
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

		case "categorization_analyzer":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := NewCategorizationAnalyzerDefinition()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.CategorizationAnalyzer = *o

			default:
				if err := localDec.Decode(&s.CategorizationAnalyzer); err != nil {
					return err
				}
			}

		case "categorization_examples_limit":
			if err := dec.Decode(&s.CategorizationExamplesLimit); err != nil {
				return err
			}

		case "daily_model_snapshot_retention_after_days":
			if err := dec.Decode(&s.DailyModelSnapshotRetentionAfterDays); err != nil {
				return err
			}

		case "model_memory_limit":
			if err := dec.Decode(&s.ModelMemoryLimit); err != nil {
				return err
			}

		case "model_snapshot_retention_days":
			if err := dec.Decode(&s.ModelSnapshotRetentionDays); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAnomalyDetectors returns a AnomalyDetectors.
func NewAnomalyDetectors() *AnomalyDetectors {
	r := &AnomalyDetectors{}

	return r
}
