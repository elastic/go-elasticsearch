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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

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
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/logstash/_types/Pipeline.ts#L28-L59
type PipelineSettings struct {
	// PipelineBatchDelay When creating pipeline event batches, how long in milliseconds to wait for
	// each event before dispatching an undersized batch to pipeline workers.
	PipelineBatchDelay int `json:"pipeline.batch.delay"`
	// PipelineBatchSize The maximum number of events an individual worker thread will collect from
	// inputs before attempting to execute its filters and outputs.
	PipelineBatchSize int `json:"pipeline.batch.size"`
	// PipelineWorkers The number of workers that will, in parallel, execute the filter and output
	// stages of the pipeline.
	PipelineWorkers int `json:"pipeline.workers"`
	// QueueCheckpointWrites The maximum number of written events before forcing a checkpoint when
	// persistent queues are enabled (`queue.type: persisted`).
	QueueCheckpointWrites int `json:"queue.checkpoint.writes"`
	// QueueMaxBytesNumber The total capacity of the queue (`queue.type: persisted`) in number of bytes.
	QueueMaxBytesNumber int `json:"queue.max_bytes.number"`
	// QueueMaxBytesUnits The total capacity of the queue (`queue.type: persisted`) in terms of units
	// of bytes.
	QueueMaxBytesUnits string `json:"queue.max_bytes.units"`
	// QueueType The internal queuing model to use for event buffering.
	QueueType string `json:"queue.type"`
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
