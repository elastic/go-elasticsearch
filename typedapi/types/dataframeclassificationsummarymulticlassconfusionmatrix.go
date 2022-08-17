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

// DataframeClassificationSummaryMulticlassConfusionMatrix type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L79-L82
type DataframeClassificationSummaryMulticlassConfusionMatrix struct {
	ConfusionMatrix       []ConfusionMatrixItem `json:"confusion_matrix"`
	OtherActualClassCount int                   `json:"other_actual_class_count"`
}

// DataframeClassificationSummaryMulticlassConfusionMatrixBuilder holds DataframeClassificationSummaryMulticlassConfusionMatrix struct and provides a builder API.
type DataframeClassificationSummaryMulticlassConfusionMatrixBuilder struct {
	v *DataframeClassificationSummaryMulticlassConfusionMatrix
}

// NewDataframeClassificationSummaryMulticlassConfusionMatrix provides a builder for the DataframeClassificationSummaryMulticlassConfusionMatrix struct.
func NewDataframeClassificationSummaryMulticlassConfusionMatrixBuilder() *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder {
	r := DataframeClassificationSummaryMulticlassConfusionMatrixBuilder{
		&DataframeClassificationSummaryMulticlassConfusionMatrix{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeClassificationSummaryMulticlassConfusionMatrix struct
func (rb *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder) Build() DataframeClassificationSummaryMulticlassConfusionMatrix {
	return *rb.v
}

func (rb *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder) ConfusionMatrix(confusion_matrix []ConfusionMatrixItemBuilder) *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder {
	tmp := make([]ConfusionMatrixItem, len(confusion_matrix))
	for _, value := range confusion_matrix {
		tmp = append(tmp, value.Build())
	}
	rb.v.ConfusionMatrix = tmp
	return rb
}

func (rb *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder) OtherActualClassCount(otheractualclasscount int) *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder {
	rb.v.OtherActualClassCount = otheractualclasscount
	return rb
}
