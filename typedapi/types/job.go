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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// Job type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Job.ts#L51-L75
type Job struct {
	AllowLazyOpen                        bool             `json:"allow_lazy_open"`
	AnalysisConfig                       AnalysisConfig   `json:"analysis_config"`
	AnalysisLimits                       *AnalysisLimits  `json:"analysis_limits,omitempty"`
	BackgroundPersistInterval            *Duration        `json:"background_persist_interval,omitempty"`
	Blocked                              *JobBlocked      `json:"blocked,omitempty"`
	CreateTime                           *DateTime        `json:"create_time,omitempty"`
	CustomSettings                       *CustomSettings  `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays *int64           `json:"daily_model_snapshot_retention_after_days,omitempty"`
	DataDescription                      DataDescription  `json:"data_description"`
	DatafeedConfig                       *Datafeed        `json:"datafeed_config,omitempty"`
	Deleting                             *bool            `json:"deleting,omitempty"`
	Description                          *string          `json:"description,omitempty"`
	FinishedTime                         *DateTime        `json:"finished_time,omitempty"`
	Groups                               []string         `json:"groups,omitempty"`
	JobId                                Id               `json:"job_id"`
	JobType                              *string          `json:"job_type,omitempty"`
	JobVersion                           *VersionString   `json:"job_version,omitempty"`
	ModelPlotConfig                      *ModelPlotConfig `json:"model_plot_config,omitempty"`
	ModelSnapshotId                      *Id              `json:"model_snapshot_id,omitempty"`
	ModelSnapshotRetentionDays           int64            `json:"model_snapshot_retention_days"`
	RenormalizationWindowDays            *int64           `json:"renormalization_window_days,omitempty"`
	ResultsIndexName                     IndexName        `json:"results_index_name"`
	ResultsRetentionDays                 *int64           `json:"results_retention_days,omitempty"`
}

// JobBuilder holds Job struct and provides a builder API.
type JobBuilder struct {
	v *Job
}

// NewJob provides a builder for the Job struct.
func NewJobBuilder() *JobBuilder {
	r := JobBuilder{
		&Job{},
	}

	return &r
}

// Build finalize the chain and returns the Job struct
func (rb *JobBuilder) Build() Job {
	return *rb.v
}

func (rb *JobBuilder) AllowLazyOpen(allowlazyopen bool) *JobBuilder {
	rb.v.AllowLazyOpen = allowlazyopen
	return rb
}

func (rb *JobBuilder) AnalysisConfig(analysisconfig *AnalysisConfigBuilder) *JobBuilder {
	v := analysisconfig.Build()
	rb.v.AnalysisConfig = v
	return rb
}

func (rb *JobBuilder) AnalysisLimits(analysislimits *AnalysisLimitsBuilder) *JobBuilder {
	v := analysislimits.Build()
	rb.v.AnalysisLimits = &v
	return rb
}

func (rb *JobBuilder) BackgroundPersistInterval(backgroundpersistinterval *DurationBuilder) *JobBuilder {
	v := backgroundpersistinterval.Build()
	rb.v.BackgroundPersistInterval = &v
	return rb
}

func (rb *JobBuilder) Blocked(blocked *JobBlockedBuilder) *JobBuilder {
	v := blocked.Build()
	rb.v.Blocked = &v
	return rb
}

func (rb *JobBuilder) CreateTime(createtime *DateTimeBuilder) *JobBuilder {
	v := createtime.Build()
	rb.v.CreateTime = &v
	return rb
}

func (rb *JobBuilder) CustomSettings(customsettings *CustomSettingsBuilder) *JobBuilder {
	v := customsettings.Build()
	rb.v.CustomSettings = &v
	return rb
}

func (rb *JobBuilder) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *JobBuilder {
	rb.v.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays
	return rb
}

func (rb *JobBuilder) DataDescription(datadescription *DataDescriptionBuilder) *JobBuilder {
	v := datadescription.Build()
	rb.v.DataDescription = v
	return rb
}

func (rb *JobBuilder) DatafeedConfig(datafeedconfig *DatafeedBuilder) *JobBuilder {
	v := datafeedconfig.Build()
	rb.v.DatafeedConfig = &v
	return rb
}

func (rb *JobBuilder) Deleting(deleting bool) *JobBuilder {
	rb.v.Deleting = &deleting
	return rb
}

func (rb *JobBuilder) Description(description string) *JobBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *JobBuilder) FinishedTime(finishedtime *DateTimeBuilder) *JobBuilder {
	v := finishedtime.Build()
	rb.v.FinishedTime = &v
	return rb
}

func (rb *JobBuilder) Groups(groups ...string) *JobBuilder {
	rb.v.Groups = groups
	return rb
}

func (rb *JobBuilder) JobId(jobid Id) *JobBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *JobBuilder) JobType(jobtype string) *JobBuilder {
	rb.v.JobType = &jobtype
	return rb
}

func (rb *JobBuilder) JobVersion(jobversion VersionString) *JobBuilder {
	rb.v.JobVersion = &jobversion
	return rb
}

func (rb *JobBuilder) ModelPlotConfig(modelplotconfig *ModelPlotConfigBuilder) *JobBuilder {
	v := modelplotconfig.Build()
	rb.v.ModelPlotConfig = &v
	return rb
}

func (rb *JobBuilder) ModelSnapshotId(modelsnapshotid Id) *JobBuilder {
	rb.v.ModelSnapshotId = &modelsnapshotid
	return rb
}

func (rb *JobBuilder) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *JobBuilder {
	rb.v.ModelSnapshotRetentionDays = modelsnapshotretentiondays
	return rb
}

func (rb *JobBuilder) RenormalizationWindowDays(renormalizationwindowdays int64) *JobBuilder {
	rb.v.RenormalizationWindowDays = &renormalizationwindowdays
	return rb
}

func (rb *JobBuilder) ResultsIndexName(resultsindexname IndexName) *JobBuilder {
	rb.v.ResultsIndexName = resultsindexname
	return rb
}

func (rb *JobBuilder) ResultsRetentionDays(resultsretentiondays int64) *JobBuilder {
	rb.v.ResultsRetentionDays = &resultsretentiondays
	return rb
}
