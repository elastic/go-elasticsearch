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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _chiSquareHeuristic struct {
	v *types.ChiSquareHeuristic
}

func NewChiSquareHeuristic(backgroundissuperset bool, includenegatives bool) *_chiSquareHeuristic {

	tmp := &_chiSquareHeuristic{v: types.NewChiSquareHeuristic()}

	tmp.BackgroundIsSuperset(backgroundissuperset)

	tmp.IncludeNegatives(includenegatives)

	return tmp

}

// Set to `false` if you defined a custom background filter that represents a
// different set of documents that you want to compare to.
func (s *_chiSquareHeuristic) BackgroundIsSuperset(backgroundissuperset bool) *_chiSquareHeuristic {

	s.v.BackgroundIsSuperset = backgroundissuperset

	return s
}

// Set to `false` to filter out the terms that appear less often in the subset
// than in documents outside the subset.
func (s *_chiSquareHeuristic) IncludeNegatives(includenegatives bool) *_chiSquareHeuristic {

	s.v.IncludeNegatives = includenegatives

	return s
}

func (s *_chiSquareHeuristic) ChiSquareHeuristicCaster() *types.ChiSquareHeuristic {
	return s.v
}
