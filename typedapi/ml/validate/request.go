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

package validate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package validate
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/validate/MlValidateJobRequest.ts#L27-L52
type Request struct {
	AnalysisConfig             *types.AnalysisConfig  `json:"analysis_config,omitempty"`
	AnalysisLimits             *types.AnalysisLimits  `json:"analysis_limits,omitempty"`
	DataDescription            *types.DataDescription `json:"data_description,omitempty"`
	Description                *string                `json:"description,omitempty"`
	JobId                      *string                `json:"job_id,omitempty"`
	ModelPlot                  *types.ModelPlotConfig `json:"model_plot,omitempty"`
	ModelSnapshotId            *string                `json:"model_snapshot_id,omitempty"`
	ModelSnapshotRetentionDays *int64                 `json:"model_snapshot_retention_days,omitempty"`
	ResultsIndexName           *string                `json:"results_index_name,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Validate request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "analysis_config":
			if err := dec.Decode(&s.AnalysisConfig); err != nil {
				return fmt.Errorf("%s | %w", "AnalysisConfig", err)
			}

		case "analysis_limits":
			if err := dec.Decode(&s.AnalysisLimits); err != nil {
				return fmt.Errorf("%s | %w", "AnalysisLimits", err)
			}

		case "data_description":
			if err := dec.Decode(&s.DataDescription); err != nil {
				return fmt.Errorf("%s | %w", "DataDescription", err)
			}

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

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return fmt.Errorf("%s | %w", "JobId", err)
			}

		case "model_plot":
			if err := dec.Decode(&s.ModelPlot); err != nil {
				return fmt.Errorf("%s | %w", "ModelPlot", err)
			}

		case "model_snapshot_id":
			if err := dec.Decode(&s.ModelSnapshotId); err != nil {
				return fmt.Errorf("%s | %w", "ModelSnapshotId", err)
			}

		case "model_snapshot_retention_days":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ModelSnapshotRetentionDays", err)
				}
				s.ModelSnapshotRetentionDays = &value
			case float64:
				f := int64(v)
				s.ModelSnapshotRetentionDays = &f
			}

		case "results_index_name":
			if err := dec.Decode(&s.ResultsIndexName); err != nil {
				return fmt.Errorf("%s | %w", "ResultsIndexName", err)
			}

		}
	}
	return nil
}
