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

// InferenceProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Processors.ts#L1027-L1059
type InferenceProcessor struct {
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// FieldMap Maps the document field names to the known field names of the model.
	// This mapping takes precedence over any default mappings provided in the model
	// configuration.
	FieldMap map[string]json.RawMessage `json:"field_map,omitempty"`
	// If Conditionally execute the processor.
	If *string `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If true and any of the input fields defined in input_ouput are missing
	// then those missing fields are quietly ignored, otherwise a missing field
	// causes a failure.
	// Only applies when using input_output configurations to explicitly list the
	// input fields.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// InferenceConfig Contains the inference type and its options.
	InferenceConfig *InferenceConfig `json:"inference_config,omitempty"`
	// InputOutput Input fields for inference and output (destination) fields for the inference
	// results.
	// This option is incompatible with the target_field and field_map options.
	InputOutput []InputConfig `json:"input_output,omitempty"`
	// ModelId The ID or alias for the trained model, or the ID of the deployment.
	ModelId string `json:"model_id"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TargetField Field added to incoming documents to contain results objects.
	TargetField *string `json:"target_field,omitempty"`
}

func (s *InferenceProcessor) UnmarshalJSON(data []byte) error {

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

		case "field_map":
			if s.FieldMap == nil {
				s.FieldMap = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.FieldMap); err != nil {
				return fmt.Errorf("%s | %w", "FieldMap", err)
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

		case "inference_config":
			if err := dec.Decode(&s.InferenceConfig); err != nil {
				return fmt.Errorf("%s | %w", "InferenceConfig", err)
			}

		case "input_output":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewInputConfig()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "InputOutput", err)
				}

				s.InputOutput = append(s.InputOutput, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.InputOutput); err != nil {
					return fmt.Errorf("%s | %w", "InputOutput", err)
				}
			}

		case "model_id":
			if err := dec.Decode(&s.ModelId); err != nil {
				return fmt.Errorf("%s | %w", "ModelId", err)
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

// NewInferenceProcessor returns a InferenceProcessor.
func NewInferenceProcessor() *InferenceProcessor {
	r := &InferenceProcessor{
		FieldMap: make(map[string]json.RawMessage),
	}

	return r
}
