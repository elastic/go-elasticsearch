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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jsonprocessorconflictstrategy"
)

// JsonProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ingest/_types/Processors.ts#L1139-L1168
type JsonProcessor struct {
	// AddToRoot Flag that forces the parsed JSON to be added at the top level of the
	// document.
	// `target_field` must not be set when this option is chosen.
	AddToRoot *bool `json:"add_to_root,omitempty"`
	// AddToRootConflictStrategy When set to `replace`, root fields that conflict with fields from the parsed
	// JSON will be overridden.
	// When set to `merge`, conflicting fields will be merged.
	// Only applicable `if add_to_root` is set to true.
	AddToRootConflictStrategy *jsonprocessorconflictstrategy.JsonProcessorConflictStrategy `json:"add_to_root_conflict_strategy,omitempty"`
	// AllowDuplicateKeys When set to `true`, the JSON parser will not fail if the JSON contains
	// duplicate keys.
	// Instead, the last encountered value for any duplicate key wins.
	AllowDuplicateKeys *bool `json:"allow_duplicate_keys,omitempty"`
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// Field The field to be parsed.
	Field string `json:"field"`
	// If Conditionally execute the processor.
	If *Script `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TargetField The field that the converted structured object will be written into.
	// Any existing content in this field will be overwritten.
	TargetField *string `json:"target_field,omitempty"`
}

func (s *JsonProcessor) UnmarshalJSON(data []byte) error {

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

		case "add_to_root":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AddToRoot", err)
				}
				s.AddToRoot = &value
			case bool:
				s.AddToRoot = &v
			}

		case "add_to_root_conflict_strategy":
			if err := dec.Decode(&s.AddToRootConflictStrategy); err != nil {
				return fmt.Errorf("%s | %w", "AddToRootConflictStrategy", err)
			}

		case "allow_duplicate_keys":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowDuplicateKeys", err)
				}
				s.AllowDuplicateKeys = &value
			case bool:
				s.AllowDuplicateKeys = &v
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

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return fmt.Errorf("%s | %w", "TargetField", err)
			}

		}
	}
	return nil
}

// NewJsonProcessor returns a JsonProcessor.
func NewJsonProcessor() *JsonProcessor {
	r := &JsonProcessor{}

	return r
}

type JsonProcessorVariant interface {
	JsonProcessorCaster() *JsonProcessor
}

func (s *JsonProcessor) JsonProcessorCaster() *JsonProcessor {
	return s
}
