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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jsonprocessorconflictstrategy"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// JsonProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ingest/_types/Processors.ts#L271-L277
type JsonProcessor struct {
	AddToRoot                 *bool                                                        `json:"add_to_root,omitempty"`
	AddToRootConflictStrategy *jsonprocessorconflictstrategy.JsonProcessorConflictStrategy `json:"add_to_root_conflict_strategy,omitempty"`
	AllowDuplicateKeys        *bool                                                        `json:"allow_duplicate_keys,omitempty"`
	Description               *string                                                      `json:"description,omitempty"`
	Field                     string                                                       `json:"field"`
	If                        *string                                                      `json:"if,omitempty"`
	IgnoreFailure             *bool                                                        `json:"ignore_failure,omitempty"`
	OnFailure                 []ProcessorContainer                                         `json:"on_failure,omitempty"`
	Tag                       *string                                                      `json:"tag,omitempty"`
	TargetField               *string                                                      `json:"target_field,omitempty"`
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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AddToRoot = &value
			case bool:
				s.AddToRoot = &v
			}

		case "add_to_root_conflict_strategy":
			if err := dec.Decode(&s.AddToRootConflictStrategy); err != nil {
				return err
			}

		case "allow_duplicate_keys":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowDuplicateKeys = &value
			case bool:
				s.AllowDuplicateKeys = &v
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Description = &o

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
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

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return err
			}

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Tag = &o

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return err
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
