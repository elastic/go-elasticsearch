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

// DataframeEvaluationMetrics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeEvaluation.ts#L64-L71
type DataframeEvaluationMetrics struct {
	// AucRoc The AUC ROC (area under the curve of the receiver operating characteristic)
	// score and optionally the curve. It is calculated for a specific class
	// (provided as "class_name") treated as positive.
	AucRoc *DataframeEvaluationClassificationMetricsAucRoc `json:"auc_roc,omitempty"`
	// Precision Precision of predictions (per-class and average).
	Precision map[string]interface{} `json:"precision,omitempty"`
	// Recall Recall of predictions (per-class and average).
	Recall map[string]interface{} `json:"recall,omitempty"`
}

// DataframeEvaluationMetricsBuilder holds DataframeEvaluationMetrics struct and provides a builder API.
type DataframeEvaluationMetricsBuilder struct {
	v *DataframeEvaluationMetrics
}

// NewDataframeEvaluationMetrics provides a builder for the DataframeEvaluationMetrics struct.
func NewDataframeEvaluationMetricsBuilder() *DataframeEvaluationMetricsBuilder {
	r := DataframeEvaluationMetricsBuilder{
		&DataframeEvaluationMetrics{
			Precision: make(map[string]interface{}, 0),
			Recall:    make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationMetrics struct
func (rb *DataframeEvaluationMetricsBuilder) Build() DataframeEvaluationMetrics {
	return *rb.v
}

// AucRoc The AUC ROC (area under the curve of the receiver operating characteristic)
// score and optionally the curve. It is calculated for a specific class
// (provided as "class_name") treated as positive.

func (rb *DataframeEvaluationMetricsBuilder) AucRoc(aucroc *DataframeEvaluationClassificationMetricsAucRocBuilder) *DataframeEvaluationMetricsBuilder {
	v := aucroc.Build()
	rb.v.AucRoc = &v
	return rb
}

// Precision Precision of predictions (per-class and average).

func (rb *DataframeEvaluationMetricsBuilder) Precision(value map[string]interface{}) *DataframeEvaluationMetricsBuilder {
	rb.v.Precision = value
	return rb
}

// Recall Recall of predictions (per-class and average).

func (rb *DataframeEvaluationMetricsBuilder) Recall(value map[string]interface{}) *DataframeEvaluationMetricsBuilder {
	rb.v.Recall = value
	return rb
}
