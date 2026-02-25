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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bucketCorrelationFunctionCountCorrelationIndicator struct {
	v *types.BucketCorrelationFunctionCountCorrelationIndicator
}

func NewBucketCorrelationFunctionCountCorrelationIndicator(doccount int) *_bucketCorrelationFunctionCountCorrelationIndicator {

	tmp := &_bucketCorrelationFunctionCountCorrelationIndicator{v: types.NewBucketCorrelationFunctionCountCorrelationIndicator()}

	tmp.DocCount(doccount)

	return tmp

}

func (s *_bucketCorrelationFunctionCountCorrelationIndicator) DocCount(doccount int) *_bucketCorrelationFunctionCountCorrelationIndicator {

	s.v.DocCount = doccount

	return s
}

func (s *_bucketCorrelationFunctionCountCorrelationIndicator) Expectations(expectations ...types.Float64) *_bucketCorrelationFunctionCountCorrelationIndicator {

	for _, v := range expectations {

		s.v.Expectations = append(s.v.Expectations, v)

	}
	return s
}

func (s *_bucketCorrelationFunctionCountCorrelationIndicator) Fractions(fractions ...types.Float64) *_bucketCorrelationFunctionCountCorrelationIndicator {

	for _, v := range fractions {

		s.v.Fractions = append(s.v.Fractions, v)

	}
	return s
}

func (s *_bucketCorrelationFunctionCountCorrelationIndicator) BucketCorrelationFunctionCountCorrelationIndicatorCaster() *types.BucketCorrelationFunctionCountCorrelationIndicator {
	return s.v
}
