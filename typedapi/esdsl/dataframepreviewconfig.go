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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframePreviewConfig struct {
	v *types.DataframePreviewConfig
}

func NewDataframePreviewConfig(analysis types.DataframeAnalysisContainerVariant, source types.DataframeAnalyticsSourceVariant) *_dataframePreviewConfig {

	tmp := &_dataframePreviewConfig{v: types.NewDataframePreviewConfig()}

	tmp.Analysis(analysis)

	tmp.Source(source)

	return tmp

}

func (s *_dataframePreviewConfig) Analysis(analysis types.DataframeAnalysisContainerVariant) *_dataframePreviewConfig {

	s.v.Analysis = *analysis.DataframeAnalysisContainerCaster()

	return s
}

func (s *_dataframePreviewConfig) AnalyzedFields(analyzedfields types.DataframeAnalysisAnalyzedFieldsVariant) *_dataframePreviewConfig {

	s.v.AnalyzedFields = analyzedfields.DataframeAnalysisAnalyzedFieldsCaster()

	return s
}

func (s *_dataframePreviewConfig) MaxNumThreads(maxnumthreads int) *_dataframePreviewConfig {

	s.v.MaxNumThreads = &maxnumthreads

	return s
}

func (s *_dataframePreviewConfig) ModelMemoryLimit(modelmemorylimit string) *_dataframePreviewConfig {

	s.v.ModelMemoryLimit = &modelmemorylimit

	return s
}

func (s *_dataframePreviewConfig) Source(source types.DataframeAnalyticsSourceVariant) *_dataframePreviewConfig {

	s.v.Source = *source.DataframeAnalyticsSourceCaster()

	return s
}

func (s *_dataframePreviewConfig) DataframePreviewConfigCaster() *types.DataframePreviewConfig {
	return s.v
}
