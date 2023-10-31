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

// MachineLearning type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/xpack/usage/types.ts#L372-L379
type MachineLearning struct {
	Available              bool                     `json:"available"`
	DataFrameAnalyticsJobs MlDataFrameAnalyticsJobs `json:"data_frame_analytics_jobs"`
	Datafeeds              map[string]XpackDatafeed `json:"datafeeds"`
	Enabled                bool                     `json:"enabled"`
	Inference              MlInference              `json:"inference"`
	// Jobs Job usage statistics. The `_all` entry is always present and gathers
	// statistics for all jobs.
	Jobs      map[string]JobUsage `json:"jobs"`
	NodeCount int                 `json:"node_count"`
}

func (s *MachineLearning) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "data_frame_analytics_jobs":
			if err := dec.Decode(&s.DataFrameAnalyticsJobs); err != nil {
				return err
			}

		case "datafeeds":
			if s.Datafeeds == nil {
				s.Datafeeds = make(map[string]XpackDatafeed, 0)
			}
			if err := dec.Decode(&s.Datafeeds); err != nil {
				return err
			}

		case "enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "inference":
			if err := dec.Decode(&s.Inference); err != nil {
				return err
			}

		case "jobs":
			if s.Jobs == nil {
				s.Jobs = make(map[string]JobUsage, 0)
			}
			if err := dec.Decode(&s.Jobs); err != nil {
				return err
			}

		case "node_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NodeCount = value
			case float64:
				f := int(v)
				s.NodeCount = f
			}

		}
	}
	return nil
}

// NewMachineLearning returns a MachineLearning.
func NewMachineLearning() *MachineLearning {
	r := &MachineLearning{
		Datafeeds: make(map[string]XpackDatafeed, 0),
		Jobs:      make(map[string]JobUsage, 0),
	}

	return r
}
