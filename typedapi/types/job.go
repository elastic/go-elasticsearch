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

// Job type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Job.ts#L51-L75
type Job struct {
	AllowLazyOpen                        bool             `json:"allow_lazy_open"`
	AnalysisConfig                       AnalysisConfig   `json:"analysis_config"`
	AnalysisLimits                       *AnalysisLimits  `json:"analysis_limits,omitempty"`
	BackgroundPersistInterval            Duration         `json:"background_persist_interval,omitempty"`
	Blocked                              *JobBlocked      `json:"blocked,omitempty"`
	CreateTime                           DateTime         `json:"create_time,omitempty"`
	CustomSettings                       json.RawMessage  `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays *int64           `json:"daily_model_snapshot_retention_after_days,omitempty"`
	DataDescription                      DataDescription  `json:"data_description"`
	DatafeedConfig                       *MLDatafeed      `json:"datafeed_config,omitempty"`
	Deleting                             *bool            `json:"deleting,omitempty"`
	Description                          *string          `json:"description,omitempty"`
	FinishedTime                         DateTime         `json:"finished_time,omitempty"`
	Groups                               []string         `json:"groups,omitempty"`
	JobId                                string           `json:"job_id"`
	JobType                              *string          `json:"job_type,omitempty"`
	JobVersion                           *string          `json:"job_version,omitempty"`
	ModelPlotConfig                      *ModelPlotConfig `json:"model_plot_config,omitempty"`
	ModelSnapshotId                      *string          `json:"model_snapshot_id,omitempty"`
	ModelSnapshotRetentionDays           int64            `json:"model_snapshot_retention_days"`
	RenormalizationWindowDays            *int64           `json:"renormalization_window_days,omitempty"`
	ResultsIndexName                     string           `json:"results_index_name"`
	ResultsRetentionDays                 *int64           `json:"results_retention_days,omitempty"`
}

func (s *Job) UnmarshalJSON(data []byte) error {

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

		case "allow_lazy_open":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowLazyOpen = value
			case bool:
				s.AllowLazyOpen = v
			}

		case "analysis_config":
			if err := dec.Decode(&s.AnalysisConfig); err != nil {
				return err
			}

		case "analysis_limits":
			if err := dec.Decode(&s.AnalysisLimits); err != nil {
				return err
			}

		case "background_persist_interval":
			if err := dec.Decode(&s.BackgroundPersistInterval); err != nil {
				return err
			}

		case "blocked":
			if err := dec.Decode(&s.Blocked); err != nil {
				return err
			}

		case "create_time":
			if err := dec.Decode(&s.CreateTime); err != nil {
				return err
			}

		case "custom_settings":
			if err := dec.Decode(&s.CustomSettings); err != nil {
				return err
			}

		case "daily_model_snapshot_retention_after_days":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DailyModelSnapshotRetentionAfterDays = &value
			case float64:
				f := int64(v)
				s.DailyModelSnapshotRetentionAfterDays = &f
			}

		case "data_description":
			if err := dec.Decode(&s.DataDescription); err != nil {
				return err
			}

		case "datafeed_config":
			if err := dec.Decode(&s.DatafeedConfig); err != nil {
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
			s.Description = &o

		case "finished_time":
			if err := dec.Decode(&s.FinishedTime); err != nil {
				return err
			}

		case "groups":
			if err := dec.Decode(&s.Groups); err != nil {
				return err
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "job_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.JobType = &o

		case "job_version":
			if err := dec.Decode(&s.JobVersion); err != nil {
				return err
			}

		case "model_plot_config":
			if err := dec.Decode(&s.ModelPlotConfig); err != nil {
				return err
			}

		case "model_snapshot_id":
			if err := dec.Decode(&s.ModelSnapshotId); err != nil {
				return err
			}

		case "model_snapshot_retention_days":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ModelSnapshotRetentionDays = value
			case float64:
				f := int64(v)
				s.ModelSnapshotRetentionDays = f
			}

		case "renormalization_window_days":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RenormalizationWindowDays = &value
			case float64:
				f := int64(v)
				s.RenormalizationWindowDays = &f
			}

		case "results_index_name":
			if err := dec.Decode(&s.ResultsIndexName); err != nil {
				return err
			}

		case "results_retention_days":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ResultsRetentionDays = &value
			case float64:
				f := int64(v)
				s.ResultsRetentionDays = &f
			}

		}
	}
	return nil
}

// NewJob returns a Job.
func NewJob() *Job {
	r := &Job{}

	return r
}
