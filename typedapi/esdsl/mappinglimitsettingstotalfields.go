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

type _mappingLimitSettingsTotalFields struct {
	v *types.MappingLimitSettingsTotalFields
}

func NewMappingLimitSettingsTotalFields() *_mappingLimitSettingsTotalFields {

	return &_mappingLimitSettingsTotalFields{v: types.NewMappingLimitSettingsTotalFields()}

}

// This setting determines what happens when a dynamically mapped field would
// exceed the total fields limit. When set
// to false (the default), the index request of the document that tries to add a
// dynamic field to the mapping will fail
// with the message Limit of total fields [X] has been exceeded. When set to
// true, the index request will not fail.
// Instead, fields that would exceed the limit are not added to the mapping,
// similar to dynamic: false.
// The fields that were not added to the mapping will be added to the _ignored
// field.
func (s *_mappingLimitSettingsTotalFields) IgnoreDynamicBeyondLimit(ignoredynamicbeyondlimit string) *_mappingLimitSettingsTotalFields {

	s.v.IgnoreDynamicBeyondLimit = ignoredynamicbeyondlimit

	return s
}

// The maximum number of fields in an index. Field and object mappings, as well
// as field aliases count towards this limit.
// The limit is in place to prevent mappings and searches from becoming too
// large. Higher values can lead to performance
// degradations and memory issues, especially in clusters with a high load or
// few resources.
func (s *_mappingLimitSettingsTotalFields) Limit(limit string) *_mappingLimitSettingsTotalFields {

	s.v.Limit = limit

	return s
}

func (s *_mappingLimitSettingsTotalFields) MappingLimitSettingsTotalFieldsCaster() *types.MappingLimitSettingsTotalFields {
	return s.v
}
