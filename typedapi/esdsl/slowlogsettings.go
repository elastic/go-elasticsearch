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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slowlogSettings struct {
	v *types.SlowlogSettings
}

func NewSlowlogSettings() *_slowlogSettings {

	return &_slowlogSettings{v: types.NewSlowlogSettings()}

}

func (s *_slowlogSettings) Level(level string) *_slowlogSettings {

	s.v.Level = &level

	return s
}

func (s *_slowlogSettings) Reformat(reformat bool) *_slowlogSettings {

	s.v.Reformat = &reformat

	return s
}

func (s *_slowlogSettings) Source(source int) *_slowlogSettings {

	s.v.Source = &source

	return s
}

func (s *_slowlogSettings) Threshold(threshold types.SlowlogTresholdsVariant) *_slowlogSettings {

	s.v.Threshold = threshold.SlowlogTresholdsCaster()

	return s
}

func (s *_slowlogSettings) SlowlogSettingsCaster() *types.SlowlogSettings {
	return s.v
}
