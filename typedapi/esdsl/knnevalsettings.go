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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnEvalSettings struct {
	v *types.KnnEvalSettings
}

func NewKnnEvalSettings() *_knnEvalSettings {

	return &_knnEvalSettings{v: types.NewKnnEvalSettings()}

}

func (s *_knnEvalSettings) Exact(exact bool) *_knnEvalSettings {

	s.v.Exact = &exact

	return s
}

func (s *_knnEvalSettings) NumCandidates(numcandidates int) *_knnEvalSettings {

	s.v.NumCandidates = &numcandidates

	return s
}

func (s *_knnEvalSettings) RescoreVector(rescorevector types.RescoreVectorVariant) *_knnEvalSettings {

	s.v.RescoreVector = rescorevector.RescoreVectorCaster()

	return s
}

func (s *_knnEvalSettings) VisitPercentage(visitpercentage float32) *_knnEvalSettings {

	s.v.VisitPercentage = &visitpercentage

	return s
}

func (s *_knnEvalSettings) KnnEvalSettingsCaster() *types.KnnEvalSettings {
	return s.v
}
