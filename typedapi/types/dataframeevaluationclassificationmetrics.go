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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"encoding/json"
)

// DataframeEvaluationClassificationMetrics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeEvaluation.ts#L73-L78
type DataframeEvaluationClassificationMetrics struct {
	// Accuracy Accuracy of predictions (per-class and overall).
	Accuracy map[string]json.RawMessage `json:"accuracy,omitempty"`
	// AucRoc The AUC ROC (area under the curve of the receiver operating characteristic)
	// score and optionally the curve. It is calculated for a specific class
	// (provided as "class_name") treated as positive.
	AucRoc *DataframeEvaluationClassificationMetricsAucRoc `json:"auc_roc,omitempty"`
	// MulticlassConfusionMatrix Multiclass confusion matrix.
	MulticlassConfusionMatrix map[string]json.RawMessage `json:"multiclass_confusion_matrix,omitempty"`
	// Precision Precision of predictions (per-class and average).
	Precision map[string]json.RawMessage `json:"precision,omitempty"`
	// Recall Recall of predictions (per-class and average).
	Recall map[string]json.RawMessage `json:"recall,omitempty"`
}

// NewDataframeEvaluationClassificationMetrics returns a DataframeEvaluationClassificationMetrics.
func NewDataframeEvaluationClassificationMetrics() *DataframeEvaluationClassificationMetrics {
	r := &DataframeEvaluationClassificationMetrics{
		Accuracy:                  make(map[string]json.RawMessage, 0),
		MulticlassConfusionMatrix: make(map[string]json.RawMessage, 0),
		Precision:                 make(map[string]json.RawMessage, 0),
		Recall:                    make(map[string]json.RawMessage, 0),
	}

	return r
}
