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

// This is provide all the types that are part of the union.
type _decisionScoreLevel struct {
	v types.DecisionScoreLevel
}

func NewDecisionScoreLevel() *_decisionScoreLevel {
	return &_decisionScoreLevel{v: nil}
}

func (u *_decisionScoreLevel) String(string string) *_decisionScoreLevel {

	u.v = &string

	return u
}

func (u *_decisionScoreLevel) DecisionScoreLevelObject(decisionscorelevelobject types.DecisionScoreLevelObjectVariant) *_decisionScoreLevel {

	u.v = decisionscorelevelobject.DecisionScoreLevelObjectCaster()

	return u
}

// Interface implementation for DecisionScoreLevelObject in DecisionScoreLevel union
func (u *_decisionScoreLevelObject) DecisionScoreLevelCaster() *types.DecisionScoreLevel {
	t := types.DecisionScoreLevel(u.v)
	return &t
}

func (u *_decisionScoreLevel) DecisionScoreLevelCaster() *types.DecisionScoreLevel {
	return &u.v
}
