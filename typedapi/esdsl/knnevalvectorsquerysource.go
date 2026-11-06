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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnEvalVectorsQuerySource struct {
	v *types.KnnEvalVectorsQuerySource
}

func NewKnnEvalVectorsQuerySource() *_knnEvalVectorsQuerySource {

	return &_knnEvalVectorsQuerySource{v: types.NewKnnEvalVectorsQuerySource()}

}

func (s *_knnEvalVectorsQuerySource) Vectors(vectors ...types.KnnEvalQueryVectorVariant) *_knnEvalVectorsQuerySource {

	for _, v := range vectors {

		s.v.Vectors = append(s.v.Vectors, *v.KnnEvalQueryVectorCaster())

	}
	return s
}

func (s *_knnEvalVectorsQuerySource) VectorsValues(vectorsvalues []types.KnnEvalQueryVector) *_knnEvalVectorsQuerySource {

	s.v.Vectors = vectorsvalues
	return s
}

func (s *_knnEvalVectorsQuerySource) KnnEvalVectorsQuerySourceCaster() *types.KnnEvalVectorsQuerySource {
	return s.v
}
