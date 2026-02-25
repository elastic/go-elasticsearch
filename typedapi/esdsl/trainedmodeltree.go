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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trainedModelTree struct {
	v *types.TrainedModelTree
}

func NewTrainedModelTree() *_trainedModelTree {

	return &_trainedModelTree{v: types.NewTrainedModelTree()}

}

func (s *_trainedModelTree) ClassificationLabels(classificationlabels ...string) *_trainedModelTree {

	for _, v := range classificationlabels {

		s.v.ClassificationLabels = append(s.v.ClassificationLabels, v)

	}
	return s
}

func (s *_trainedModelTree) FeatureNames(featurenames ...string) *_trainedModelTree {

	for _, v := range featurenames {

		s.v.FeatureNames = append(s.v.FeatureNames, v)

	}
	return s
}

func (s *_trainedModelTree) TargetType(targettype string) *_trainedModelTree {

	s.v.TargetType = &targettype

	return s
}

func (s *_trainedModelTree) TreeStructure(treestructures ...types.TrainedModelTreeNodeVariant) *_trainedModelTree {

	for _, v := range treestructures {

		s.v.TreeStructure = append(s.v.TreeStructure, *v.TrainedModelTreeNodeCaster())

	}
	return s
}

func (s *_trainedModelTree) TrainedModelTreeCaster() *types.TrainedModelTree {
	return s.v
}
