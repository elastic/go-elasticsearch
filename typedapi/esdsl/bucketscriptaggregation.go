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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _bucketScriptAggregation struct {
	v *types.BucketScriptAggregation
}

// A parent pipeline aggregation which runs a script which can perform per
// bucket computations on metrics in the parent multi-bucket aggregation.
func NewBucketScriptAggregation() *_bucketScriptAggregation {

	return &_bucketScriptAggregation{v: types.NewBucketScriptAggregation()}

}

func (s *_bucketScriptAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketScriptAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_bucketScriptAggregation) Format(format string) *_bucketScriptAggregation {

	s.v.Format = &format

	return s
}

func (s *_bucketScriptAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_bucketScriptAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_bucketScriptAggregation) Script(script types.ScriptVariant) *_bucketScriptAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_bucketScriptAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.BucketScript = s.v

	return container
}

func (s *_bucketScriptAggregation) BucketScriptAggregationCaster() *types.BucketScriptAggregation {
	return s.v
}
