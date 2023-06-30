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

// PipelineSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/logstash/_types/Pipeline.ts#L28-L36
type PipelineSettings struct {
	PipelineBatchDelay    int    `json:"pipeline.batch.delay"`
	PipelineBatchSize     int    `json:"pipeline.batch.size"`
	PipelineWorkers       int    `json:"pipeline.workers"`
	QueueCheckpointWrites int    `json:"queue.checkpoint.writes"`
	QueueMaxBytesNumber   int    `json:"queue.max_bytes.number"`
	QueueMaxBytesUnits    string `json:"queue.max_bytes.units"`
	QueueType             string `json:"queue.type"`
}

func (s *PipelineSettings) UnmarshalJSON(data []byte) error {

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

		case "pipeline.batch.delay":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PipelineBatchDelay = value
			case float64:
				f := int(v)
				s.PipelineBatchDelay = f
			}

		case "pipeline.batch.size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PipelineBatchSize = value
			case float64:
				f := int(v)
				s.PipelineBatchSize = f
			}

		case "pipeline.workers":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PipelineWorkers = value
			case float64:
				f := int(v)
				s.PipelineWorkers = f
			}

		case "queue.checkpoint.writes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.QueueCheckpointWrites = value
			case float64:
				f := int(v)
				s.QueueCheckpointWrites = f
			}

		case "queue.max_bytes.number":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.QueueMaxBytesNumber = value
			case float64:
				f := int(v)
				s.QueueMaxBytesNumber = f
			}

		case "queue.max_bytes.units":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueueMaxBytesUnits = o

		case "queue.type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueueType = o

		}
	}
	return nil
}

// NewPipelineSettings returns a PipelineSettings.
func NewPipelineSettings() *PipelineSettings {
	r := &PipelineSettings{}

	return r
}
