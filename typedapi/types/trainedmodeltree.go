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

// TrainedModelTree type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L74-L79
type TrainedModelTree struct {
	ClassificationLabels []string               `json:"classification_labels,omitempty"`
	FeatureNames         []string               `json:"feature_names"`
	TargetType           *string                `json:"target_type,omitempty"`
	TreeStructure        []TrainedModelTreeNode `json:"tree_structure"`
}

// TrainedModelTreeBuilder holds TrainedModelTree struct and provides a builder API.
type TrainedModelTreeBuilder struct {
	v *TrainedModelTree
}

// NewTrainedModelTree provides a builder for the TrainedModelTree struct.
func NewTrainedModelTreeBuilder() *TrainedModelTreeBuilder {
	r := TrainedModelTreeBuilder{
		&TrainedModelTree{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelTree struct
func (rb *TrainedModelTreeBuilder) Build() TrainedModelTree {
	return *rb.v
}

func (rb *TrainedModelTreeBuilder) ClassificationLabels(classification_labels ...string) *TrainedModelTreeBuilder {
	rb.v.ClassificationLabels = classification_labels
	return rb
}

func (rb *TrainedModelTreeBuilder) FeatureNames(feature_names ...string) *TrainedModelTreeBuilder {
	rb.v.FeatureNames = feature_names
	return rb
}

func (rb *TrainedModelTreeBuilder) TargetType(targettype string) *TrainedModelTreeBuilder {
	rb.v.TargetType = &targettype
	return rb
}

func (rb *TrainedModelTreeBuilder) TreeStructure(tree_structure []TrainedModelTreeNodeBuilder) *TrainedModelTreeBuilder {
	tmp := make([]TrainedModelTreeNode, len(tree_structure))
	for _, value := range tree_structure {
		tmp = append(tmp, value.Build())
	}
	rb.v.TreeStructure = tmp
	return rb
}
