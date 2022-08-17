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

// DataframeOutlierDetectionSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/evaluate_data_frame/types.ts#L24-L29
type DataframeOutlierDetectionSummary struct {
	AucRoc          *DataframeEvaluationSummaryAucRoc   `json:"auc_roc,omitempty"`
	ConfusionMatrix map[string]ConfusionMatrixThreshold `json:"confusion_matrix,omitempty"`
	Precision       map[string]float64                  `json:"precision,omitempty"`
	Recall          map[string]float64                  `json:"recall,omitempty"`
}

// DataframeOutlierDetectionSummaryBuilder holds DataframeOutlierDetectionSummary struct and provides a builder API.
type DataframeOutlierDetectionSummaryBuilder struct {
	v *DataframeOutlierDetectionSummary
}

// NewDataframeOutlierDetectionSummary provides a builder for the DataframeOutlierDetectionSummary struct.
func NewDataframeOutlierDetectionSummaryBuilder() *DataframeOutlierDetectionSummaryBuilder {
	r := DataframeOutlierDetectionSummaryBuilder{
		&DataframeOutlierDetectionSummary{
			ConfusionMatrix: make(map[string]ConfusionMatrixThreshold, 0),
			Precision:       make(map[string]float64, 0),
			Recall:          make(map[string]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DataframeOutlierDetectionSummary struct
func (rb *DataframeOutlierDetectionSummaryBuilder) Build() DataframeOutlierDetectionSummary {
	return *rb.v
}

func (rb *DataframeOutlierDetectionSummaryBuilder) AucRoc(aucroc *DataframeEvaluationSummaryAucRocBuilder) *DataframeOutlierDetectionSummaryBuilder {
	v := aucroc.Build()
	rb.v.AucRoc = &v
	return rb
}

func (rb *DataframeOutlierDetectionSummaryBuilder) ConfusionMatrix(values map[string]*ConfusionMatrixThresholdBuilder) *DataframeOutlierDetectionSummaryBuilder {
	tmp := make(map[string]ConfusionMatrixThreshold, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ConfusionMatrix = tmp
	return rb
}

func (rb *DataframeOutlierDetectionSummaryBuilder) Precision(value map[string]float64) *DataframeOutlierDetectionSummaryBuilder {
	rb.v.Precision = value
	return rb
}

func (rb *DataframeOutlierDetectionSummaryBuilder) Recall(value map[string]float64) *DataframeOutlierDetectionSummaryBuilder {
	rb.v.Recall = value
	return rb
}
