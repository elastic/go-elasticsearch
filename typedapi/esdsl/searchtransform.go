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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchTransform struct {
	v *types.SearchTransform
}

func NewSearchTransform(request types.SearchInputRequestDefinitionVariant) *_searchTransform {

	tmp := &_searchTransform{v: types.NewSearchTransform()}

	tmp.Request(request)

	return tmp

}

func (s *_searchTransform) Request(request types.SearchInputRequestDefinitionVariant) *_searchTransform {

	s.v.Request = *request.SearchInputRequestDefinitionCaster()

	return s
}

func (s *_searchTransform) Timeout(duration types.DurationVariant) *_searchTransform {

	s.v.Timeout = *duration.DurationCaster()

	return s
}

func (s *_searchTransform) TransformContainerCaster() *types.TransformContainer {
	container := types.NewTransformContainer()

	container.Search = s.v

	return container
}

func (s *_searchTransform) SearchTransformCaster() *types.SearchTransform {
	return s.v
}
