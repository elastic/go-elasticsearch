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

// PipelineSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/logstash/_types/Pipeline.ts#L28-L55
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
	// QueueMaxBytes The total capacity of the queue (`queue.type: persisted`) in number of bytes.
	QueueMaxBytes string `json:"queue.max_bytes"`
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

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PipelineBatchDelay", err)
				}
				s.PipelineBatchDelay = value
			case float64:
				f := int(v)
				s.PipelineBatchDelay = f
			}

		case "pipeline.batch.size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PipelineBatchSize", err)
				}
				s.PipelineBatchSize = value
			case float64:
				f := int(v)
				s.PipelineBatchSize = f
			}

		case "pipeline.workers":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PipelineWorkers", err)
				}
				s.PipelineWorkers = value
			case float64:
				f := int(v)
				s.PipelineWorkers = f
			}

		case "queue.checkpoint.writes":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "QueueCheckpointWrites", err)
				}
				s.QueueCheckpointWrites = value
			case float64:
				f := int(v)
				s.QueueCheckpointWrites = f
			}

		case "queue.max_bytes":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueueMaxBytes", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueueMaxBytes = o

		case "queue.type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueueType", err)
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
