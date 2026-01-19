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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _mappingLimitSettingsTotalFields struct {
	v *types.MappingLimitSettingsTotalFields
}

func NewMappingLimitSettingsTotalFields() *_mappingLimitSettingsTotalFields {

	return &_mappingLimitSettingsTotalFields{v: types.NewMappingLimitSettingsTotalFields()}

}

func (s *_mappingLimitSettingsTotalFields) IgnoreDynamicBeyondLimit(ignoredynamicbeyondlimit string) *_mappingLimitSettingsTotalFields {

	s.v.IgnoreDynamicBeyondLimit = &ignoredynamicbeyondlimit

	return s
}

func (s *_mappingLimitSettingsTotalFields) Limit(limit string) *_mappingLimitSettingsTotalFields {

	s.v.Limit = &limit

	return s
}

func (s *_mappingLimitSettingsTotalFields) MappingLimitSettingsTotalFieldsCaster() *types.MappingLimitSettingsTotalFields {
	return s.v
}
