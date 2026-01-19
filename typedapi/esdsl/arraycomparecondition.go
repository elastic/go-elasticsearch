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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conditionop"
)

type _arrayCompareCondition struct {
	k string
	v *types.ArrayCompareCondition
}

func NewArrayCompareCondition(key string) *_arrayCompareCondition {
	return &_arrayCompareCondition{
		k: key,
		v: types.NewArrayCompareCondition(),
	}
}

func (s *_arrayCompareCondition) ArrayCompareCondition(arraycomparecondition map[conditionop.ConditionOp]types.ArrayCompareOpParams) *_arrayCompareCondition {

	s.v.ArrayCompareCondition = arraycomparecondition
	return s
}

func (s *_arrayCompareCondition) AddArrayCompareCondition(key conditionop.ConditionOp, value types.ArrayCompareOpParamsVariant) *_arrayCompareCondition {

	var tmp map[conditionop.ConditionOp]types.ArrayCompareOpParams
	if s.v.ArrayCompareCondition == nil {
		s.v.ArrayCompareCondition = make(map[conditionop.ConditionOp]types.ArrayCompareOpParams)
	} else {
		tmp = s.v.ArrayCompareCondition
	}

	tmp[key] = *value.ArrayCompareOpParamsCaster()

	s.v.ArrayCompareCondition = tmp
	return s
}

func (s *_arrayCompareCondition) Path(path string) *_arrayCompareCondition {

	s.v.Path = path

	return s
}

func (s *_arrayCompareCondition) WatcherConditionCaster() *types.WatcherCondition {
	container := types.NewWatcherCondition()
	container.ArrayCompare = map[string]types.ArrayCompareCondition{
		s.k: *s.v,
	}
	return container
}

// NewSingleArrayCompareCondition should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleArrayCompareCondition() *_arrayCompareCondition {
	return &_arrayCompareCondition{
		k: "",
		v: types.NewArrayCompareCondition(),
	}
}

func (s *_arrayCompareCondition) ArrayCompareConditionCaster() *types.ArrayCompareCondition {
	return s.v.ArrayCompareConditionCaster()
}
