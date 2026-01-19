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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ChunkRescorerChunkingSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/_types/mapping/ChunkingSettings.ts#L38-L52
type ChunkRescorerChunkingSettings struct {
	// MaxChunkSize The maximum size of a chunk in words.
	// This value cannot be lower than `20` (for `sentence` strategy) or `10` (for
	// `word` strategy).
	// This value should not exceed the window size for the associated model.
	MaxChunkSize int `json:"max_chunk_size"`
	// Overlap The number of overlapping words for chunks.
	// It is applicable only to a `word` chunking strategy.
	// This value cannot be higher than half the `max_chunk_size` value.
	Overlap *int `json:"overlap,omitempty"`
	// SentenceOverlap The number of overlapping sentences for chunks.
	// It is applicable only for a `sentence` chunking strategy.
	// It can be either `1` or `0`.
	SentenceOverlap *int `json:"sentence_overlap,omitempty"`
	// SeparatorGroup Only applicable to the `recursive` strategy and required when using it.
	//
	// Sets a predefined list of separators in the saved chunking settings based on
	// the selected text type.
	// Values can be `markdown` or `plaintext`.
	//
	// Using this parameter is an alternative to manually specifying a custom
	// `separators` list.
	SeparatorGroup *string `json:"separator_group,omitempty"`
	// Separators Only applicable to the `recursive` strategy and required when using it.
	//
	// A list of strings used as possible split points when chunking text.
	//
	// Each string can be a plain string or a regular expression (regex) pattern.
	// The system tries each separator in order to split the text, starting from the
	// first item in the list.
	//
	// After splitting, it attempts to recombine smaller pieces into larger chunks
	// that stay within
	// the `max_chunk_size` limit, to reduce the total number of chunks generated.
	Separators []string `json:"separators,omitempty"`
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

func (s *ChunkRescorerChunkingSettings) UnmarshalJSON(data []byte) error {

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
				s.MaxChunkSize = value
			case float64:
				f := int(v)
				s.MaxChunkSize = f
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
			s.SeparatorGroup = &o

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

// NewChunkRescorerChunkingSettings returns a ChunkRescorerChunkingSettings.
func NewChunkRescorerChunkingSettings() *ChunkRescorerChunkingSettings {
	r := &ChunkRescorerChunkingSettings{}

	return r
}

type ChunkRescorerChunkingSettingsVariant interface {
	ChunkRescorerChunkingSettingsCaster() *ChunkRescorerChunkingSettings
}

func (s *ChunkRescorerChunkingSettings) ChunkRescorerChunkingSettingsCaster() *ChunkRescorerChunkingSettings {
	return s
}
