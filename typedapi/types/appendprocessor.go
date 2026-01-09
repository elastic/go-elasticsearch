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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// AppendProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/ingest/_types/Processors.ts#L334-L366
type AppendProcessor struct {
	// AllowDuplicates If `false`, the processor does not append values already present in the
	// field.
	AllowDuplicates *bool `json:"allow_duplicates,omitempty"`
	// CopyFrom The origin field which will be appended to `field`, cannot set `value`
	// simultaneously.
	CopyFrom *string `json:"copy_from,omitempty"`
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// Field The field to be appended to.
	// Supports template snippets.
	Field string `json:"field"`
	// If Conditionally execute the processor.
	If *Script `json:"if,omitempty"`
	// IgnoreEmptyValues If `true`, the processor will skip empty values from the source (e.g. empty
	// strings, and null values),
	// rather than appending them to the field.
	IgnoreEmptyValues *bool `json:"ignore_empty_values,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// MediaType The media type for encoding `value`.
	// Applies only when value is a template snippet.
	// Must be one of `application/json`, `text/plain`, or
	// `application/x-www-form-urlencoded`.
	MediaType *string `json:"media_type,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// Value The value to be appended. Supports template snippets. May specify only one of
	// `value` or `copy_from`.
	Value []json.RawMessage `json:"value,omitempty"`
}

func (s *AppendProcessor) UnmarshalJSON(data []byte) error {

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

		case "allow_duplicates":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowDuplicates", err)
				}
				s.AllowDuplicates = &value
			case bool:
				s.AllowDuplicates = &v
			}

		case "copy_from":
			if err := dec.Decode(&s.CopyFrom); err != nil {
				return fmt.Errorf("%s | %w", "CopyFrom", err)
			}

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

		case "ignore_empty_values":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreEmptyValues", err)
				}
				s.IgnoreEmptyValues = &value
			case bool:
				s.IgnoreEmptyValues = &v
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

		case "media_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MediaType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MediaType = &o

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

		case "value":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(json.RawMessage)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Value", err)
				}

				s.Value = append(s.Value, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Value); err != nil {
					return fmt.Errorf("%s | %w", "Value", err)
				}
			}

		}
	}
	return nil
}

// NewAppendProcessor returns a AppendProcessor.
func NewAppendProcessor() *AppendProcessor {
	r := &AppendProcessor{}

	return r
}

type AppendProcessorVariant interface {
	AppendProcessorCaster() *AppendProcessor
}

func (s *AppendProcessor) AppendProcessorCaster() *AppendProcessor {
	return s
}
