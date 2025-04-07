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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/httpinputmethod"
)

type _webhookAction struct {
	v *types.WebhookAction
}

func NewWebhookAction() *_webhookAction {

	return &_webhookAction{v: types.NewWebhookAction()}

}

func (s *_webhookAction) Auth(auth types.HttpInputAuthenticationVariant) *_webhookAction {

	s.v.Auth = auth.HttpInputAuthenticationCaster()

	return s
}

func (s *_webhookAction) Body(body string) *_webhookAction {

	s.v.Body = &body

	return s
}

func (s *_webhookAction) ConnectionTimeout(duration types.DurationVariant) *_webhookAction {

	s.v.ConnectionTimeout = *duration.DurationCaster()

	return s
}

func (s *_webhookAction) Headers(headers map[string]string) *_webhookAction {

	s.v.Headers = headers
	return s
}

func (s *_webhookAction) AddHeader(key string, value string) *_webhookAction {

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

func (s *_webhookAction) Host(host string) *_webhookAction {

	s.v.Host = &host

	return s
}

func (s *_webhookAction) Method(method httpinputmethod.HttpInputMethod) *_webhookAction {

	s.v.Method = &method
	return s
}

func (s *_webhookAction) Params(params map[string]string) *_webhookAction {

	s.v.Params = params
	return s
}

func (s *_webhookAction) AddParam(key string, value string) *_webhookAction {

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

func (s *_webhookAction) Path(path string) *_webhookAction {

	s.v.Path = &path

	return s
}

func (s *_webhookAction) Port(port uint) *_webhookAction {

	s.v.Port = &port

	return s
}

func (s *_webhookAction) Proxy(proxy types.HttpInputProxyVariant) *_webhookAction {

	s.v.Proxy = proxy.HttpInputProxyCaster()

	return s
}

func (s *_webhookAction) ReadTimeout(duration types.DurationVariant) *_webhookAction {

	s.v.ReadTimeout = *duration.DurationCaster()

	return s
}

func (s *_webhookAction) Scheme(scheme connectionscheme.ConnectionScheme) *_webhookAction {

	s.v.Scheme = &scheme
	return s
}

func (s *_webhookAction) Url(url string) *_webhookAction {

	s.v.Url = &url

	return s
}

func (s *_webhookAction) WebhookActionCaster() *types.WebhookAction {
	return s.v
}
