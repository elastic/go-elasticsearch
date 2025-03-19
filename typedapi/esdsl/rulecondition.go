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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/appliesto"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conditionoperator"
)

type _ruleCondition struct {
	v *types.RuleCondition
}

func NewRuleCondition(appliesto appliesto.AppliesTo, operator conditionoperator.ConditionOperator, value types.Float64) *_ruleCondition {

	tmp := &_ruleCondition{v: types.NewRuleCondition()}

	tmp.AppliesTo(appliesto)

	tmp.Operator(operator)

	tmp.Value(value)

	return tmp

}

// Specifies the result property to which the condition applies. If your
// detector uses `lat_long`, `metric`, `rare`, or `freq_rare` functions, you can
// only specify conditions that apply to time.
func (s *_ruleCondition) AppliesTo(appliesto appliesto.AppliesTo) *_ruleCondition {

	s.v.AppliesTo = appliesto
	return s
}

// Specifies the condition operator. The available options are greater than,
// greater than or equals, less than, and less than or equals.
func (s *_ruleCondition) Operator(operator conditionoperator.ConditionOperator) *_ruleCondition {

	s.v.Operator = operator
	return s
}

// The value that is compared against the `applies_to` field using the operator.
func (s *_ruleCondition) Value(value types.Float64) *_ruleCondition {

	s.v.Value = value

	return s
}

func (s *_ruleCondition) RuleConditionCaster() *types.RuleCondition {
	return s.v
}
