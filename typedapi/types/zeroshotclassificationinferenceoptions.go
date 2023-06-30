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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ZeroShotClassificationInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/inference.ts#L191-L212
type ZeroShotClassificationInferenceOptions struct {
	// ClassificationLabels The zero shot classification labels indicating entailment, neutral, and
	// contradiction
	// Must contain exactly and only entailment, neutral, and contradiction
	ClassificationLabels []string `json:"classification_labels"`
	// HypothesisTemplate Hypothesis template used when tokenizing labels for prediction
	HypothesisTemplate *string `json:"hypothesis_template,omitempty"`
	// Labels The labels to predict.
	Labels []string `json:"labels,omitempty"`
	// MultiLabel Indicates if more than one true label exists.
	MultiLabel *bool `json:"multi_label,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

func (s *ZeroShotClassificationInferenceOptions) UnmarshalJSON(data []byte) error {

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

		case "classification_labels":
			if err := dec.Decode(&s.ClassificationLabels); err != nil {
				return err
			}

		case "hypothesis_template":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.HypothesisTemplate = &o

		case "labels":
			if err := dec.Decode(&s.Labels); err != nil {
				return err
			}

		case "multi_label":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.MultiLabel = &value
			case bool:
				s.MultiLabel = &v
			}

		case "results_field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResultsField = &o

		case "tokenization":
			if err := dec.Decode(&s.Tokenization); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewZeroShotClassificationInferenceOptions returns a ZeroShotClassificationInferenceOptions.
func NewZeroShotClassificationInferenceOptions() *ZeroShotClassificationInferenceOptions {
	r := &ZeroShotClassificationInferenceOptions{}

	return r
}
