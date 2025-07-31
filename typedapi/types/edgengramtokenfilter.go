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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/edgengramside"
)

// EdgeNGramTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L97-L107
type EdgeNGramTokenFilter struct {
	// MaxGram Maximum character length of a gram. For custom token filters, defaults to
	// `2`. For the built-in edge_ngram filter, defaults to `1`.
	MaxGram *int `json:"max_gram,omitempty"`
	// MinGram Minimum character length of a gram. Defaults to `1`.
	MinGram *int `json:"min_gram,omitempty"`
	// PreserveOriginal Emits original token when set to `true`. Defaults to `false`.
	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`
	// Side Indicates whether to truncate tokens from the `front` or `back`. Defaults to
	// `front`.
	Side    *edgengramside.EdgeNGramSide `json:"side,omitempty"`
	Type    string                       `json:"type,omitempty"`
	Version *string                      `json:"version,omitempty"`
}

func (s *EdgeNGramTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "max_gram":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxGram", err)
				}
				s.MaxGram = &value
			case float64:
				f := int(v)
				s.MaxGram = &f
			}

		case "min_gram":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinGram", err)
				}
				s.MinGram = &value
			case float64:
				f := int(v)
				s.MinGram = &f
			}

		case "preserve_original":
			if err := dec.Decode(&s.PreserveOriginal); err != nil {
				return fmt.Errorf("%s | %w", "PreserveOriginal", err)
			}

		case "side":
			if err := dec.Decode(&s.Side); err != nil {
				return fmt.Errorf("%s | %w", "Side", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s EdgeNGramTokenFilter) MarshalJSON() ([]byte, error) {
	type innerEdgeNGramTokenFilter EdgeNGramTokenFilter
	tmp := innerEdgeNGramTokenFilter{
		MaxGram:          s.MaxGram,
		MinGram:          s.MinGram,
		PreserveOriginal: s.PreserveOriginal,
		Side:             s.Side,
		Type:             s.Type,
		Version:          s.Version,
	}

	tmp.Type = "edge_ngram"

	return json.Marshal(tmp)
}

// NewEdgeNGramTokenFilter returns a EdgeNGramTokenFilter.
func NewEdgeNGramTokenFilter() *EdgeNGramTokenFilter {
	r := &EdgeNGramTokenFilter{}

	return r
}
