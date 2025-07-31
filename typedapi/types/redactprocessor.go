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

// RedactProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Processors.ts#L1297-L1338
type RedactProcessor struct {
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// Field The field to be redacted
	Field string `json:"field"`
	// If Conditionally execute the processor.
	If *string `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If `true` and `field` does not exist or is `null`, the processor quietly
	// exits without modifying the document.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure          []ProcessorContainer `json:"on_failure,omitempty"`
	PatternDefinitions map[string]string    `json:"pattern_definitions,omitempty"`
	// Patterns A list of grok expressions to match and redact named captures with
	Patterns []string `json:"patterns"`
	// Prefix Start a redacted section with this token
	Prefix *string `json:"prefix,omitempty"`
	// SkipIfUnlicensed If `true` and the current license does not support running redact processors,
	// then the processor quietly exits without modifying the document
	SkipIfUnlicensed *bool `json:"skip_if_unlicensed,omitempty"`
	// Suffix End a redacted section with this token
	Suffix *string `json:"suffix,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TraceRedact If `true` then ingest metadata `_ingest._redact._is_redacted` is set to
	// `true` if the document has been redacted
	TraceRedact *bool `json:"trace_redact,omitempty"`
}

func (s *RedactProcessor) UnmarshalJSON(data []byte) error {

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

		case "prefix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Prefix", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Prefix = &o

		case "skip_if_unlicensed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SkipIfUnlicensed", err)
				}
				s.SkipIfUnlicensed = &value
			case bool:
				s.SkipIfUnlicensed = &v
			}

		case "suffix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Suffix", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Suffix = &o

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

		case "trace_redact":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TraceRedact", err)
				}
				s.TraceRedact = &value
			case bool:
				s.TraceRedact = &v
			}

		}
	}
	return nil
}

// NewRedactProcessor returns a RedactProcessor.
func NewRedactProcessor() *RedactProcessor {
	r := &RedactProcessor{
		PatternDefinitions: make(map[string]string),
	}

	return r
}
