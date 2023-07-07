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
// https://github.com/elastic/elasticsearch-specification/tree/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c

package types

// DataframeOutlierDetectionSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c/specification/ml/evaluate_data_frame/types.ts#L24-L29
type DataframeOutlierDetectionSummary struct {
	AucRoc          *DataframeEvaluationSummaryAucRoc   `json:"auc_roc,omitempty"`
	ConfusionMatrix map[string]ConfusionMatrixThreshold `json:"confusion_matrix,omitempty"`
	Precision       map[string]Float64                  `json:"precision,omitempty"`
	Recall          map[string]Float64                  `json:"recall,omitempty"`
}

// NewDataframeOutlierDetectionSummary returns a DataframeOutlierDetectionSummary.
func NewDataframeOutlierDetectionSummary() *DataframeOutlierDetectionSummary {
	r := &DataframeOutlierDetectionSummary{
		ConfusionMatrix: make(map[string]ConfusionMatrixThreshold, 0),
		Precision:       make(map[string]Float64, 0),
		Recall:          make(map[string]Float64, 0),
	}

	return r
}
