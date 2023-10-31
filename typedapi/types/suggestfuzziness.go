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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// SuggestFuzziness type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/search/_types/suggester.ts#L193-L221
type SuggestFuzziness struct {
	// Fuzziness The fuzziness factor.
	Fuzziness Fuzziness `json:"fuzziness,omitempty"`
	// MinLength Minimum length of the input before fuzzy suggestions are returned.
	MinLength *int `json:"min_length,omitempty"`
	// PrefixLength Minimum length of the input, which is not checked for fuzzy alternatives.
	PrefixLength *int `json:"prefix_length,omitempty"`
	// Transpositions If set to `true`, transpositions are counted as one change instead of two.
	Transpositions *bool `json:"transpositions,omitempty"`
	// UnicodeAware If `true`, all measurements (like fuzzy edit distance, transpositions, and
	// lengths) are measured in Unicode code points instead of in bytes.
	// This is slightly slower than raw bytes.
	UnicodeAware *bool `json:"unicode_aware,omitempty"`
}

func (s *SuggestFuzziness) UnmarshalJSON(data []byte) error {

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

		case "fuzziness":
			if err := dec.Decode(&s.Fuzziness); err != nil {
				return err
			}

		case "min_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinLength = &value
			case float64:
				f := int(v)
				s.MinLength = &f
			}

		case "prefix_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PrefixLength = &value
			case float64:
				f := int(v)
				s.PrefixLength = &f
			}

		case "transpositions":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Transpositions = &value
			case bool:
				s.Transpositions = &v
			}

		case "unicode_aware":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.UnicodeAware = &value
			case bool:
				s.UnicodeAware = &v
			}

		}
	}
	return nil
}

// NewSuggestFuzziness returns a SuggestFuzziness.
func NewSuggestFuzziness() *SuggestFuzziness {
	r := &SuggestFuzziness{}

	return r
}
