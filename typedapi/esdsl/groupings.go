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

type _groupings struct {
	v *types.Groupings
}

func NewGroupings() *_groupings {

	return &_groupings{v: types.NewGroupings()}

}

// A date histogram group aggregates a date field into time-based buckets.
// This group is mandatory; you currently cannot roll up documents without a
// timestamp and a `date_histogram` group.
func (s *_groupings) DateHistogram(datehistogram types.DateHistogramGroupingVariant) *_groupings {

	s.v.DateHistogram = datehistogram.DateHistogramGroupingCaster()

	return s
}

// The histogram group aggregates one or more numeric fields into numeric
// histogram intervals.
func (s *_groupings) Histogram(histogram types.HistogramGroupingVariant) *_groupings {

	s.v.Histogram = histogram.HistogramGroupingCaster()

	return s
}

// The terms group can be used on keyword or numeric fields to allow bucketing
// via the terms aggregation at a later point.
// The indexer enumerates and stores all values of a field for each time-period.
// This can be potentially costly for high-cardinality groups such as IP
// addresses, especially if the time-bucket is particularly sparse.
func (s *_groupings) Terms(terms types.TermsGroupingVariant) *_groupings {

	s.v.Terms = terms.TermsGroupingCaster()

	return s
}

func (s *_groupings) GroupingsCaster() *types.Groupings {
	return s.v
}
