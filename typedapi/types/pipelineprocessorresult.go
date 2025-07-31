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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/pipelinesimulationstatusoptions"
)

// PipelineProcessorResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Simulation.ts#L59-L67
type PipelineProcessorResult struct {
	Description   *string                                                          `json:"description,omitempty"`
	Doc           *DocumentSimulation                                              `json:"doc,omitempty"`
	Error         *ErrorCause                                                      `json:"error,omitempty"`
	IgnoredError  *ErrorCause                                                      `json:"ignored_error,omitempty"`
	ProcessorType *string                                                          `json:"processor_type,omitempty"`
	Status        *pipelinesimulationstatusoptions.PipelineSimulationStatusOptions `json:"status,omitempty"`
	Tag           *string                                                          `json:"tag,omitempty"`
}

func (s *PipelineProcessorResult) UnmarshalJSON(data []byte) error {

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

		case "doc":
			if err := dec.Decode(&s.Doc); err != nil {
				return fmt.Errorf("%s | %w", "Doc", err)
			}

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}

		case "ignored_error":
			if err := dec.Decode(&s.IgnoredError); err != nil {
				return fmt.Errorf("%s | %w", "IgnoredError", err)
			}

		case "processor_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ProcessorType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ProcessorType = &o

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
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

		}
	}
	return nil
}

// NewPipelineProcessorResult returns a PipelineProcessorResult.
func NewPipelineProcessorResult() *PipelineProcessorResult {
	r := &PipelineProcessorResult{}

	return r
}
