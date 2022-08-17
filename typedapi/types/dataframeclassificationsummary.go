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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// DataframeClassificationSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L31-L37
type DataframeClassificationSummary struct {
	Accuracy                  *DataframeClassificationSummaryAccuracy                  `json:"accuracy,omitempty"`
	AucRoc                    *DataframeEvaluationSummaryAucRoc                        `json:"auc_roc,omitempty"`
	MulticlassConfusionMatrix *DataframeClassificationSummaryMulticlassConfusionMatrix `json:"multiclass_confusion_matrix,omitempty"`
	Precision                 *DataframeClassificationSummaryPrecision                 `json:"precision,omitempty"`
	Recall                    *DataframeClassificationSummaryRecall                    `json:"recall,omitempty"`
}

// DataframeClassificationSummaryBuilder holds DataframeClassificationSummary struct and provides a builder API.
type DataframeClassificationSummaryBuilder struct {
	v *DataframeClassificationSummary
}

// NewDataframeClassificationSummary provides a builder for the DataframeClassificationSummary struct.
func NewDataframeClassificationSummaryBuilder() *DataframeClassificationSummaryBuilder {
	r := DataframeClassificationSummaryBuilder{
		&DataframeClassificationSummary{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeClassificationSummary struct
func (rb *DataframeClassificationSummaryBuilder) Build() DataframeClassificationSummary {
	return *rb.v
}

func (rb *DataframeClassificationSummaryBuilder) Accuracy(accuracy *DataframeClassificationSummaryAccuracyBuilder) *DataframeClassificationSummaryBuilder {
	v := accuracy.Build()
	rb.v.Accuracy = &v
	return rb
}

func (rb *DataframeClassificationSummaryBuilder) AucRoc(aucroc *DataframeEvaluationSummaryAucRocBuilder) *DataframeClassificationSummaryBuilder {
	v := aucroc.Build()
	rb.v.AucRoc = &v
	return rb
}

func (rb *DataframeClassificationSummaryBuilder) MulticlassConfusionMatrix(multiclassconfusionmatrix *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder) *DataframeClassificationSummaryBuilder {
	v := multiclassconfusionmatrix.Build()
	rb.v.MulticlassConfusionMatrix = &v
	return rb
}

func (rb *DataframeClassificationSummaryBuilder) Precision(precision *DataframeClassificationSummaryPrecisionBuilder) *DataframeClassificationSummaryBuilder {
	v := precision.Build()
	rb.v.Precision = &v
	return rb
}

func (rb *DataframeClassificationSummaryBuilder) Recall(recall *DataframeClassificationSummaryRecallBuilder) *DataframeClassificationSummaryBuilder {
	v := recall.Build()
	rb.v.Recall = &v
	return rb
}
