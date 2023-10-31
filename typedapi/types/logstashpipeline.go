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

// LogstashPipeline type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/logstash/_types/Pipeline.ts#L60-L92
type LogstashPipeline struct {
	// Description Description of the pipeline.
	// This description is not used by Elasticsearch or Logstash.
	Description string `json:"description"`
	// LastModified Date the pipeline was last updated.
	// Must be in the `yyyy-MM-dd'T'HH:mm:ss.SSSZZ` strict_date_time format.
	LastModified DateTime `json:"last_modified"`
	// Pipeline Configuration for the pipeline.
	Pipeline string `json:"pipeline"`
	// PipelineMetadata Optional metadata about the pipeline.
	// May have any contents.
	// This metadata is not generated or used by Elasticsearch or Logstash.
	PipelineMetadata PipelineMetadata `json:"pipeline_metadata"`
	// PipelineSettings Settings for the pipeline.
	// Supports only flat keys in dot notation.
	PipelineSettings PipelineSettings `json:"pipeline_settings"`
	// Username User who last updated the pipeline.
	Username string `json:"username"`
}

func (s *LogstashPipeline) UnmarshalJSON(data []byte) error {

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
			s.Description = o

		case "last_modified":
			if err := dec.Decode(&s.LastModified); err != nil {
				return err
			}

		case "pipeline":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pipeline = o

		case "pipeline_metadata":
			if err := dec.Decode(&s.PipelineMetadata); err != nil {
				return err
			}

		case "pipeline_settings":
			if err := dec.Decode(&s.PipelineSettings); err != nil {
				return err
			}

		case "username":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Username = o

		}
	}
	return nil
}

// NewLogstashPipeline returns a LogstashPipeline.
func NewLogstashPipeline() *LogstashPipeline {
	r := &LogstashPipeline{}

	return r
}
