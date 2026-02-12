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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _searchTemplateRequestBody struct {
	v *types.SearchTemplateRequestBody
}

func NewSearchTemplateRequestBody() *_searchTemplateRequestBody {

	return &_searchTemplateRequestBody{v: types.NewSearchTemplateRequestBody()}

}

func (s *_searchTemplateRequestBody) Explain(explain bool) *_searchTemplateRequestBody {

	s.v.Explain = &explain

	return s
}

func (s *_searchTemplateRequestBody) Id(id string) *_searchTemplateRequestBody {

	s.v.Id = &id

	return s
}

func (s *_searchTemplateRequestBody) Params(params map[string]json.RawMessage) *_searchTemplateRequestBody {

	s.v.Params = params
	return s
}

func (s *_searchTemplateRequestBody) AddParam(key string, value json.RawMessage) *_searchTemplateRequestBody {

	var tmp map[string]json.RawMessage
	if s.v.Params == nil {
		s.v.Params = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

func (s *_searchTemplateRequestBody) Profile(profile bool) *_searchTemplateRequestBody {

	s.v.Profile = &profile

	return s
}

func (s *_searchTemplateRequestBody) Source(source string) *_searchTemplateRequestBody {

	s.v.Source = &source

	return s
}

func (s *_searchTemplateRequestBody) SearchTemplateRequestBodyCaster() *types.SearchTemplateRequestBody {
	return s.v
}
