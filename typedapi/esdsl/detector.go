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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/excludefrequent"
)

type _detector struct {
	v *types.Detector
}

func NewDetector() *_detector {

	return &_detector{v: types.NewDetector()}

}

// The field used to split the data. In particular, this property is used for
// analyzing the splits with respect to their own history. It is used for
// finding unusual values in the context of the split.
func (s *_detector) ByFieldName(field string) *_detector {

	s.v.ByFieldName = &field

	return s
}

// Custom rules enable you to customize the way detectors operate. For example,
// a rule may dictate conditions under which results should be skipped. Kibana
// refers to custom rules as job rules.
func (s *_detector) CustomRules(customrules ...types.DetectionRuleVariant) *_detector {

	for _, v := range customrules {

		s.v.CustomRules = append(s.v.CustomRules, *v.DetectionRuleCaster())

	}
	return s
}

// A description of the detector.
func (s *_detector) DetectorDescription(detectordescription string) *_detector {

	s.v.DetectorDescription = &detectordescription

	return s
}

// A unique identifier for the detector. This identifier is based on the order
// of the detectors in the `analysis_config`, starting at zero. If you specify a
// value for this property, it is ignored.
func (s *_detector) DetectorIndex(detectorindex int) *_detector {

	s.v.DetectorIndex = &detectorindex

	return s
}

// If set, frequent entities are excluded from influencing the anomaly results.
// Entities can be considered frequent over time or frequent in a population. If
// you are working with both over and by fields, you can set `exclude_frequent`
// to `all` for both fields, or to `by` or `over` for those specific fields.
func (s *_detector) ExcludeFrequent(excludefrequent excludefrequent.ExcludeFrequent) *_detector {

	s.v.ExcludeFrequent = &excludefrequent
	return s
}

// The field that the detector uses in the function. If you use an event rate
// function such as count or rare, do not specify this field. The `field_name`
// cannot contain double quotes or backslashes.
func (s *_detector) FieldName(field string) *_detector {

	s.v.FieldName = &field

	return s
}

// The analysis function that is used. For example, `count`, `rare`, `mean`,
// `min`, `max`, or `sum`.
func (s *_detector) Function(function string) *_detector {

	s.v.Function = &function

	return s
}

// The field used to split the data. In particular, this property is used for
// analyzing the splits with respect to the history of all splits. It is used
// for finding unusual values in the population of all splits.
func (s *_detector) OverFieldName(field string) *_detector {

	s.v.OverFieldName = &field

	return s
}

// The field used to segment the analysis. When you use this property, you have
// completely independent baselines for each value of this field.
func (s *_detector) PartitionFieldName(field string) *_detector {

	s.v.PartitionFieldName = &field

	return s
}

// Defines whether a new series is used as the null series when there is no
// value for the by or partition fields.
func (s *_detector) UseNull(usenull bool) *_detector {

	s.v.UseNull = &usenull

	return s
}

func (s *_detector) DetectorCaster() *types.Detector {
	return s.v
}
