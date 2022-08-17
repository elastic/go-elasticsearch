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

// TrainedModelTreeNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L81-L91
type TrainedModelTreeNode struct {
	DecisionType *string  `json:"decision_type,omitempty"`
	DefaultLeft  *bool    `json:"default_left,omitempty"`
	LeafValue    *float64 `json:"leaf_value,omitempty"`
	LeftChild    *int     `json:"left_child,omitempty"`
	NodeIndex    int      `json:"node_index"`
	RightChild   *int     `json:"right_child,omitempty"`
	SplitFeature *int     `json:"split_feature,omitempty"`
	SplitGain    *int     `json:"split_gain,omitempty"`
	Threshold    *float64 `json:"threshold,omitempty"`
}

// TrainedModelTreeNodeBuilder holds TrainedModelTreeNode struct and provides a builder API.
type TrainedModelTreeNodeBuilder struct {
	v *TrainedModelTreeNode
}

// NewTrainedModelTreeNode provides a builder for the TrainedModelTreeNode struct.
func NewTrainedModelTreeNodeBuilder() *TrainedModelTreeNodeBuilder {
	r := TrainedModelTreeNodeBuilder{
		&TrainedModelTreeNode{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelTreeNode struct
func (rb *TrainedModelTreeNodeBuilder) Build() TrainedModelTreeNode {
	return *rb.v
}

func (rb *TrainedModelTreeNodeBuilder) DecisionType(decisiontype string) *TrainedModelTreeNodeBuilder {
	rb.v.DecisionType = &decisiontype
	return rb
}

func (rb *TrainedModelTreeNodeBuilder) DefaultLeft(defaultleft bool) *TrainedModelTreeNodeBuilder {
	rb.v.DefaultLeft = &defaultleft
	return rb
}

func (rb *TrainedModelTreeNodeBuilder) LeafValue(leafvalue float64) *TrainedModelTreeNodeBuilder {
	rb.v.LeafValue = &leafvalue
	return rb
}

func (rb *TrainedModelTreeNodeBuilder) LeftChild(leftchild int) *TrainedModelTreeNodeBuilder {
	rb.v.LeftChild = &leftchild
	return rb
}

func (rb *TrainedModelTreeNodeBuilder) NodeIndex(nodeindex int) *TrainedModelTreeNodeBuilder {
	rb.v.NodeIndex = nodeindex
	return rb
}

func (rb *TrainedModelTreeNodeBuilder) RightChild(rightchild int) *TrainedModelTreeNodeBuilder {
	rb.v.RightChild = &rightchild
	return rb
}

func (rb *TrainedModelTreeNodeBuilder) SplitFeature(splitfeature int) *TrainedModelTreeNodeBuilder {
	rb.v.SplitFeature = &splitfeature
	return rb
}

func (rb *TrainedModelTreeNodeBuilder) SplitGain(splitgain int) *TrainedModelTreeNodeBuilder {
	rb.v.SplitGain = &splitgain
	return rb
}

func (rb *TrainedModelTreeNodeBuilder) Threshold(threshold float64) *TrainedModelTreeNodeBuilder {
	rb.v.Threshold = &threshold
	return rb
}
