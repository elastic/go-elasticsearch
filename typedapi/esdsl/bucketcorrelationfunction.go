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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _bucketCorrelationFunction struct {
	v *types.BucketCorrelationFunction
}

func NewBucketCorrelationFunction(countcorrelation types.BucketCorrelationFunctionCountCorrelationVariant) *_bucketCorrelationFunction {

	tmp := &_bucketCorrelationFunction{v: types.NewBucketCorrelationFunction()}

	tmp.CountCorrelation(countcorrelation)

	return tmp

}

// The configuration to calculate a count correlation. This function is designed
// for determining the correlation of a term value and a given metric.
func (s *_bucketCorrelationFunction) CountCorrelation(countcorrelation types.BucketCorrelationFunctionCountCorrelationVariant) *_bucketCorrelationFunction {

	s.v.CountCorrelation = *countcorrelation.BucketCorrelationFunctionCountCorrelationCaster()

	return s
}

func (s *_bucketCorrelationFunction) BucketCorrelationFunctionCaster() *types.BucketCorrelationFunction {
	return s.v
}
