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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _customServiceSettings struct {
	v *types.CustomServiceSettings
}

func NewCustomServiceSettings(request types.CustomRequestParamsVariant, response types.CustomResponseParamsVariant, secretparameters json.RawMessage) *_customServiceSettings {

	tmp := &_customServiceSettings{v: types.NewCustomServiceSettings()}

	tmp.Request(request)

	tmp.Response(response)

	tmp.SecretParameters(secretparameters)

	return tmp

}

func (s *_customServiceSettings) Headers(headers json.RawMessage) *_customServiceSettings {

	s.v.Headers = headers

	return s
}

func (s *_customServiceSettings) InputType(inputtype json.RawMessage) *_customServiceSettings {

	s.v.InputType = inputtype

	return s
}

func (s *_customServiceSettings) QueryParameters(queryparameters json.RawMessage) *_customServiceSettings {

	s.v.QueryParameters = queryparameters

	return s
}

func (s *_customServiceSettings) Request(request types.CustomRequestParamsVariant) *_customServiceSettings {

	s.v.Request = *request.CustomRequestParamsCaster()

	return s
}

func (s *_customServiceSettings) Response(response types.CustomResponseParamsVariant) *_customServiceSettings {

	s.v.Response = *response.CustomResponseParamsCaster()

	return s
}

func (s *_customServiceSettings) SecretParameters(secretparameters json.RawMessage) *_customServiceSettings {

	s.v.SecretParameters = secretparameters

	return s
}

func (s *_customServiceSettings) Url(url string) *_customServiceSettings {

	s.v.Url = &url

	return s
}

func (s *_customServiceSettings) CustomServiceSettingsCaster() *types.CustomServiceSettings {
	return s.v
}
