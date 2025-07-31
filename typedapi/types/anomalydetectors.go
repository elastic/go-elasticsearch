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

// AnomalyDetectors type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/info/types.ts#L46-L52
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
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "CategorizationAnalyzer", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		categorizationanalyzer_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "CategorizationAnalyzer", err)
				}

				switch t {

				case "char_filter", "filter", "tokenizer":
					o := NewCategorizationAnalyzerDefinition()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "CategorizationAnalyzer", err)
					}
					s.CategorizationAnalyzer = o
					break categorizationanalyzer_field

				}
			}
			if s.CategorizationAnalyzer == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.CategorizationAnalyzer); err != nil {
					return fmt.Errorf("%s | %w", "CategorizationAnalyzer", err)
				}
			}

		case "categorization_examples_limit":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CategorizationExamplesLimit", err)
				}
				s.CategorizationExamplesLimit = value
			case float64:
				f := int(v)
				s.CategorizationExamplesLimit = f
			}

		case "daily_model_snapshot_retention_after_days":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DailyModelSnapshotRetentionAfterDays", err)
				}
				s.DailyModelSnapshotRetentionAfterDays = value
			case float64:
				f := int(v)
				s.DailyModelSnapshotRetentionAfterDays = f
			}

		case "model_memory_limit":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ModelMemoryLimit", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelMemoryLimit = o

		case "model_snapshot_retention_days":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ModelSnapshotRetentionDays", err)
				}
				s.ModelSnapshotRetentionDays = value
			case float64:
				f := int(v)
				s.ModelSnapshotRetentionDays = f
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
