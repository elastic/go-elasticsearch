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
	"encoding/json"
	"fmt"
)

// Preprocessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/put_trained_model/types.ts#L31-L36
type Preprocessor struct {
	AdditionalPreprocessorProperty map[string]json.RawMessage      `json:"-"`
	FrequencyEncoding              *FrequencyEncodingPreprocessor  `json:"frequency_encoding,omitempty"`
	OneHotEncoding                 *OneHotEncodingPreprocessor     `json:"one_hot_encoding,omitempty"`
	TargetMeanEncoding             *TargetMeanEncodingPreprocessor `json:"target_mean_encoding,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s Preprocessor) MarshalJSON() ([]byte, error) {
	type opt Preprocessor
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalPreprocessorProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalPreprocessorProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewPreprocessor returns a Preprocessor.
func NewPreprocessor() *Preprocessor {
	r := &Preprocessor{
		AdditionalPreprocessorProperty: make(map[string]json.RawMessage),
	}

	return r
}
