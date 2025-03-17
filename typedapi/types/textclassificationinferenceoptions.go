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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TextClassificationInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/ml/_types/inference.ts#L173-L183
type TextClassificationInferenceOptions struct {
	// ClassificationLabels Classification labels to apply other than the stored labels. Must have the
	// same deminsions as the default configured labels
	ClassificationLabels []string `json:"classification_labels,omitempty"`
	// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

func (s *TextClassificationInferenceOptions) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "ClassificationLabels", err)
			}

		case "num_top_classes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumTopClasses", err)
				}
				s.NumTopClasses = &value
			case float64:
				f := int(v)
				s.NumTopClasses = &f
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

// NewTextClassificationInferenceOptions returns a TextClassificationInferenceOptions.
func NewTextClassificationInferenceOptions() *TextClassificationInferenceOptions {
	r := &TextClassificationInferenceOptions{}

	return r
}

// true

type TextClassificationInferenceOptionsVariant interface {
	TextClassificationInferenceOptionsCaster() *TextClassificationInferenceOptions
}

func (s *TextClassificationInferenceOptions) TextClassificationInferenceOptionsCaster() *TextClassificationInferenceOptions {
	return s
}
