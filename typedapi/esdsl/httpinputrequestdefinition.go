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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/httpinputmethod"
)

type _httpInputRequestDefinition struct {
	v *types.HttpInputRequestDefinition
}

func NewHttpInputRequestDefinition() *_httpInputRequestDefinition {

	return &_httpInputRequestDefinition{v: types.NewHttpInputRequestDefinition()}

}

func (s *_httpInputRequestDefinition) Auth(auth types.HttpInputAuthenticationVariant) *_httpInputRequestDefinition {

	s.v.Auth = auth.HttpInputAuthenticationCaster()

	return s
}

func (s *_httpInputRequestDefinition) Body(body string) *_httpInputRequestDefinition {

	s.v.Body = &body

	return s
}

func (s *_httpInputRequestDefinition) ConnectionTimeout(duration types.DurationVariant) *_httpInputRequestDefinition {

	s.v.ConnectionTimeout = *duration.DurationCaster()

	return s
}

func (s *_httpInputRequestDefinition) Headers(headers map[string]string) *_httpInputRequestDefinition {

	s.v.Headers = headers
	return s
}

func (s *_httpInputRequestDefinition) AddHeader(key string, value string) *_httpInputRequestDefinition {

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

func (s *_httpInputRequestDefinition) Host(host string) *_httpInputRequestDefinition {

	s.v.Host = &host

	return s
}

func (s *_httpInputRequestDefinition) Method(method httpinputmethod.HttpInputMethod) *_httpInputRequestDefinition {

	s.v.Method = &method
	return s
}

func (s *_httpInputRequestDefinition) Params(params map[string]string) *_httpInputRequestDefinition {

	s.v.Params = params
	return s
}

func (s *_httpInputRequestDefinition) AddParam(key string, value string) *_httpInputRequestDefinition {

	var tmp map[string]string
	if s.v.Params == nil {
		s.v.Params = make(map[string]string)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

func (s *_httpInputRequestDefinition) Path(path string) *_httpInputRequestDefinition {

	s.v.Path = &path

	return s
}

func (s *_httpInputRequestDefinition) Port(port uint) *_httpInputRequestDefinition {

	s.v.Port = &port

	return s
}

func (s *_httpInputRequestDefinition) Proxy(proxy types.HttpInputProxyVariant) *_httpInputRequestDefinition {

	s.v.Proxy = proxy.HttpInputProxyCaster()

	return s
}

func (s *_httpInputRequestDefinition) ReadTimeout(duration types.DurationVariant) *_httpInputRequestDefinition {

	s.v.ReadTimeout = *duration.DurationCaster()

	return s
}

func (s *_httpInputRequestDefinition) Scheme(scheme connectionscheme.ConnectionScheme) *_httpInputRequestDefinition {

	s.v.Scheme = &scheme
	return s
}

func (s *_httpInputRequestDefinition) Url(url string) *_httpInputRequestDefinition {

	s.v.Url = &url

	return s
}

func (s *_httpInputRequestDefinition) HttpInputRequestDefinitionCaster() *types.HttpInputRequestDefinition {
	return s.v
}
