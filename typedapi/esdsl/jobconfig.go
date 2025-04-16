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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

func (s *_jobConfig) AllowLazyOpen(allowlazyopen bool) *_jobConfig {

	s.v.AllowLazyOpen = &allowlazyopen

	return s
}

func (s *_jobConfig) AnalysisConfig(analysisconfig types.AnalysisConfigVariant) *_jobConfig {

	s.v.AnalysisConfig = *analysisconfig.AnalysisConfigCaster()

	return s
}

func (s *_jobConfig) AnalysisLimits(analysislimits types.AnalysisLimitsVariant) *_jobConfig {

	s.v.AnalysisLimits = analysislimits.AnalysisLimitsCaster()

	return s
}

func (s *_jobConfig) BackgroundPersistInterval(duration types.DurationVariant) *_jobConfig {

	s.v.BackgroundPersistInterval = *duration.DurationCaster()

	return s
}

func (s *_jobConfig) CustomSettings(customsettings json.RawMessage) *_jobConfig {

	s.v.CustomSettings = customsettings

	return s
}

func (s *_jobConfig) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *_jobConfig {

	s.v.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays

	return s
}

func (s *_jobConfig) DataDescription(datadescription types.DataDescriptionVariant) *_jobConfig {

	s.v.DataDescription = *datadescription.DataDescriptionCaster()

	return s
}

func (s *_jobConfig) DatafeedConfig(datafeedconfig types.DatafeedConfigVariant) *_jobConfig {

	s.v.DatafeedConfig = datafeedconfig.DatafeedConfigCaster()

	return s
}

func (s *_jobConfig) Description(description string) *_jobConfig {

	s.v.Description = &description

	return s
}

func (s *_jobConfig) Groups(groups ...string) *_jobConfig {

	for _, v := range groups {

		s.v.Groups = append(s.v.Groups, v)

	}
	return s
}

func (s *_jobConfig) JobId(id string) *_jobConfig {

	s.v.JobId = &id

	return s
}

func (s *_jobConfig) JobType(jobtype string) *_jobConfig {

	s.v.JobType = &jobtype

	return s
}

func (s *_jobConfig) ModelPlotConfig(modelplotconfig types.ModelPlotConfigVariant) *_jobConfig {

	s.v.ModelPlotConfig = modelplotconfig.ModelPlotConfigCaster()

	return s
}

func (s *_jobConfig) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *_jobConfig {

	s.v.ModelSnapshotRetentionDays = &modelsnapshotretentiondays

	return s
}

func (s *_jobConfig) RenormalizationWindowDays(renormalizationwindowdays int64) *_jobConfig {

	s.v.RenormalizationWindowDays = &renormalizationwindowdays

	return s
}

func (s *_jobConfig) ResultsIndexName(indexname string) *_jobConfig {

	s.v.ResultsIndexName = &indexname

	return s
}

func (s *_jobConfig) ResultsRetentionDays(resultsretentiondays int64) *_jobConfig {

	s.v.ResultsRetentionDays = &resultsretentiondays

	return s
}

func (s *_jobConfig) JobConfigCaster() *types.JobConfig {
	return s.v
}
