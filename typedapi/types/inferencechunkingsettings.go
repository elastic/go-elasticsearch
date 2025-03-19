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
// https://github.com/elastic/elasticsearch-specification/tree/3ea9ce260df22d3244bff5bace485dd97ff4046d

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
// https://github.com/elastic/elasticsearch-specification/blob/3ea9ce260df22d3244bff5bace485dd97ff4046d/specification/inference/_types/Services.ts#L60-L89
type InferenceChunkingSettings struct {
	// ChunkingSettings Chunking configuration object
	ChunkingSettings *InferenceChunkingSettings `json:"chunking_settings,omitempty"`
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
	// Service The service type
	Service string `json:"service"`
	// ServiceSettings Settings specific to the service
	ServiceSettings json.RawMessage `json:"service_settings"`
	// Strategy The chunking strategy: `sentence` or `word`.
	Strategy *string `json:"strategy,omitempty"`
	// TaskSettings Task settings specific to the service and task type
	TaskSettings json.RawMessage `json:"task_settings,omitempty"`
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

		case "chunking_settings":
			if err := dec.Decode(&s.ChunkingSettings); err != nil {
				return fmt.Errorf("%s | %w", "ChunkingSettings", err)
			}

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

		case "service":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Service", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Service = o

		case "service_settings":
			if err := dec.Decode(&s.ServiceSettings); err != nil {
				return fmt.Errorf("%s | %w", "ServiceSettings", err)
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

		case "task_settings":
			if err := dec.Decode(&s.TaskSettings); err != nil {
				return fmt.Errorf("%s | %w", "TaskSettings", err)
			}

		}
	}
	return nil
}

// NewInferenceChunkingSettings returns a InferenceChunkingSettings.
func NewInferenceChunkingSettings() *InferenceChunkingSettings {
	r := &InferenceChunkingSettings{}

	return r
}

// true

type InferenceChunkingSettingsVariant interface {
	InferenceChunkingSettingsCaster() *InferenceChunkingSettings
}

func (s *InferenceChunkingSettings) InferenceChunkingSettingsCaster() *InferenceChunkingSettings {
	return s
}
