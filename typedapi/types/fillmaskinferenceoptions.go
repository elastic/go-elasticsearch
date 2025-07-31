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

// FillMaskInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/inference.ts#L284-L299
type FillMaskInferenceOptions struct {
	// MaskToken The string/token which will be removed from incoming documents and replaced
	// with the inference prediction(s).
	// In a response, this field contains the mask token for the specified
	// model/tokenizer. Each model and tokenizer
	// has a predefined mask token which cannot be changed. Thus, it is recommended
	// not to set this value in requests.
	// However, if this field is present in a request, its value must match the
	// predefined value for that model/tokenizer,
	// otherwise the request will fail.
	MaskToken *string `json:"mask_token,omitempty"`
	// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
	Vocabulary   Vocabulary                   `json:"vocabulary"`
}

func (s *FillMaskInferenceOptions) UnmarshalJSON(data []byte) error {

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

		case "mask_token":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaskToken", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaskToken = &o

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

		case "vocabulary":
			if err := dec.Decode(&s.Vocabulary); err != nil {
				return fmt.Errorf("%s | %w", "Vocabulary", err)
			}

		}
	}
	return nil
}

// NewFillMaskInferenceOptions returns a FillMaskInferenceOptions.
func NewFillMaskInferenceOptions() *FillMaskInferenceOptions {
	r := &FillMaskInferenceOptions{}

	return r
}
