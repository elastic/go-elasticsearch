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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DateIndexNameProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ingest/_types/Processors.ts#L730-L768
type DateIndexNameProcessor struct {
	// DateFormats An array of the expected date formats for parsing dates / timestamps in the
	// document being preprocessed.
	// Can be a java time pattern or one of the following formats: ISO8601, UNIX,
	// UNIX_MS, or TAI64N.
	DateFormats []string `json:"date_formats,omitempty"`
	// DateRounding How to round the date when formatting the date into the index name. Valid
	// values are:
	// `y` (year), `M` (month), `w` (week), `d` (day), `h` (hour), `m` (minute) and
	// `s` (second).
	// Supports template snippets.
	DateRounding string `json:"date_rounding"`
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// Field The field to get the date or timestamp from.
	Field string `json:"field"`
	// If Conditionally execute the processor.
	If *Script `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IndexNameFormat The format to be used when printing the parsed date into the index name.
	// A valid java time pattern is expected here.
	// Supports template snippets.
	IndexNameFormat *string `json:"index_name_format,omitempty"`
	// IndexNamePrefix A prefix of the index name to be prepended before the printed date.
	// Supports template snippets.
	IndexNamePrefix *string `json:"index_name_prefix,omitempty"`
	// Locale The locale to use when parsing the date from the document being preprocessed,
	// relevant when parsing month names or week days.
	Locale *string `json:"locale,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// Timezone The timezone to use when parsing the date and when date math index supports
	// resolves expressions into concrete index names.
	Timezone *string `json:"timezone,omitempty"`
}

func (s *DateIndexNameProcessor) UnmarshalJSON(data []byte) error {

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

		case "date_formats":
			if err := dec.Decode(&s.DateFormats); err != nil {
				return fmt.Errorf("%s | %w", "DateFormats", err)
			}

		case "date_rounding":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DateRounding", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DateRounding = o

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
			if err := dec.Decode(&s.If); err != nil {
				return fmt.Errorf("%s | %w", "If", err)
			}

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

		case "index_name_format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "IndexNameFormat", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexNameFormat = &o

		case "index_name_prefix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "IndexNamePrefix", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexNamePrefix = &o

		case "locale":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Locale", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Locale = &o

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
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

		case "timezone":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Timezone", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Timezone = &o

		}
	}
	return nil
}

// NewDateIndexNameProcessor returns a DateIndexNameProcessor.
func NewDateIndexNameProcessor() *DateIndexNameProcessor {
	r := &DateIndexNameProcessor{}

	return r
}

type DateIndexNameProcessorVariant interface {
	DateIndexNameProcessorCaster() *DateIndexNameProcessor
}

func (s *DateIndexNameProcessor) DateIndexNameProcessorCaster() *DateIndexNameProcessor {
	return s
}
