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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rolloverConditions struct {
	v *types.RolloverConditions
}

func NewRolloverConditions() *_rolloverConditions {

	return &_rolloverConditions{v: types.NewRolloverConditions()}

}

func (s *_rolloverConditions) MaxAge(duration types.DurationVariant) *_rolloverConditions {

	s.v.MaxAge = *duration.DurationCaster()

	return s
}

func (s *_rolloverConditions) MaxAgeMillis(durationvalueunitmillis int64) *_rolloverConditions {

	s.v.MaxAgeMillis = &durationvalueunitmillis

	return s
}

func (s *_rolloverConditions) MaxDocs(maxdocs int64) *_rolloverConditions {

	s.v.MaxDocs = &maxdocs

	return s
}

func (s *_rolloverConditions) MaxPrimaryShardDocs(maxprimarysharddocs int64) *_rolloverConditions {

	s.v.MaxPrimaryShardDocs = &maxprimarysharddocs

	return s
}

func (s *_rolloverConditions) MaxPrimaryShardSize(bytesize types.ByteSizeVariant) *_rolloverConditions {

	s.v.MaxPrimaryShardSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_rolloverConditions) MaxPrimaryShardSizeBytes(maxprimaryshardsizebytes int64) *_rolloverConditions {

	s.v.MaxPrimaryShardSizeBytes = &maxprimaryshardsizebytes

	return s
}

func (s *_rolloverConditions) MaxSize(bytesize types.ByteSizeVariant) *_rolloverConditions {

	s.v.MaxSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_rolloverConditions) MaxSizeBytes(maxsizebytes int64) *_rolloverConditions {

	s.v.MaxSizeBytes = &maxsizebytes

	return s
}

func (s *_rolloverConditions) MinAge(duration types.DurationVariant) *_rolloverConditions {

	s.v.MinAge = *duration.DurationCaster()

	return s
}

func (s *_rolloverConditions) MinDocs(mindocs int64) *_rolloverConditions {

	s.v.MinDocs = &mindocs

	return s
}

func (s *_rolloverConditions) MinPrimaryShardDocs(minprimarysharddocs int64) *_rolloverConditions {

	s.v.MinPrimaryShardDocs = &minprimarysharddocs

	return s
}

func (s *_rolloverConditions) MinPrimaryShardSize(bytesize types.ByteSizeVariant) *_rolloverConditions {

	s.v.MinPrimaryShardSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_rolloverConditions) MinPrimaryShardSizeBytes(minprimaryshardsizebytes int64) *_rolloverConditions {

	s.v.MinPrimaryShardSizeBytes = &minprimaryshardsizebytes

	return s
}

func (s *_rolloverConditions) MinSize(bytesize types.ByteSizeVariant) *_rolloverConditions {

	s.v.MinSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_rolloverConditions) MinSizeBytes(minsizebytes int64) *_rolloverConditions {

	s.v.MinSizeBytes = &minsizebytes

	return s
}

func (s *_rolloverConditions) RolloverConditionsCaster() *types.RolloverConditions {
	return s.v
}
