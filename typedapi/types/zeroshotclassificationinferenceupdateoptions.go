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

// ZeroShotClassificationInferenceUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/inference.ts#L393-L402
type ZeroShotClassificationInferenceUpdateOptions struct {
	// Labels The labels to predict.
	Labels []string `json:"labels"`
	// MultiLabel Update the configured multi label option. Indicates if more than one true
	// label exists. Defaults to the configured value.
	MultiLabel *bool `json:"multi_label,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *ZeroShotClassificationInferenceUpdateOptions) UnmarshalJSON(data []byte) error {

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

		case "labels":
			if err := dec.Decode(&s.Labels); err != nil {
				return fmt.Errorf("%s | %w", "Labels", err)
			}

		case "multi_label":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MultiLabel", err)
				}
				s.MultiLabel = &value
			case bool:
				s.MultiLabel = &v
			}

		case "results_field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ResultsField", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResultsField = &o

		case "tokenization":
			if err := dec.Decode(&s.Tokenization); err != nil {
				return fmt.Errorf("%s | %w", "Tokenization", err)
			}

		}
	}
	return nil
}

// NewZeroShotClassificationInferenceUpdateOptions returns a ZeroShotClassificationInferenceUpdateOptions.
func NewZeroShotClassificationInferenceUpdateOptions() *ZeroShotClassificationInferenceUpdateOptions {
	r := &ZeroShotClassificationInferenceUpdateOptions{}

	return r
}

type ZeroShotClassificationInferenceUpdateOptionsVariant interface {
	ZeroShotClassificationInferenceUpdateOptionsCaster() *ZeroShotClassificationInferenceUpdateOptions
}

func (s *ZeroShotClassificationInferenceUpdateOptions) ZeroShotClassificationInferenceUpdateOptionsCaster() *ZeroShotClassificationInferenceUpdateOptions {
	return s
}
