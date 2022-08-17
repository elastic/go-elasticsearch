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

// JobConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Job.ts#L77-L95
type JobConfig struct {
	AllowLazyOpen                        *bool            `json:"allow_lazy_open,omitempty"`
	AnalysisConfig                       AnalysisConfig   `json:"analysis_config"`
	AnalysisLimits                       *AnalysisLimits  `json:"analysis_limits,omitempty"`
	BackgroundPersistInterval            *Duration        `json:"background_persist_interval,omitempty"`
	CustomSettings                       *CustomSettings  `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays *int64           `json:"daily_model_snapshot_retention_after_days,omitempty"`
	DataDescription                      DataDescription  `json:"data_description"`
	DatafeedConfig                       *DatafeedConfig  `json:"datafeed_config,omitempty"`
	Description                          *string          `json:"description,omitempty"`
	Groups                               []string         `json:"groups,omitempty"`
	JobId                                *Id              `json:"job_id,omitempty"`
	JobType                              *string          `json:"job_type,omitempty"`
	ModelPlotConfig                      *ModelPlotConfig `json:"model_plot_config,omitempty"`
	ModelSnapshotRetentionDays           *int64           `json:"model_snapshot_retention_days,omitempty"`
	RenormalizationWindowDays            *int64           `json:"renormalization_window_days,omitempty"`
	ResultsIndexName                     *IndexName       `json:"results_index_name,omitempty"`
	ResultsRetentionDays                 *int64           `json:"results_retention_days,omitempty"`
}

// JobConfigBuilder holds JobConfig struct and provides a builder API.
type JobConfigBuilder struct {
	v *JobConfig
}

// NewJobConfig provides a builder for the JobConfig struct.
func NewJobConfigBuilder() *JobConfigBuilder {
	r := JobConfigBuilder{
		&JobConfig{},
	}

	return &r
}

// Build finalize the chain and returns the JobConfig struct
func (rb *JobConfigBuilder) Build() JobConfig {
	return *rb.v
}

func (rb *JobConfigBuilder) AllowLazyOpen(allowlazyopen bool) *JobConfigBuilder {
	rb.v.AllowLazyOpen = &allowlazyopen
	return rb
}

func (rb *JobConfigBuilder) AnalysisConfig(analysisconfig *AnalysisConfigBuilder) *JobConfigBuilder {
	v := analysisconfig.Build()
	rb.v.AnalysisConfig = v
	return rb
}

func (rb *JobConfigBuilder) AnalysisLimits(analysislimits *AnalysisLimitsBuilder) *JobConfigBuilder {
	v := analysislimits.Build()
	rb.v.AnalysisLimits = &v
	return rb
}

func (rb *JobConfigBuilder) BackgroundPersistInterval(backgroundpersistinterval *DurationBuilder) *JobConfigBuilder {
	v := backgroundpersistinterval.Build()
	rb.v.BackgroundPersistInterval = &v
	return rb
}

func (rb *JobConfigBuilder) CustomSettings(customsettings *CustomSettingsBuilder) *JobConfigBuilder {
	v := customsettings.Build()
	rb.v.CustomSettings = &v
	return rb
}

func (rb *JobConfigBuilder) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *JobConfigBuilder {
	rb.v.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays
	return rb
}

func (rb *JobConfigBuilder) DataDescription(datadescription *DataDescriptionBuilder) *JobConfigBuilder {
	v := datadescription.Build()
	rb.v.DataDescription = v
	return rb
}

func (rb *JobConfigBuilder) DatafeedConfig(datafeedconfig *DatafeedConfigBuilder) *JobConfigBuilder {
	v := datafeedconfig.Build()
	rb.v.DatafeedConfig = &v
	return rb
}

func (rb *JobConfigBuilder) Description(description string) *JobConfigBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *JobConfigBuilder) Groups(groups ...string) *JobConfigBuilder {
	rb.v.Groups = groups
	return rb
}

func (rb *JobConfigBuilder) JobId(jobid Id) *JobConfigBuilder {
	rb.v.JobId = &jobid
	return rb
}

func (rb *JobConfigBuilder) JobType(jobtype string) *JobConfigBuilder {
	rb.v.JobType = &jobtype
	return rb
}

func (rb *JobConfigBuilder) ModelPlotConfig(modelplotconfig *ModelPlotConfigBuilder) *JobConfigBuilder {
	v := modelplotconfig.Build()
	rb.v.ModelPlotConfig = &v
	return rb
}

func (rb *JobConfigBuilder) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *JobConfigBuilder {
	rb.v.ModelSnapshotRetentionDays = &modelsnapshotretentiondays
	return rb
}

func (rb *JobConfigBuilder) RenormalizationWindowDays(renormalizationwindowdays int64) *JobConfigBuilder {
	rb.v.RenormalizationWindowDays = &renormalizationwindowdays
	return rb
}

func (rb *JobConfigBuilder) ResultsIndexName(resultsindexname IndexName) *JobConfigBuilder {
	rb.v.ResultsIndexName = &resultsindexname
	return rb
}

func (rb *JobConfigBuilder) ResultsRetentionDays(resultsretentiondays int64) *JobConfigBuilder {
	rb.v.ResultsRetentionDays = &resultsretentiondays
	return rb
}
