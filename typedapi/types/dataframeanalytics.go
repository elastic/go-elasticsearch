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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dataframestate"
)

// DataframeAnalytics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeAnalytics.ts#L324-L344
type DataframeAnalytics struct {
	// AnalysisStats An object containing information about the analysis job.
	AnalysisStats *DataframeAnalyticsStatsContainer `json:"analysis_stats,omitempty"`
	// AssignmentExplanation For running jobs only, contains messages relating to the selection of a node
	// to run the job.
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// DataCounts An object that provides counts for the quantity of documents skipped, used in
	// training, or available for testing.
	DataCounts DataframeAnalyticsStatsDataCounts `json:"data_counts"`
	// Id The unique identifier of the data frame analytics job.
	Id string `json:"id"`
	// MemoryUsage An object describing memory usage of the analytics. It is present only after
	// the job is started and memory usage is reported.
	MemoryUsage DataframeAnalyticsStatsMemoryUsage `json:"memory_usage"`
	// Node Contains properties for the node that runs the job. This information is
	// available only for running jobs.
	Node *NodeAttributes `json:"node,omitempty"`
	// Progress The progress report of the data frame analytics job by phase.
	Progress []DataframeAnalyticsStatsProgress `json:"progress"`
	// State The status of the data frame analytics job, which can be one of the following
	// values: failed, started, starting, stopping, stopped.
	State dataframestate.DataframeState `json:"state"`
}

func (s *DataframeAnalytics) UnmarshalJSON(data []byte) error {

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

		case "analysis_stats":
			if err := dec.Decode(&s.AnalysisStats); err != nil {
				return err
			}

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "memory_usage":
			if err := dec.Decode(&s.MemoryUsage); err != nil {
				return err
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return err
			}

		case "progress":
			if err := dec.Decode(&s.Progress); err != nil {
				return err
			}

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeAnalytics returns a DataframeAnalytics.
func NewDataframeAnalytics() *DataframeAnalytics {
	r := &DataframeAnalytics{}

	return r
}
