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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _requestEmbedding struct {
	v *types.RequestEmbedding
}

func NewRequestEmbedding() *_requestEmbedding {

	return &_requestEmbedding{v: types.NewRequestEmbedding()}

}

func (s *_requestEmbedding) Input(embeddinginput types.EmbeddingInputVariant) *_requestEmbedding {

	s.v.Input = *embeddinginput.EmbeddingInputCaster()

	return s
}

func (s *_requestEmbedding) InputType(inputtype string) *_requestEmbedding {

	s.v.InputType = &inputtype

	return s
}

func (s *_requestEmbedding) TaskSettings(tasksettings json.RawMessage) *_requestEmbedding {

	s.v.TaskSettings = tasksettings

	return s
}

func (s *_requestEmbedding) RequestEmbeddingCaster() *types.RequestEmbedding {
	return s.v
}
