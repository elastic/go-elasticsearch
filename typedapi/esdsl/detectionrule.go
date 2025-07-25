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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ruleaction"
)

type _detectionRule struct {
	v *types.DetectionRule
}

func NewDetectionRule() *_detectionRule {

	return &_detectionRule{v: types.NewDetectionRule()}

}

func (s *_detectionRule) Actions(actions ...ruleaction.RuleAction) *_detectionRule {

	for _, v := range actions {

		s.v.Actions = append(s.v.Actions, v)

	}
	return s
}

func (s *_detectionRule) Conditions(conditions ...types.RuleConditionVariant) *_detectionRule {

	for _, v := range conditions {

		s.v.Conditions = append(s.v.Conditions, *v.RuleConditionCaster())

	}
	return s
}

func (s *_detectionRule) Scope(scope map[string]types.FilterRef) *_detectionRule {

	s.v.Scope = scope
	return s
}

func (s *_detectionRule) AddScope(key string, value types.FilterRefVariant) *_detectionRule {

	var tmp map[string]types.FilterRef
	if s.v.Scope == nil {
		s.v.Scope = make(map[string]types.FilterRef)
	} else {
		tmp = s.v.Scope
	}

	tmp[key] = *value.FilterRefCaster()

	s.v.Scope = tmp
	return s
}

func (s *_detectionRule) DetectionRuleCaster() *types.DetectionRule {
	return s.v
}
