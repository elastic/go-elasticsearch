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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/excludefrequent"
)

type _detector struct {
	v *types.Detector
}

func NewDetector() *_detector {

	return &_detector{v: types.NewDetector()}

}

func (s *_detector) ByFieldName(field string) *_detector {

	s.v.ByFieldName = &field

	return s
}

func (s *_detector) CustomRules(customrules ...types.DetectionRuleVariant) *_detector {

	for _, v := range customrules {

		s.v.CustomRules = append(s.v.CustomRules, *v.DetectionRuleCaster())

	}
	return s
}

func (s *_detector) DetectorDescription(detectordescription string) *_detector {

	s.v.DetectorDescription = &detectordescription

	return s
}

func (s *_detector) DetectorIndex(detectorindex int) *_detector {

	s.v.DetectorIndex = &detectorindex

	return s
}

func (s *_detector) ExcludeFrequent(excludefrequent excludefrequent.ExcludeFrequent) *_detector {

	s.v.ExcludeFrequent = &excludefrequent
	return s
}

func (s *_detector) FieldName(field string) *_detector {

	s.v.FieldName = &field

	return s
}

func (s *_detector) Function(function string) *_detector {

	s.v.Function = &function

	return s
}

func (s *_detector) OverFieldName(field string) *_detector {

	s.v.OverFieldName = &field

	return s
}

func (s *_detector) PartitionFieldName(field string) *_detector {

	s.v.PartitionFieldName = &field

	return s
}

func (s *_detector) UseNull(usenull bool) *_detector {

	s.v.UseNull = &usenull

	return s
}

func (s *_detector) DetectorCaster() *types.Detector {
	return s.v
}
