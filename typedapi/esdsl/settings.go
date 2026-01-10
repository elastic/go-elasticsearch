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

type _settings struct {
	v *types.Settings
}

func NewSettings() *_settings {

	return &_settings{v: types.NewSettings()}

}

func (s *_settings) AlignCheckpoints(aligncheckpoints bool) *_settings {

	s.v.AlignCheckpoints = &aligncheckpoints

	return s
}

func (s *_settings) DatesAsEpochMillis(datesasepochmillis bool) *_settings {

	s.v.DatesAsEpochMillis = &datesasepochmillis

	return s
}

func (s *_settings) DeduceMappings(deducemappings bool) *_settings {

	s.v.DeduceMappings = &deducemappings

	return s
}

func (s *_settings) DocsPerSecond(docspersecond float32) *_settings {

	s.v.DocsPerSecond = &docspersecond

	return s
}

func (s *_settings) MaxPageSearchSize(maxpagesearchsize int) *_settings {

	s.v.MaxPageSearchSize = &maxpagesearchsize

	return s
}

func (s *_settings) Unattended(unattended bool) *_settings {

	s.v.Unattended = &unattended

	return s
}

func (s *_settings) UsePointInTime(usepointintime bool) *_settings {

	s.v.UsePointInTime = &usepointintime

	return s
}

func (s *_settings) SettingsCaster() *types.Settings {
	return s.v
}
