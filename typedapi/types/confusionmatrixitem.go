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

// ConfusionMatrixItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L84-L89
type ConfusionMatrixItem struct {
	ActualClass                 Name                        `json:"actual_class"`
	ActualClassDocCount         int                         `json:"actual_class_doc_count"`
	OtherPredictedClassDocCount int                         `json:"other_predicted_class_doc_count"`
	PredictedClasses            []ConfusionMatrixPrediction `json:"predicted_classes"`
}

// ConfusionMatrixItemBuilder holds ConfusionMatrixItem struct and provides a builder API.
type ConfusionMatrixItemBuilder struct {
	v *ConfusionMatrixItem
}

// NewConfusionMatrixItem provides a builder for the ConfusionMatrixItem struct.
func NewConfusionMatrixItemBuilder() *ConfusionMatrixItemBuilder {
	r := ConfusionMatrixItemBuilder{
		&ConfusionMatrixItem{},
	}

	return &r
}

// Build finalize the chain and returns the ConfusionMatrixItem struct
func (rb *ConfusionMatrixItemBuilder) Build() ConfusionMatrixItem {
	return *rb.v
}

func (rb *ConfusionMatrixItemBuilder) ActualClass(actualclass Name) *ConfusionMatrixItemBuilder {
	rb.v.ActualClass = actualclass
	return rb
}

func (rb *ConfusionMatrixItemBuilder) ActualClassDocCount(actualclassdoccount int) *ConfusionMatrixItemBuilder {
	rb.v.ActualClassDocCount = actualclassdoccount
	return rb
}

func (rb *ConfusionMatrixItemBuilder) OtherPredictedClassDocCount(otherpredictedclassdoccount int) *ConfusionMatrixItemBuilder {
	rb.v.OtherPredictedClassDocCount = otherpredictedclassdoccount
	return rb
}

func (rb *ConfusionMatrixItemBuilder) PredictedClasses(predicted_classes []ConfusionMatrixPredictionBuilder) *ConfusionMatrixItemBuilder {
	tmp := make([]ConfusionMatrixPrediction, len(predicted_classes))
	for _, value := range predicted_classes {
		tmp = append(tmp, value.Build())
	}
	rb.v.PredictedClasses = tmp
	return rb
}
