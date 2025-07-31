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

// GrokProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Processors.ts#L950-L981
type GrokProcessor struct {
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// EcsCompatibility Must be disabled or v1. If v1, the processor uses patterns with Elastic
	// Common Schema (ECS) field names.
	EcsCompatibility *string `json:"ecs_compatibility,omitempty"`
	// Field The field to use for grok expression parsing.
	Field string `json:"field"`
	// If Conditionally execute the processor.
	If *string `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If `true` and `field` does not exist or is `null`, the processor quietly
	// exits without modifying the document.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// PatternDefinitions A map of pattern-name and pattern tuples defining custom patterns to be used
	// by the current processor.
	// Patterns matching existing names will override the pre-existing definition.
	PatternDefinitions map[string]string `json:"pattern_definitions,omitempty"`
	// Patterns An ordered list of grok expression to match and extract named captures with.
	// Returns on the first expression in the list that matches.
	Patterns []string `json:"patterns"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TraceMatch When `true`, `_ingest._grok_match_index` will be inserted into your matched
	// document’s metadata with the index into the pattern found in `patterns` that
	// matched.
	TraceMatch *bool `json:"trace_match,omitempty"`
}

func (s *GrokProcessor) UnmarshalJSON(data []byte) error {

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

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "ecs_compatibility":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "EcsCompatibility", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.EcsCompatibility = &o

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "If", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.If = &o

		case "ignore_failure":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreFailure", err)
				}
				s.IgnoreFailure = &value
			case bool:
				s.IgnoreFailure = &v
			}

		case "ignore_missing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreMissing", err)
				}
				s.IgnoreMissing = &value
			case bool:
				s.IgnoreMissing = &v
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
			}

		case "pattern_definitions":
			if s.PatternDefinitions == nil {
				s.PatternDefinitions = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.PatternDefinitions); err != nil {
				return fmt.Errorf("%s | %w", "PatternDefinitions", err)
			}

		case "patterns":
			if err := dec.Decode(&s.Patterns); err != nil {
				return fmt.Errorf("%s | %w", "Patterns", err)
			}

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tag", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tag = &o

		case "trace_match":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TraceMatch", err)
				}
				s.TraceMatch = &value
			case bool:
				s.TraceMatch = &v
			}

		}
	}
	return nil
}

// NewGrokProcessor returns a GrokProcessor.
func NewGrokProcessor() *GrokProcessor {
	r := &GrokProcessor{
		PatternDefinitions: make(map[string]string),
	}

	return r
}
