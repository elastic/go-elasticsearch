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

// DataframeEvaluationClassificationMetricsAucRoc type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeEvaluation.ts#L85-L90
type DataframeEvaluationClassificationMetricsAucRoc struct {
	// ClassName Name of the only class that is treated as positive during AUC ROC
	// calculation. Other classes are treated as negative ("one-vs-all" strategy).
	// All the evaluated documents must have class_name in the list of their top
	// classes.
	ClassName *Name `json:"class_name,omitempty"`
	// IncludeCurve Whether or not the curve should be returned in addition to the score. Default
	// value is false.
	IncludeCurve *bool `json:"include_curve,omitempty"`
}

// DataframeEvaluationClassificationMetricsAucRocBuilder holds DataframeEvaluationClassificationMetricsAucRoc struct and provides a builder API.
type DataframeEvaluationClassificationMetricsAucRocBuilder struct {
	v *DataframeEvaluationClassificationMetricsAucRoc
}

// NewDataframeEvaluationClassificationMetricsAucRoc provides a builder for the DataframeEvaluationClassificationMetricsAucRoc struct.
func NewDataframeEvaluationClassificationMetricsAucRocBuilder() *DataframeEvaluationClassificationMetricsAucRocBuilder {
	r := DataframeEvaluationClassificationMetricsAucRocBuilder{
		&DataframeEvaluationClassificationMetricsAucRoc{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationClassificationMetricsAucRoc struct
func (rb *DataframeEvaluationClassificationMetricsAucRocBuilder) Build() DataframeEvaluationClassificationMetricsAucRoc {
	return *rb.v
}

// ClassName Name of the only class that is treated as positive during AUC ROC
// calculation. Other classes are treated as negative ("one-vs-all" strategy).
// All the evaluated documents must have class_name in the list of their top
// classes.

func (rb *DataframeEvaluationClassificationMetricsAucRocBuilder) ClassName(classname Name) *DataframeEvaluationClassificationMetricsAucRocBuilder {
	rb.v.ClassName = &classname
	return rb
}

// IncludeCurve Whether or not the curve should be returned in addition to the score. Default
// value is false.

func (rb *DataframeEvaluationClassificationMetricsAucRocBuilder) IncludeCurve(includecurve bool) *DataframeEvaluationClassificationMetricsAucRocBuilder {
	rb.v.IncludeCurve = &includecurve
	return rb
}
