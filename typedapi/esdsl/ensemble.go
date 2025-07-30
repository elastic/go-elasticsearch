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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ensemble struct {
	v *types.Ensemble
}

func NewEnsemble() *_ensemble {

	return &_ensemble{v: types.NewEnsemble()}

}

func (s *_ensemble) AggregateOutput(aggregateoutput types.AggregateOutputVariant) *_ensemble {

	s.v.AggregateOutput = aggregateoutput.AggregateOutputCaster()

	return s
}

func (s *_ensemble) ClassificationLabels(classificationlabels ...string) *_ensemble {

	for _, v := range classificationlabels {

		s.v.ClassificationLabels = append(s.v.ClassificationLabels, v)

	}
	return s
}

func (s *_ensemble) FeatureNames(featurenames ...string) *_ensemble {

	for _, v := range featurenames {

		s.v.FeatureNames = append(s.v.FeatureNames, v)

	}
	return s
}

func (s *_ensemble) TargetType(targettype string) *_ensemble {

	s.v.TargetType = &targettype

	return s
}

func (s *_ensemble) TrainedModels(trainedmodels ...types.TrainedModelVariant) *_ensemble {

	for _, v := range trainedmodels {

		s.v.TrainedModels = append(s.v.TrainedModels, *v.TrainedModelCaster())

	}
	return s
}

func (s *_ensemble) EnsembleCaster() *types.Ensemble {
	return s.v
}
