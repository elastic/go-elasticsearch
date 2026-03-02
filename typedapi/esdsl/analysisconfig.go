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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _analysisConfig struct {
	v *types.AnalysisConfig
}

func NewAnalysisConfig() *_analysisConfig {

	return &_analysisConfig{v: types.NewAnalysisConfig()}

}

func (s *_analysisConfig) BucketSpan(duration types.DurationVariant) *_analysisConfig {

	s.v.BucketSpan = *duration.DurationCaster()

	return s
}

func (s *_analysisConfig) CategorizationAnalyzer(categorizationanalyzer types.CategorizationAnalyzerVariant) *_analysisConfig {

	s.v.CategorizationAnalyzer = *categorizationanalyzer.CategorizationAnalyzerCaster()

	return s
}

func (s *_analysisConfig) CategorizationFieldName(field string) *_analysisConfig {

	s.v.CategorizationFieldName = &field

	return s
}

func (s *_analysisConfig) CategorizationFilters(categorizationfilters ...string) *_analysisConfig {

	for _, v := range categorizationfilters {

		s.v.CategorizationFilters = append(s.v.CategorizationFilters, v)

	}
	return s
}

func (s *_analysisConfig) Detectors(detectors ...types.DetectorVariant) *_analysisConfig {

	for _, v := range detectors {

		s.v.Detectors = append(s.v.Detectors, *v.DetectorCaster())

	}
	return s
}

func (s *_analysisConfig) Influencers(influencers ...string) *_analysisConfig {

	for _, v := range influencers {

		s.v.Influencers = append(s.v.Influencers, v)

	}
	return s
}

func (s *_analysisConfig) Latency(duration types.DurationVariant) *_analysisConfig {

	s.v.Latency = *duration.DurationCaster()

	return s
}

func (s *_analysisConfig) ModelPruneWindow(duration types.DurationVariant) *_analysisConfig {

	s.v.ModelPruneWindow = *duration.DurationCaster()

	return s
}

func (s *_analysisConfig) MultivariateByFields(multivariatebyfields bool) *_analysisConfig {

	s.v.MultivariateByFields = &multivariatebyfields

	return s
}

func (s *_analysisConfig) PerPartitionCategorization(perpartitioncategorization types.PerPartitionCategorizationVariant) *_analysisConfig {

	s.v.PerPartitionCategorization = perpartitioncategorization.PerPartitionCategorizationCaster()

	return s
}

func (s *_analysisConfig) SummaryCountFieldName(field string) *_analysisConfig {

	s.v.SummaryCountFieldName = &field

	return s
}

func (s *_analysisConfig) AnalysisConfigCaster() *types.AnalysisConfig {
	return s.v
}
