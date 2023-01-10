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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// Job type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/Job.ts#L51-L75
type Job struct {
	AllowLazyOpen                        bool             `json:"allow_lazy_open"`
	AnalysisConfig                       AnalysisConfig   `json:"analysis_config"`
	AnalysisLimits                       *AnalysisLimits  `json:"analysis_limits,omitempty"`
	BackgroundPersistInterval            *Duration        `json:"background_persist_interval,omitempty"`
	Blocked                              *JobBlocked      `json:"blocked,omitempty"`
	CreateTime                           *DateTime        `json:"create_time,omitempty"`
	CustomSettings                       interface{}      `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays *int64           `json:"daily_model_snapshot_retention_after_days,omitempty"`
	DataDescription                      DataDescription  `json:"data_description"`
	DatafeedConfig                       *MLDatafeed      `json:"datafeed_config,omitempty"`
	Deleting                             *bool            `json:"deleting,omitempty"`
	Description                          *string          `json:"description,omitempty"`
	FinishedTime                         *DateTime        `json:"finished_time,omitempty"`
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

// NewJob returns a Job.
func NewJob() *Job {
	r := &Job{}

	return r
}
