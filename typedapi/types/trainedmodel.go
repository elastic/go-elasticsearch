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

// TrainedModel type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L60-L72
type TrainedModel struct {
	// Ensemble The definition for an ensemble model
	Ensemble *Ensemble `json:"ensemble,omitempty"`
	// Tree The definition for a binary decision tree.
	Tree *TrainedModelTree `json:"tree,omitempty"`
	// TreeNode The definition of a node in a tree.
	// There are two major types of nodes: leaf nodes and not-leaf nodes.
	// - Leaf nodes only need node_index and leaf_value defined.
	// - All other nodes need split_feature, left_child, right_child, threshold,
	// decision_type, and default_left defined.
	TreeNode *TrainedModelTreeNode `json:"tree_node,omitempty"`
}

// TrainedModelBuilder holds TrainedModel struct and provides a builder API.
type TrainedModelBuilder struct {
	v *TrainedModel
}

// NewTrainedModel provides a builder for the TrainedModel struct.
func NewTrainedModelBuilder() *TrainedModelBuilder {
	r := TrainedModelBuilder{
		&TrainedModel{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModel struct
func (rb *TrainedModelBuilder) Build() TrainedModel {
	return *rb.v
}

// Ensemble The definition for an ensemble model

func (rb *TrainedModelBuilder) Ensemble(ensemble *EnsembleBuilder) *TrainedModelBuilder {
	v := ensemble.Build()
	rb.v.Ensemble = &v
	return rb
}

// Tree The definition for a binary decision tree.

func (rb *TrainedModelBuilder) Tree(tree *TrainedModelTreeBuilder) *TrainedModelBuilder {
	v := tree.Build()
	rb.v.Tree = &v
	return rb
}

// TreeNode The definition of a node in a tree.
// There are two major types of nodes: leaf nodes and not-leaf nodes.
// - Leaf nodes only need node_index and leaf_value defined.
// - All other nodes need split_feature, left_child, right_child, threshold,
// decision_type, and default_left defined.

func (rb *TrainedModelBuilder) TreeNode(treenode *TrainedModelTreeNodeBuilder) *TrainedModelBuilder {
	v := treenode.Build()
	rb.v.TreeNode = &v
	return rb
}
