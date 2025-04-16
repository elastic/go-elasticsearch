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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorindexoptionstype"
)

type _denseVectorIndexOptions struct {
	v *types.DenseVectorIndexOptions
}

func NewDenseVectorIndexOptions(type_ densevectorindexoptionstype.DenseVectorIndexOptionsType) *_denseVectorIndexOptions {

	tmp := &_denseVectorIndexOptions{v: types.NewDenseVectorIndexOptions()}

	tmp.Type(type_)

	return tmp

}

func (s *_denseVectorIndexOptions) ConfidenceInterval(confidenceinterval float32) *_denseVectorIndexOptions {

	s.v.ConfidenceInterval = &confidenceinterval

	return s
}

func (s *_denseVectorIndexOptions) EfConstruction(efconstruction int) *_denseVectorIndexOptions {

	s.v.EfConstruction = &efconstruction

	return s
}

func (s *_denseVectorIndexOptions) M(m int) *_denseVectorIndexOptions {

	s.v.M = &m

	return s
}

func (s *_denseVectorIndexOptions) RescoreVector(rescorevector types.DenseVectorIndexOptionsRescoreVectorVariant) *_denseVectorIndexOptions {

	s.v.RescoreVector = rescorevector.DenseVectorIndexOptionsRescoreVectorCaster()

	return s
}

func (s *_denseVectorIndexOptions) Type(type_ densevectorindexoptionstype.DenseVectorIndexOptionsType) *_denseVectorIndexOptions {

	s.v.Type = type_
	return s
}

func (s *_denseVectorIndexOptions) DenseVectorIndexOptionsCaster() *types.DenseVectorIndexOptions {
	return s.v
}
