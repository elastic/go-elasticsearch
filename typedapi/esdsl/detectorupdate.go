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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _detectorUpdate struct {
	v *types.DetectorUpdate
}

func NewDetectorUpdate(detectorindex int) *_detectorUpdate {

	tmp := &_detectorUpdate{v: types.NewDetectorUpdate()}

	tmp.DetectorIndex(detectorindex)

	return tmp

}

// An array of custom rule objects, which enable you to customize the way
// detectors operate.
// For example, a rule may dictate to the detector conditions under which
// results should be skipped.
// Kibana refers to custom rules as job rules.
func (s *_detectorUpdate) CustomRules(customrules ...types.DetectionRuleVariant) *_detectorUpdate {

	for _, v := range customrules {

		s.v.CustomRules = append(s.v.CustomRules, *v.DetectionRuleCaster())

	}
	return s
}

// A description of the detector.
func (s *_detectorUpdate) Description(description string) *_detectorUpdate {

	s.v.Description = &description

	return s
}

// A unique identifier for the detector.
// This identifier is based on the order of the detectors in the
// `analysis_config`, starting at zero.
func (s *_detectorUpdate) DetectorIndex(detectorindex int) *_detectorUpdate {

	s.v.DetectorIndex = detectorindex

	return s
}

func (s *_detectorUpdate) DetectorUpdateCaster() *types.DetectorUpdate {
	return s.v
}
