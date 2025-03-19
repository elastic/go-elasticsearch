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

type _bucketCorrelationFunctionCountCorrelationIndicator struct {
	v *types.BucketCorrelationFunctionCountCorrelationIndicator
}

func NewBucketCorrelationFunctionCountCorrelationIndicator(doccount int) *_bucketCorrelationFunctionCountCorrelationIndicator {

	tmp := &_bucketCorrelationFunctionCountCorrelationIndicator{v: types.NewBucketCorrelationFunctionCountCorrelationIndicator()}

	tmp.DocCount(doccount)

	return tmp

}

// The total number of documents that initially created the expectations. It’s
// required to be greater
// than or equal to the sum of all values in the buckets_path as this is the
// originating superset of data
// to which the term values are correlated.
func (s *_bucketCorrelationFunctionCountCorrelationIndicator) DocCount(doccount int) *_bucketCorrelationFunctionCountCorrelationIndicator {

	s.v.DocCount = doccount

	return s
}

// An array of numbers with which to correlate the configured `bucket_path`
// values.
// The length of this value must always equal the number of buckets returned by
// the `bucket_path`.
func (s *_bucketCorrelationFunctionCountCorrelationIndicator) Expectations(expectations ...types.Float64) *_bucketCorrelationFunctionCountCorrelationIndicator {

	for _, v := range expectations {

		s.v.Expectations = append(s.v.Expectations, v)

	}
	return s
}

// An array of fractions to use when averaging and calculating variance. This
// should be used if
// the pre-calculated data and the buckets_path have known gaps. The length of
// fractions, if provided,
// must equal expectations.
func (s *_bucketCorrelationFunctionCountCorrelationIndicator) Fractions(fractions ...types.Float64) *_bucketCorrelationFunctionCountCorrelationIndicator {

	for _, v := range fractions {

		s.v.Fractions = append(s.v.Fractions, v)

	}
	return s
}

func (s *_bucketCorrelationFunctionCountCorrelationIndicator) BucketCorrelationFunctionCountCorrelationIndicatorCaster() *types.BucketCorrelationFunctionCountCorrelationIndicator {
	return s.v
}
