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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _embedding struct {
	v *types.Embedding
}

func NewEmbedding() *_embedding {

	return &_embedding{v: types.NewEmbedding()}

}

func (s *_embedding) InferenceId(inferenceid string) *_embedding {

	s.v.InferenceId = &inferenceid

	return s
}

func (s *_embedding) Input(knnembeddinginput types.KnnEmbeddingInputVariant) *_embedding {

	s.v.Input = *knnembeddinginput.KnnEmbeddingInputCaster()

	return s
}

func (s *_embedding) Timeout(duration types.DurationVariant) *_embedding {

	s.v.Timeout = *duration.DurationCaster()

	return s
}

func (s *_embedding) QueryVectorBuilderCaster() *types.QueryVectorBuilder {
	container := types.NewQueryVectorBuilder()

	container.Embedding = s.v

	return container
}

func (s *_embedding) EmbeddingCaster() *types.Embedding {
	return s.v
}
