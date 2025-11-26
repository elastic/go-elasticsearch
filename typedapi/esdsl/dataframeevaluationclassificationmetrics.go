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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _dataframeEvaluationClassificationMetrics struct {
	v *types.DataframeEvaluationClassificationMetrics
}

func NewDataframeEvaluationClassificationMetrics() *_dataframeEvaluationClassificationMetrics {

	return &_dataframeEvaluationClassificationMetrics{v: types.NewDataframeEvaluationClassificationMetrics()}

}

func (s *_dataframeEvaluationClassificationMetrics) Accuracy(accuracy map[string]json.RawMessage) *_dataframeEvaluationClassificationMetrics {

	s.v.Accuracy = accuracy
	return s
}

func (s *_dataframeEvaluationClassificationMetrics) AddAccuracy(key string, value json.RawMessage) *_dataframeEvaluationClassificationMetrics {

	var tmp map[string]json.RawMessage
	if s.v.Accuracy == nil {
		s.v.Accuracy = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Accuracy
	}

	tmp[key] = value

	s.v.Accuracy = tmp
	return s
}

func (s *_dataframeEvaluationClassificationMetrics) AucRoc(aucroc types.DataframeEvaluationClassificationMetricsAucRocVariant) *_dataframeEvaluationClassificationMetrics {

	s.v.AucRoc = aucroc.DataframeEvaluationClassificationMetricsAucRocCaster()

	return s
}

func (s *_dataframeEvaluationClassificationMetrics) MulticlassConfusionMatrix(multiclassconfusionmatrix map[string]json.RawMessage) *_dataframeEvaluationClassificationMetrics {

	s.v.MulticlassConfusionMatrix = multiclassconfusionmatrix
	return s
}

func (s *_dataframeEvaluationClassificationMetrics) AddMulticlassConfusionMatrix(key string, value json.RawMessage) *_dataframeEvaluationClassificationMetrics {

	var tmp map[string]json.RawMessage
	if s.v.MulticlassConfusionMatrix == nil {
		s.v.MulticlassConfusionMatrix = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.MulticlassConfusionMatrix
	}

	tmp[key] = value

	s.v.MulticlassConfusionMatrix = tmp
	return s
}

func (s *_dataframeEvaluationClassificationMetrics) Precision(precision map[string]json.RawMessage) *_dataframeEvaluationClassificationMetrics {

	s.v.Precision = precision
	return s
}

func (s *_dataframeEvaluationClassificationMetrics) AddPrecision(key string, value json.RawMessage) *_dataframeEvaluationClassificationMetrics {

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

func (s *_dataframeEvaluationClassificationMetrics) Recall(recall map[string]json.RawMessage) *_dataframeEvaluationClassificationMetrics {

	s.v.Recall = recall
	return s
}

func (s *_dataframeEvaluationClassificationMetrics) AddRecall(key string, value json.RawMessage) *_dataframeEvaluationClassificationMetrics {

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

func (s *_dataframeEvaluationClassificationMetrics) DataframeEvaluationClassificationMetricsCaster() *types.DataframeEvaluationClassificationMetrics {
	return s.v
}
