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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _transformContainer struct {
	v *types.TransformContainer
}

func NewTransformContainer() *_transformContainer {
	return &_transformContainer{v: types.NewTransformContainer()}
}

// AdditionalTransformContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_transformContainer) AdditionalTransformContainerProperty(key string, value json.RawMessage) *_transformContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalTransformContainerProperty = tmp
	return s
}

func (s *_transformContainer) Chain(chains ...types.TransformContainerVariant) *_transformContainer {

	for _, v := range chains {

		s.v.Chain = append(s.v.Chain, *v.TransformContainerCaster())

	}
	return s
}

func (s *_transformContainer) Script(script types.ScriptTransformVariant) *_transformContainer {

	s.v.Script = script.ScriptTransformCaster()

	return s
}

func (s *_transformContainer) Search(search types.SearchTransformVariant) *_transformContainer {

	s.v.Search = search.SearchTransformCaster()

	return s
}

func (s *_transformContainer) TransformContainerCaster() *types.TransformContainer {
	return s.v
}
