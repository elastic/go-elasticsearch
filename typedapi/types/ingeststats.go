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

// IngestStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/_types/Stats.ts#L358-L396
type IngestStats struct {
	// Count Total number of documents ingested during the lifetime of this node.
	Count int64 `json:"count"`
	// Current Total number of documents currently being ingested.
	Current int64 `json:"current"`
	// Failed Total number of failed ingest operations during the lifetime of this node.
	Failed int64 `json:"failed"`
	// IngestedAsFirstPipelineInBytes Total number of bytes of all documents ingested by the pipeline.
	// This field is only present on pipelines which are the first to process a
	// document.
	// Thus, it is not present on pipelines which only serve as a final pipeline
	// after a default pipeline, a pipeline run after a reroute processor, or
	// pipelines in pipeline processors.
	IngestedAsFirstPipelineInBytes int64 `json:"ingested_as_first_pipeline_in_bytes"`
	// Processors Total number of ingest processors.
	Processors []map[string]KeyedProcessor `json:"processors"`
	// ProducedAsFirstPipelineInBytes Total number of bytes of all documents produced by the pipeline.
	// This field is only present on pipelines which are the first to process a
	// document.
	// Thus, it is not present on pipelines which only serve as a final pipeline
	// after a default pipeline, a pipeline run after a reroute processor, or
	// pipelines in pipeline processors.
	// In situations where there are subsequent pipelines, the value represents the
	// size of the document after all pipelines have run.
	ProducedAsFirstPipelineInBytes int64 `json:"produced_as_first_pipeline_in_bytes"`
	// TimeInMillis Total time, in milliseconds, spent preprocessing ingest documents during the
	// lifetime of this node.
	TimeInMillis int64 `json:"time_in_millis"`
}

func (s *IngestStats) UnmarshalJSON(data []byte) error {

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

		case "count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "current":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Current", err)
				}
				s.Current = value
			case float64:
				f := int64(v)
				s.Current = f
			}

		case "failed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Failed", err)
				}
				s.Failed = value
			case float64:
				f := int64(v)
				s.Failed = f
			}

		case "ingested_as_first_pipeline_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IngestedAsFirstPipelineInBytes", err)
				}
				s.IngestedAsFirstPipelineInBytes = value
			case float64:
				f := int64(v)
				s.IngestedAsFirstPipelineInBytes = f
			}

		case "processors":
			if err := dec.Decode(&s.Processors); err != nil {
				return fmt.Errorf("%s | %w", "Processors", err)
			}

		case "produced_as_first_pipeline_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ProducedAsFirstPipelineInBytes", err)
				}
				s.ProducedAsFirstPipelineInBytes = value
			case float64:
				f := int64(v)
				s.ProducedAsFirstPipelineInBytes = f
			}

		case "time_in_millis":
			if err := dec.Decode(&s.TimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TimeInMillis", err)
			}

		}
	}
	return nil
}

// NewIngestStats returns a IngestStats.
func NewIngestStats() *IngestStats {
	r := &IngestStats{}

	return r
}
