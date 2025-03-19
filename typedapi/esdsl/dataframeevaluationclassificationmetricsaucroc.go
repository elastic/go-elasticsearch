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

type _dataframeEvaluationClassificationMetricsAucRoc struct {
	v *types.DataframeEvaluationClassificationMetricsAucRoc
}

func NewDataframeEvaluationClassificationMetricsAucRoc() *_dataframeEvaluationClassificationMetricsAucRoc {

	return &_dataframeEvaluationClassificationMetricsAucRoc{v: types.NewDataframeEvaluationClassificationMetricsAucRoc()}

}

// Name of the only class that is treated as positive during AUC ROC
// calculation. Other classes are treated as negative ("one-vs-all" strategy).
// All the evaluated documents must have class_name in the list of their top
// classes.
func (s *_dataframeEvaluationClassificationMetricsAucRoc) ClassName(name string) *_dataframeEvaluationClassificationMetricsAucRoc {

	s.v.ClassName = &name

	return s
}

// Whether or not the curve should be returned in addition to the score. Default
// value is false.
func (s *_dataframeEvaluationClassificationMetricsAucRoc) IncludeCurve(includecurve bool) *_dataframeEvaluationClassificationMetricsAucRoc {

	s.v.IncludeCurve = &includecurve

	return s
}

func (s *_dataframeEvaluationClassificationMetricsAucRoc) DataframeEvaluationClassificationMetricsAucRocCaster() *types.DataframeEvaluationClassificationMetricsAucRoc {
	return s.v
}
