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

package updatejob

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package updatejob
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/update_job/MlUpdateJobResponse.ts#L29-L53

type Response struct {
	AllowLazyOpen                        bool                     `json:"allow_lazy_open"`
	AnalysisConfig                       types.AnalysisConfigRead `json:"analysis_config"`
	AnalysisLimits                       types.AnalysisLimits     `json:"analysis_limits"`
	BackgroundPersistInterval            types.Duration           `json:"background_persist_interval,omitempty"`
	CreateTime                           int64                    `json:"create_time"`
	CustomSettings                       map[string]string        `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays int64                    `json:"daily_model_snapshot_retention_after_days"`
	DataDescription                      types.DataDescription    `json:"data_description"`
	DatafeedConfig                       *types.MLDatafeed        `json:"datafeed_config,omitempty"`
	Description                          *string                  `json:"description,omitempty"`
	FinishedTime                         *int64                   `json:"finished_time,omitempty"`
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
	r := &Response{
		CustomSettings: make(map[string]string, 0),
	}
	return r
}
