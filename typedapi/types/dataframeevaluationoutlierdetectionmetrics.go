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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// DataframeEvaluationOutlierDetectionMetrics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/DataframeEvaluation.ts#L80-L83
type DataframeEvaluationOutlierDetectionMetrics struct {
	// AucRoc The AUC ROC (area under the curve of the receiver operating characteristic)
	// score and optionally the curve. It is calculated for a specific class
	// (provided as "class_name") treated as positive.
	AucRoc *DataframeEvaluationClassificationMetricsAucRoc `json:"auc_roc,omitempty"`
	// ConfusionMatrix Accuracy of predictions (per-class and overall).
	ConfusionMatrix map[string]interface{} `json:"confusion_matrix,omitempty"`
	// Precision Precision of predictions (per-class and average).
	Precision map[string]interface{} `json:"precision,omitempty"`
	// Recall Recall of predictions (per-class and average).
	Recall map[string]interface{} `json:"recall,omitempty"`
}

// NewDataframeEvaluationOutlierDetectionMetrics returns a DataframeEvaluationOutlierDetectionMetrics.
func NewDataframeEvaluationOutlierDetectionMetrics() *DataframeEvaluationOutlierDetectionMetrics {
	r := &DataframeEvaluationOutlierDetectionMetrics{
		ConfusionMatrix: make(map[string]interface{}, 0),
		Precision:       make(map[string]interface{}, 0),
		Recall:          make(map[string]interface{}, 0),
	}

	return r
}
