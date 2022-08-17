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

// Ensemble type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L93-L99
type Ensemble struct {
	AggregateOutput      *AggregateOutput `json:"aggregate_output,omitempty"`
	ClassificationLabels []string         `json:"classification_labels,omitempty"`
	FeatureNames         []string         `json:"feature_names,omitempty"`
	TargetType           *string          `json:"target_type,omitempty"`
	TrainedModels        []TrainedModel   `json:"trained_models"`
}

// EnsembleBuilder holds Ensemble struct and provides a builder API.
type EnsembleBuilder struct {
	v *Ensemble
}

// NewEnsemble provides a builder for the Ensemble struct.
func NewEnsembleBuilder() *EnsembleBuilder {
	r := EnsembleBuilder{
		&Ensemble{},
	}

	return &r
}

// Build finalize the chain and returns the Ensemble struct
func (rb *EnsembleBuilder) Build() Ensemble {
	return *rb.v
}

func (rb *EnsembleBuilder) AggregateOutput(aggregateoutput *AggregateOutputBuilder) *EnsembleBuilder {
	v := aggregateoutput.Build()
	rb.v.AggregateOutput = &v
	return rb
}

func (rb *EnsembleBuilder) ClassificationLabels(classification_labels ...string) *EnsembleBuilder {
	rb.v.ClassificationLabels = classification_labels
	return rb
}

func (rb *EnsembleBuilder) FeatureNames(feature_names ...string) *EnsembleBuilder {
	rb.v.FeatureNames = feature_names
	return rb
}

func (rb *EnsembleBuilder) TargetType(targettype string) *EnsembleBuilder {
	rb.v.TargetType = &targettype
	return rb
}

func (rb *EnsembleBuilder) TrainedModels(trained_models []TrainedModelBuilder) *EnsembleBuilder {
	tmp := make([]TrainedModel, len(trained_models))
	for _, value := range trained_models {
		tmp = append(tmp, value.Build())
	}
	rb.v.TrainedModels = tmp
	return rb
}
