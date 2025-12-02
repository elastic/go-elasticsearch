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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rolloverAction struct {
	v *types.RolloverAction
}

func NewRolloverAction() *_rolloverAction {

	return &_rolloverAction{v: types.NewRolloverAction()}

}

func (s *_rolloverAction) MaxAge(duration types.DurationVariant) *_rolloverAction {

	s.v.MaxAge = *duration.DurationCaster()

	return s
}

func (s *_rolloverAction) MaxDocs(maxdocs int64) *_rolloverAction {

	s.v.MaxDocs = &maxdocs

	return s
}

func (s *_rolloverAction) MaxPrimaryShardDocs(maxprimarysharddocs int64) *_rolloverAction {

	s.v.MaxPrimaryShardDocs = &maxprimarysharddocs

	return s
}

func (s *_rolloverAction) MaxPrimaryShardSize(bytesize types.ByteSizeVariant) *_rolloverAction {

	s.v.MaxPrimaryShardSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_rolloverAction) MaxSize(bytesize types.ByteSizeVariant) *_rolloverAction {

	s.v.MaxSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_rolloverAction) MinAge(duration types.DurationVariant) *_rolloverAction {

	s.v.MinAge = *duration.DurationCaster()

	return s
}

func (s *_rolloverAction) MinDocs(mindocs int64) *_rolloverAction {

	s.v.MinDocs = &mindocs

	return s
}

func (s *_rolloverAction) MinPrimaryShardDocs(minprimarysharddocs int64) *_rolloverAction {

	s.v.MinPrimaryShardDocs = &minprimarysharddocs

	return s
}

func (s *_rolloverAction) MinPrimaryShardSize(bytesize types.ByteSizeVariant) *_rolloverAction {

	s.v.MinPrimaryShardSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_rolloverAction) MinSize(bytesize types.ByteSizeVariant) *_rolloverAction {

	s.v.MinSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_rolloverAction) RolloverActionCaster() *types.RolloverAction {
	return s.v
}
