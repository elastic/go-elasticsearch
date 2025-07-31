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

package putjob

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package putjob
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/put_job/MlPutJobResponse.ts#L29-L52
type Response struct {
	AllowLazyOpen                        bool                     `json:"allow_lazy_open"`
	AnalysisConfig                       types.AnalysisConfigRead `json:"analysis_config"`
	AnalysisLimits                       types.AnalysisLimits     `json:"analysis_limits"`
	BackgroundPersistInterval            types.Duration           `json:"background_persist_interval,omitempty"`
	CreateTime                           types.DateTime           `json:"create_time"`
	CustomSettings                       json.RawMessage          `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays int64                    `json:"daily_model_snapshot_retention_after_days"`
	DataDescription                      types.DataDescription    `json:"data_description"`
	DatafeedConfig                       *types.MLDatafeed        `json:"datafeed_config,omitempty"`
	Description                          *string                  `json:"description,omitempty"`
	Groups                               []string                 `json:"groups,omitempty"`
	JobId                                string                   `json:"job_id"`
	JobType                              string                   `json:"job_type"`
	JobVersion                           string                   `json:"job_version"`
	ModelPlotConfig                      *types.ModelPlotConfig   `json:"model_plot_config,omitempty"`
	ModelSnapshotId                      *string                  `json:"model_snapshot_id,omitempty"`
	ModelSnapshotRetentionDays           int64                    `json:"model_snapshot_retention_days"`
	RenormalizationWindowDays            *int64                   `json:"renormalization_window_days,omitempty"`
	ResultsIndexName                     string                   `json:"results_index_name"`
	ResultsRetentionDays                 *int64                   `json:"results_retention_days,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowLazyOpen", err)
				}
				s.AllowLazyOpen = value
			case bool:
				s.AllowLazyOpen = v
			}

		case "analysis_config":
			if err := dec.Decode(&s.AnalysisConfig); err != nil {
				return fmt.Errorf("%s | %w", "AnalysisConfig", err)
			}

		case "analysis_limits":
			if err := dec.Decode(&s.AnalysisLimits); err != nil {
				return fmt.Errorf("%s | %w", "AnalysisLimits", err)
			}

		case "background_persist_interval":
			if err := dec.Decode(&s.BackgroundPersistInterval); err != nil {
				return fmt.Errorf("%s | %w", "BackgroundPersistInterval", err)
			}

		case "create_time":
			if err := dec.Decode(&s.CreateTime); err != nil {
				return fmt.Errorf("%s | %w", "CreateTime", err)
			}

		case "custom_settings":
			if err := dec.Decode(&s.CustomSettings); err != nil {
				return fmt.Errorf("%s | %w", "CustomSettings", err)
			}

		case "daily_model_snapshot_retention_after_days":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DailyModelSnapshotRetentionAfterDays", err)
				}
				s.DailyModelSnapshotRetentionAfterDays = value
			case float64:
				f := int64(v)
				s.DailyModelSnapshotRetentionAfterDays = f
			}

		case "data_description":
			if err := dec.Decode(&s.DataDescription); err != nil {
				return fmt.Errorf("%s | %w", "DataDescription", err)
			}

		case "datafeed_config":
			if err := dec.Decode(&s.DatafeedConfig); err != nil {
				return fmt.Errorf("%s | %w", "DatafeedConfig", err)
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

		case "groups":
			if err := dec.Decode(&s.Groups); err != nil {
				return fmt.Errorf("%s | %w", "Groups", err)
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return fmt.Errorf("%s | %w", "JobId", err)
			}

		case "job_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "JobType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.JobType = o

		case "job_version":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "JobVersion", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.JobVersion = o

		case "model_plot_config":
			if err := dec.Decode(&s.ModelPlotConfig); err != nil {
				return fmt.Errorf("%s | %w", "ModelPlotConfig", err)
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
				s.ModelSnapshotRetentionDays = value
			case float64:
				f := int64(v)
				s.ModelSnapshotRetentionDays = f
			}

		case "renormalization_window_days":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RenormalizationWindowDays", err)
				}
				s.RenormalizationWindowDays = &value
			case float64:
				f := int64(v)
				s.RenormalizationWindowDays = &f
			}

		case "results_index_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ResultsIndexName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResultsIndexName = o

		case "results_retention_days":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ResultsRetentionDays", err)
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
