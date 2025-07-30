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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// InferenceChunkingSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e585438d116b00ff34643179e6286e402c0bcaaf/specification/inference/_types/Services.ts#L288-L344
type InferenceChunkingSettings struct {
	// MaxChunkSize The maximum size of a chunk in words.
	// This value cannot be higher than `300` or lower than `20` (for `sentence`
	// strategy) or `10` (for `word` strategy).
	MaxChunkSize *int `json:"max_chunk_size,omitempty"`
	// Overlap The number of overlapping words for chunks.
	// It is applicable only to a `word` chunking strategy.
	// This value cannot be higher than half the `max_chunk_size` value.
	Overlap *int `json:"overlap,omitempty"`
	// SentenceOverlap The number of overlapping sentences for chunks.
	// It is applicable only for a `sentence` chunking strategy.
	// It can be either `1` or `0`.
	SentenceOverlap *int `json:"sentence_overlap,omitempty"`
	// SeparatorGroup This parameter is only applicable when using the `recursive` chunking
	// strategy.
	//
	// Sets a predefined list of separators in the saved chunking settings based on
	// the selected text type.
	// Values can be `markdown` or `plaintext`.
	//
	// Using this parameter is an alternative to manually specifying a custom
	// `separators` list.
	SeparatorGroup string `json:"separator_group"`
	// Separators A list of strings used as possible split points when chunking text with the
	// `recursive` strategy.
	//
	// Each string can be a plain string or a regular expression (regex) pattern.
	// The system tries each separator in order to split the text, starting from the
	// first item in the list.
	//
	// After splitting, it attempts to recombine smaller pieces into larger chunks
	// that stay within
	// the `max_chunk_size` limit, to reduce the total number of chunks generated.
	Separators []string `json:"separators"`
	// Strategy The chunking strategy: `sentence`, `word`, `none` or `recursive`.
	//
	//   - If `strategy` is set to `recursive`, you must also specify:
	//
	// - `max_chunk_size`
	// - either `separators` or`separator_group`
	//
	// Learn more about different chunking strategies in the linked documentation.
	Strategy *string `json:"strategy,omitempty"`
}

func (s *InferenceChunkingSettings) UnmarshalJSON(data []byte) error {

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

		case "max_chunk_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxChunkSize", err)
				}
				s.MaxChunkSize = &value
			case float64:
				f := int(v)
				s.MaxChunkSize = &f
			}

		case "overlap":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Overlap", err)
				}
				s.Overlap = &value
			case float64:
				f := int(v)
				s.Overlap = &f
			}

		case "sentence_overlap":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SentenceOverlap", err)
				}
				s.SentenceOverlap = &value
			case float64:
				f := int(v)
				s.SentenceOverlap = &f
			}

		case "separator_group":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SeparatorGroup", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SeparatorGroup = o

		case "separators":
			if err := dec.Decode(&s.Separators); err != nil {
				return fmt.Errorf("%s | %w", "Separators", err)
			}

		case "strategy":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Strategy", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Strategy = &o

		}
	}
	return nil
}

// NewInferenceChunkingSettings returns a InferenceChunkingSettings.
func NewInferenceChunkingSettings() *InferenceChunkingSettings {
	r := &InferenceChunkingSettings{}

	return r
}

type InferenceChunkingSettingsVariant interface {
	InferenceChunkingSettingsCaster() *InferenceChunkingSettings
}

func (s *InferenceChunkingSettings) InferenceChunkingSettingsCaster() *InferenceChunkingSettings {
	return s
}
