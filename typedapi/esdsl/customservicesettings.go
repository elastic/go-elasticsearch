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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/customserviceinputtype"
)

type _customServiceSettings struct {
	v *types.CustomServiceSettings
}

func NewCustomServiceSettings(request types.CustomRequestParamsVariant, response types.CustomResponseParamsVariant) *_customServiceSettings {

	tmp := &_customServiceSettings{v: types.NewCustomServiceSettings()}

	tmp.Request(request)

	tmp.Response(response)

	return tmp

}

func (s *_customServiceSettings) BatchSize(batchsize int) *_customServiceSettings {

	s.v.BatchSize = &batchsize

	return s
}

func (s *_customServiceSettings) Headers(headers map[string]string) *_customServiceSettings {

	s.v.Headers = headers
	return s
}

func (s *_customServiceSettings) AddHeader(key string, value string) *_customServiceSettings {

	var tmp map[string]string
	if s.v.Headers == nil {
		s.v.Headers = make(map[string]string)
	} else {
		tmp = s.v.Headers
	}

	tmp[key] = value

	s.v.Headers = tmp
	return s
}

func (s *_customServiceSettings) InputType(inputtype map[customserviceinputtype.CustomServiceInputType]string) *_customServiceSettings {

	s.v.InputType = inputtype
	return s
}

func (s *_customServiceSettings) AddInputType(key customserviceinputtype.CustomServiceInputType, value string) *_customServiceSettings {

	var tmp map[customserviceinputtype.CustomServiceInputType]string
	if s.v.InputType == nil {
		s.v.InputType = make(map[customserviceinputtype.CustomServiceInputType]string)
	} else {
		tmp = s.v.InputType
	}

	tmp[key] = value

	s.v.InputType = tmp
	return s
}

func (s *_customServiceSettings) QueryParameters(queryparameters ...[]string) *_customServiceSettings {

	for _, v := range queryparameters {

		s.v.QueryParameters = append(s.v.QueryParameters, v)

	}
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

func (s *_customServiceSettings) SecretParameters(secretparameters map[string]string) *_customServiceSettings {

	s.v.SecretParameters = secretparameters
	return s
}

func (s *_customServiceSettings) AddSecretParameter(key string, value string) *_customServiceSettings {

	var tmp map[string]string
	if s.v.SecretParameters == nil {
		s.v.SecretParameters = make(map[string]string)
	} else {
		tmp = s.v.SecretParameters
	}

	tmp[key] = value

	s.v.SecretParameters = tmp
	return s
}

func (s *_customServiceSettings) Url(url string) *_customServiceSettings {

	s.v.Url = &url

	return s
}

func (s *_customServiceSettings) CustomServiceSettingsCaster() *types.CustomServiceSettings {
	return s.v
}
