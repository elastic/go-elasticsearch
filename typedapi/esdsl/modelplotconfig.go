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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _modelPlotConfig struct {
	v *types.ModelPlotConfig
}

func NewModelPlotConfig() *_modelPlotConfig {

	return &_modelPlotConfig{v: types.NewModelPlotConfig()}

}

// If true, enables calculation and storage of the model change annotations for
// each entity that is being analyzed.
func (s *_modelPlotConfig) AnnotationsEnabled(annotationsenabled bool) *_modelPlotConfig {

	s.v.AnnotationsEnabled = &annotationsenabled

	return s
}

// If true, enables calculation and storage of the model bounds for each entity
// that is being analyzed.
func (s *_modelPlotConfig) Enabled(enabled bool) *_modelPlotConfig {

	s.v.Enabled = &enabled

	return s
}

// Limits data collection to this comma separated list of partition or by field
// values. If terms are not specified or it is an empty string, no filtering is
// applied. Wildcards are not supported. Only the specified terms can be viewed
// when using the Single Metric Viewer.
func (s *_modelPlotConfig) Terms(field string) *_modelPlotConfig {

	s.v.Terms = &field

	return s
}

func (s *_modelPlotConfig) ModelPlotConfigCaster() *types.ModelPlotConfig {
	return s.v
}
