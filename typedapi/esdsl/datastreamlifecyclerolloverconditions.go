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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataStreamLifecycleRolloverConditions struct {
	v *types.DataStreamLifecycleRolloverConditions
}

func NewDataStreamLifecycleRolloverConditions() *_dataStreamLifecycleRolloverConditions {

	return &_dataStreamLifecycleRolloverConditions{v: types.NewDataStreamLifecycleRolloverConditions()}

}

func (s *_dataStreamLifecycleRolloverConditions) MaxAge(maxage string) *_dataStreamLifecycleRolloverConditions {

	s.v.MaxAge = &maxage

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MaxDocs(maxdocs int64) *_dataStreamLifecycleRolloverConditions {

	s.v.MaxDocs = &maxdocs

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MaxPrimaryShardDocs(maxprimarysharddocs int64) *_dataStreamLifecycleRolloverConditions {

	s.v.MaxPrimaryShardDocs = &maxprimarysharddocs

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MaxPrimaryShardSize(bytesize types.ByteSizeVariant) *_dataStreamLifecycleRolloverConditions {

	s.v.MaxPrimaryShardSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MaxSize(bytesize types.ByteSizeVariant) *_dataStreamLifecycleRolloverConditions {

	s.v.MaxSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MinAge(duration types.DurationVariant) *_dataStreamLifecycleRolloverConditions {

	s.v.MinAge = *duration.DurationCaster()

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MinDocs(mindocs int64) *_dataStreamLifecycleRolloverConditions {

	s.v.MinDocs = &mindocs

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MinPrimaryShardDocs(minprimarysharddocs int64) *_dataStreamLifecycleRolloverConditions {

	s.v.MinPrimaryShardDocs = &minprimarysharddocs

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MinPrimaryShardSize(bytesize types.ByteSizeVariant) *_dataStreamLifecycleRolloverConditions {

	s.v.MinPrimaryShardSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) MinSize(bytesize types.ByteSizeVariant) *_dataStreamLifecycleRolloverConditions {

	s.v.MinSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_dataStreamLifecycleRolloverConditions) DataStreamLifecycleRolloverConditionsCaster() *types.DataStreamLifecycleRolloverConditions {
	return s.v
}
