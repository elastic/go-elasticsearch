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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

// DataframeClassificationSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/evaluate_data_frame/types.ts#L44-L66
type DataframeClassificationSummary struct {
	// Accuracy Accuracy of predictions (per-class and overall).
	Accuracy *DataframeClassificationSummaryAccuracy `json:"accuracy,omitempty"`
	// AucRoc The AUC ROC (area under the curve of the receiver operating characteristic)
	// score and optionally the curve.
	// It is calculated for a specific class (provided as "class_name") treated as
	// positive.
	AucRoc *DataframeEvaluationSummaryAucRoc `json:"auc_roc,omitempty"`
	// MulticlassConfusionMatrix Multiclass confusion matrix.
	MulticlassConfusionMatrix *DataframeClassificationSummaryMulticlassConfusionMatrix `json:"multiclass_confusion_matrix,omitempty"`
	// Precision Precision of predictions (per-class and average).
	Precision *DataframeClassificationSummaryPrecision `json:"precision,omitempty"`
	// Recall Recall of predictions (per-class and average).
	Recall *DataframeClassificationSummaryRecall `json:"recall,omitempty"`
}

// NewDataframeClassificationSummary returns a DataframeClassificationSummary.
func NewDataframeClassificationSummary() *DataframeClassificationSummary {
	r := &DataframeClassificationSummary{}

	return r
}
