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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/responsecontenttype"
)

type _httpInput struct {
	v *types.HttpInput
}

func NewHttpInput() *_httpInput {

	return &_httpInput{v: types.NewHttpInput()}

}

func (s *_httpInput) Extract(extracts ...string) *_httpInput {

	for _, v := range extracts {

		s.v.Extract = append(s.v.Extract, v)

	}
	return s
}

func (s *_httpInput) Request(request types.HttpInputRequestDefinitionVariant) *_httpInput {

	s.v.Request = request.HttpInputRequestDefinitionCaster()

	return s
}

func (s *_httpInput) ResponseContentType(responsecontenttype responsecontenttype.ResponseContentType) *_httpInput {

	s.v.ResponseContentType = &responsecontenttype
	return s
}

func (s *_httpInput) WatcherInputCaster() *types.WatcherInput {
	container := types.NewWatcherInput()

	container.Http = s.v

	return container
}

func (s *_httpInput) HttpInputCaster() *types.HttpInput {
	return s.v
}
