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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jobstate"
)

// JobStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Job.ts#L96-L107
type JobStats struct {
	AssignmentExplanation *string               `json:"assignment_explanation,omitempty"`
	DataCounts            DataCounts            `json:"data_counts"`
	Deleting              *bool                 `json:"deleting,omitempty"`
	ForecastsStats        JobForecastStatistics `json:"forecasts_stats"`
	JobId                 string                `json:"job_id"`
	ModelSizeStats        ModelSizeStats        `json:"model_size_stats"`
	Node                  *DiscoveryNode        `json:"node,omitempty"`
	OpenTime              DateTime              `json:"open_time,omitempty"`
	State                 jobstate.JobState     `json:"state"`
	TimingStats           JobTimingStats        `json:"timing_stats"`
}

func (s *JobStats) UnmarshalJSON(data []byte) error {

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

		case "assignment_explanation":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AssignmentExplanation = &o

		case "data_counts":
			if err := dec.Decode(&s.DataCounts); err != nil {
				return err
			}

		case "deleting":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Deleting = &value
			case bool:
				s.Deleting = &v
			}

		case "forecasts_stats":
			if err := dec.Decode(&s.ForecastsStats); err != nil {
				return err
			}

		case "job_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.JobId = o

		case "model_size_stats":
			if err := dec.Decode(&s.ModelSizeStats); err != nil {
				return err
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return err
			}

		case "open_time":
			if err := dec.Decode(&s.OpenTime); err != nil {
				return err
			}

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		case "timing_stats":
			if err := dec.Decode(&s.TimingStats); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewJobStats returns a JobStats.
func NewJobStats() *JobStats {
	r := &JobStats{}

	return r
}
