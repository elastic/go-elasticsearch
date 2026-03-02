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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeEvaluationClassificationMetricsAucRoc struct {
	v *types.DataframeEvaluationClassificationMetricsAucRoc
}

func NewDataframeEvaluationClassificationMetricsAucRoc() *_dataframeEvaluationClassificationMetricsAucRoc {

	return &_dataframeEvaluationClassificationMetricsAucRoc{v: types.NewDataframeEvaluationClassificationMetricsAucRoc()}

}

func (s *_dataframeEvaluationClassificationMetricsAucRoc) ClassName(name string) *_dataframeEvaluationClassificationMetricsAucRoc {

	s.v.ClassName = &name

	return s
}

func (s *_dataframeEvaluationClassificationMetricsAucRoc) IncludeCurve(includecurve bool) *_dataframeEvaluationClassificationMetricsAucRoc {

	s.v.IncludeCurve = &includecurve

	return s
}

func (s *_dataframeEvaluationClassificationMetricsAucRoc) DataframeEvaluationClassificationMetricsAucRocCaster() *types.DataframeEvaluationClassificationMetricsAucRoc {
	return s.v
}
