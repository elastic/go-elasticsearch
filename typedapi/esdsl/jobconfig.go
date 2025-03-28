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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _jobConfig struct {
	v *types.JobConfig
}

func NewJobConfig(analysisconfig types.AnalysisConfigVariant, datadescription types.DataDescriptionVariant) *_jobConfig {

	tmp := &_jobConfig{v: types.NewJobConfig()}

	tmp.AnalysisConfig(analysisconfig)

	tmp.DataDescription(datadescription)

	return tmp

}

// Advanced configuration option. Specifies whether this job can open when there
// is insufficient machine learning node capacity for it to be immediately
// assigned to a node.
func (s *_jobConfig) AllowLazyOpen(allowlazyopen bool) *_jobConfig {

	s.v.AllowLazyOpen = &allowlazyopen

	return s
}

// The analysis configuration, which specifies how to analyze the data.
// After you create a job, you cannot change the analysis configuration; all the
// properties are informational.
func (s *_jobConfig) AnalysisConfig(analysisconfig types.AnalysisConfigVariant) *_jobConfig {

	s.v.AnalysisConfig = *analysisconfig.AnalysisConfigCaster()

	return s
}

// Limits can be applied for the resources required to hold the mathematical
// models in memory.
// These limits are approximate and can be set per job.
// They do not control the memory used by other processes, for example the
// Elasticsearch Java processes.
func (s *_jobConfig) AnalysisLimits(analysislimits types.AnalysisLimitsVariant) *_jobConfig {

	s.v.AnalysisLimits = analysislimits.AnalysisLimitsCaster()

	return s
}

// Advanced configuration option.
// The time between each periodic persistence of the model.
// The default value is a randomized value between 3 to 4 hours, which avoids
// all jobs persisting at exactly the same time.
// The smallest allowed value is 1 hour.
func (s *_jobConfig) BackgroundPersistInterval(duration types.DurationVariant) *_jobConfig {

	s.v.BackgroundPersistInterval = *duration.DurationCaster()

	return s
}

// Advanced configuration option.
// Contains custom metadata about the job.
func (s *_jobConfig) CustomSettings(customsettings json.RawMessage) *_jobConfig {

	s.v.CustomSettings = customsettings

	return s
}

// Advanced configuration option, which affects the automatic removal of old
// model snapshots for this job.
// It specifies a period of time (in days) after which only the first snapshot
// per day is retained.
// This period is relative to the timestamp of the most recent snapshot for this
// job.
func (s *_jobConfig) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *_jobConfig {

	s.v.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays

	return s
}

// The data description defines the format of the input data when you send data
// to the job by using the post data API.
// Note that when configure a datafeed, these properties are automatically set.
func (s *_jobConfig) DataDescription(datadescription types.DataDescriptionVariant) *_jobConfig {

	s.v.DataDescription = *datadescription.DataDescriptionCaster()

	return s
}

// The datafeed, which retrieves data from Elasticsearch for analysis by the
// job.
// You can associate only one datafeed with each anomaly detection job.
func (s *_jobConfig) DatafeedConfig(datafeedconfig types.DatafeedConfigVariant) *_jobConfig {

	s.v.DatafeedConfig = datafeedconfig.DatafeedConfigCaster()

	return s
}

// A description of the job.
func (s *_jobConfig) Description(description string) *_jobConfig {

	s.v.Description = &description

	return s
}

// A list of job groups. A job can belong to no groups or many.
func (s *_jobConfig) Groups(groups ...string) *_jobConfig {

	for _, v := range groups {

		s.v.Groups = append(s.v.Groups, v)

	}
	return s
}

// Identifier for the anomaly detection job.
// This identifier can contain lowercase alphanumeric characters (a-z and 0-9),
// hyphens, and underscores.
// It must start and end with alphanumeric characters.
func (s *_jobConfig) JobId(id string) *_jobConfig {

	s.v.JobId = &id

	return s
}

// Reserved for future use, currently set to `anomaly_detector`.
func (s *_jobConfig) JobType(jobtype string) *_jobConfig {

	s.v.JobType = &jobtype

	return s
}

// This advanced configuration option stores model information along with the
// results.
// It provides a more detailed view into anomaly detection.
// Model plot provides a simplified and indicative view of the model and its
// bounds.
func (s *_jobConfig) ModelPlotConfig(modelplotconfig types.ModelPlotConfigVariant) *_jobConfig {

	s.v.ModelPlotConfig = modelplotconfig.ModelPlotConfigCaster()

	return s
}

// Advanced configuration option, which affects the automatic removal of old
// model snapshots for this job.
// It specifies the maximum period of time (in days) that snapshots are
// retained.
// This period is relative to the timestamp of the most recent snapshot for this
// job.
// The default value is `10`, which means snapshots ten days older than the
// newest snapshot are deleted.
func (s *_jobConfig) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *_jobConfig {

	s.v.ModelSnapshotRetentionDays = &modelsnapshotretentiondays

	return s
}

// Advanced configuration option.
// The period over which adjustments to the score are applied, as new data is
// seen.
// The default value is the longer of 30 days or 100 `bucket_spans`.
func (s *_jobConfig) RenormalizationWindowDays(renormalizationwindowdays int64) *_jobConfig {

	s.v.RenormalizationWindowDays = &renormalizationwindowdays

	return s
}

// A text string that affects the name of the machine learning results index.
// The default value is `shared`, which generates an index named
// `.ml-anomalies-shared`.
func (s *_jobConfig) ResultsIndexName(indexname string) *_jobConfig {

	s.v.ResultsIndexName = &indexname

	return s
}

// Advanced configuration option.
// The period of time (in days) that results are retained.
// Age is calculated relative to the timestamp of the latest bucket result.
// If this property has a non-null value, once per day at 00:30 (server time),
// results that are the specified number of days older than the latest bucket
// result are deleted from Elasticsearch.
// The default value is null, which means all results are retained.
// Annotations generated by the system also count as results for retention
// purposes; they are deleted after the same number of days as results.
// Annotations added by users are retained forever.
func (s *_jobConfig) ResultsRetentionDays(resultsretentiondays int64) *_jobConfig {

	s.v.ResultsRetentionDays = &resultsretentiondays

	return s
}

func (s *_jobConfig) JobConfigCaster() *types.JobConfig {
	return s.v
}
