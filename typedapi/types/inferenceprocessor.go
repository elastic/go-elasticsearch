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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// InferenceProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ingest/_types/Processors.ts#L237-L242
type InferenceProcessor struct {
	Description     *string                    `json:"description,omitempty"`
	FieldMap        map[string]json.RawMessage `json:"field_map,omitempty"`
	If              *string                    `json:"if,omitempty"`
	IgnoreFailure   *bool                      `json:"ignore_failure,omitempty"`
	InferenceConfig *InferenceConfig           `json:"inference_config,omitempty"`
	ModelId         string                     `json:"model_id"`
	OnFailure       []ProcessorContainer       `json:"on_failure,omitempty"`
	Tag             *string                    `json:"tag,omitempty"`
	TargetField     *string                    `json:"target_field,omitempty"`
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
				return err
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
				return err
			}

		case "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.If = &o

		case "ignore_failure":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreFailure = &value
			case bool:
				s.IgnoreFailure = &v
			}

		case "inference_config":
			if err := dec.Decode(&s.InferenceConfig); err != nil {
				return err
			}

		case "model_id":
			if err := dec.Decode(&s.ModelId); err != nil {
				return err
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return err
			}

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tag = &o

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewInferenceProcessor returns a InferenceProcessor.
func NewInferenceProcessor() *InferenceProcessor {
	r := &InferenceProcessor{
		FieldMap: make(map[string]json.RawMessage, 0),
	}

	return r
}
