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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _dataframeEvaluationOutlierDetectionMetrics struct {
	v *types.DataframeEvaluationOutlierDetectionMetrics
}

func NewDataframeEvaluationOutlierDetectionMetrics() *_dataframeEvaluationOutlierDetectionMetrics {

	return &_dataframeEvaluationOutlierDetectionMetrics{v: types.NewDataframeEvaluationOutlierDetectionMetrics()}

}

func (s *_dataframeEvaluationOutlierDetectionMetrics) AucRoc(aucroc types.DataframeEvaluationClassificationMetricsAucRocVariant) *_dataframeEvaluationOutlierDetectionMetrics {

	s.v.AucRoc = aucroc.DataframeEvaluationClassificationMetricsAucRocCaster()

	return s
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) ConfusionMatrix(confusionmatrix map[string]json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {

	s.v.ConfusionMatrix = confusionmatrix
	return s
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) AddConfusionMatrix(key string, value json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {

	var tmp map[string]json.RawMessage
	if s.v.ConfusionMatrix == nil {
		s.v.ConfusionMatrix = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.ConfusionMatrix
	}

	tmp[key] = value

	s.v.ConfusionMatrix = tmp
	return s
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) Precision(precision map[string]json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {

	s.v.Precision = precision
	return s
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) AddPrecision(key string, value json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {

	var tmp map[string]json.RawMessage
	if s.v.Precision == nil {
		s.v.Precision = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Precision
	}

	tmp[key] = value

	s.v.Precision = tmp
	return s
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) Recall(recall map[string]json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {

	s.v.Recall = recall
	return s
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) AddRecall(key string, value json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {

	var tmp map[string]json.RawMessage
	if s.v.Recall == nil {
		s.v.Recall = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Recall
	}

	tmp[key] = value

	s.v.Recall = tmp
	return s
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) DataframeEvaluationOutlierDetectionMetricsCaster() *types.DataframeEvaluationOutlierDetectionMetrics {
	return s.v
}
