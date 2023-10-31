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

// DataframeOutlierDetectionSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/evaluate_data_frame/types.ts#L24-L42
type DataframeOutlierDetectionSummary struct {
	// AucRoc The AUC ROC (area under the curve of the receiver operating characteristic)
	// score and optionally the curve.
	AucRoc *DataframeEvaluationSummaryAucRoc `json:"auc_roc,omitempty"`
	// ConfusionMatrix Set the different thresholds of the outlier score at where the metrics (`tp`
	// - true positive, `fp` - false positive, `tn` - true negative, `fn` - false
	// negative) are calculated.
	ConfusionMatrix map[string]ConfusionMatrixThreshold `json:"confusion_matrix,omitempty"`
	// Precision Set the different thresholds of the outlier score at where the metric is
	// calculated.
	Precision map[string]Float64 `json:"precision,omitempty"`
	// Recall Set the different thresholds of the outlier score at where the metric is
	// calculated.
	Recall map[string]Float64 `json:"recall,omitempty"`
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
