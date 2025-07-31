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

// MlInferenceTrainedModelsCount type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/xpack/usage/types.ts#L254-L263
type MlInferenceTrainedModelsCount struct {
	Classification *int64 `json:"classification,omitempty"`
	Ner            *int64 `json:"ner,omitempty"`
	Other          int64  `json:"other"`
	PassThrough    *int64 `json:"pass_through,omitempty"`
	Prepackaged    int64  `json:"prepackaged"`
	Regression     *int64 `json:"regression,omitempty"`
	TextEmbedding  *int64 `json:"text_embedding,omitempty"`
	Total          int64  `json:"total"`
}

func (s *MlInferenceTrainedModelsCount) UnmarshalJSON(data []byte) error {

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

		case "classification":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Classification", err)
				}
				s.Classification = &value
			case float64:
				f := int64(v)
				s.Classification = &f
			}

		case "ner":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Ner", err)
				}
				s.Ner = &value
			case float64:
				f := int64(v)
				s.Ner = &f
			}

		case "other":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Other", err)
				}
				s.Other = value
			case float64:
				f := int64(v)
				s.Other = f
			}

		case "pass_through":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PassThrough", err)
				}
				s.PassThrough = &value
			case float64:
				f := int64(v)
				s.PassThrough = &f
			}

		case "prepackaged":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Prepackaged", err)
				}
				s.Prepackaged = value
			case float64:
				f := int64(v)
				s.Prepackaged = f
			}

		case "regression":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Regression", err)
				}
				s.Regression = &value
			case float64:
				f := int64(v)
				s.Regression = &f
			}

		case "text_embedding":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TextEmbedding", err)
				}
				s.TextEmbedding = &value
			case float64:
				f := int64(v)
				s.TextEmbedding = &f
			}

		case "total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int64(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewMlInferenceTrainedModelsCount returns a MlInferenceTrainedModelsCount.
func NewMlInferenceTrainedModelsCount() *MlInferenceTrainedModelsCount {
	r := &MlInferenceTrainedModelsCount{}

	return r
}
