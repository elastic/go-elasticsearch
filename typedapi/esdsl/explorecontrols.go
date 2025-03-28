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

type _exploreControls struct {
	v *types.ExploreControls
}

func NewExploreControls(usesignificance bool) *_exploreControls {

	tmp := &_exploreControls{v: types.NewExploreControls()}

	tmp.UseSignificance(usesignificance)

	return tmp

}

// To avoid the top-matching documents sample being dominated by a single source
// of results, it is sometimes necessary to request diversity in the sample.
// You can do this by selecting a single-value field and setting a maximum
// number of documents per value for that field.
func (s *_exploreControls) SampleDiversity(samplediversity types.SampleDiversityVariant) *_exploreControls {

	s.v.SampleDiversity = samplediversity.SampleDiversityCaster()

	return s
}

// Each hop considers a sample of the best-matching documents on each shard.
// Using samples improves the speed of execution and keeps exploration focused
// on meaningfully-connected terms.
// Very small values (less than 50) might not provide sufficient
// weight-of-evidence to identify significant connections between terms.
// Very large sample sizes can dilute the quality of the results and increase
// execution times.
func (s *_exploreControls) SampleSize(samplesize int) *_exploreControls {

	s.v.SampleSize = &samplesize

	return s
}

// The length of time in milliseconds after which exploration will be halted and
// the results gathered so far are returned.
// This timeout is honored on a best-effort basis.
// Execution might overrun this timeout if, for example, a long pause is
// encountered while FieldData is loaded for a field.
func (s *_exploreControls) Timeout(duration types.DurationVariant) *_exploreControls {

	s.v.Timeout = *duration.DurationCaster()

	return s
}

// Filters associated terms so only those that are significantly associated with
// your query are included.
func (s *_exploreControls) UseSignificance(usesignificance bool) *_exploreControls {

	s.v.UseSignificance = usesignificance

	return s
}

func (s *_exploreControls) ExploreControlsCaster() *types.ExploreControls {
	return s.v
}
