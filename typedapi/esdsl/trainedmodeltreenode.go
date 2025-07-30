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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trainedModelTreeNode struct {
	v *types.TrainedModelTreeNode
}

func NewTrainedModelTreeNode(nodeindex int) *_trainedModelTreeNode {

	tmp := &_trainedModelTreeNode{v: types.NewTrainedModelTreeNode()}

	tmp.NodeIndex(nodeindex)

	return tmp

}

func (s *_trainedModelTreeNode) DecisionType(decisiontype string) *_trainedModelTreeNode {

	s.v.DecisionType = &decisiontype

	return s
}

func (s *_trainedModelTreeNode) DefaultLeft(defaultleft bool) *_trainedModelTreeNode {

	s.v.DefaultLeft = &defaultleft

	return s
}

func (s *_trainedModelTreeNode) LeafValue(leafvalue types.Float64) *_trainedModelTreeNode {

	s.v.LeafValue = &leafvalue

	return s
}

func (s *_trainedModelTreeNode) LeftChild(leftchild int) *_trainedModelTreeNode {

	s.v.LeftChild = &leftchild

	return s
}

func (s *_trainedModelTreeNode) NodeIndex(nodeindex int) *_trainedModelTreeNode {

	s.v.NodeIndex = nodeindex

	return s
}

func (s *_trainedModelTreeNode) RightChild(rightchild int) *_trainedModelTreeNode {

	s.v.RightChild = &rightchild

	return s
}

func (s *_trainedModelTreeNode) SplitFeature(splitfeature int) *_trainedModelTreeNode {

	s.v.SplitFeature = &splitfeature

	return s
}

func (s *_trainedModelTreeNode) SplitGain(splitgain int) *_trainedModelTreeNode {

	s.v.SplitGain = &splitgain

	return s
}

func (s *_trainedModelTreeNode) Threshold(threshold types.Float64) *_trainedModelTreeNode {

	s.v.Threshold = &threshold

	return s
}

func (s *_trainedModelTreeNode) TrainedModelTreeNodeCaster() *types.TrainedModelTreeNode {
	return s.v
}
