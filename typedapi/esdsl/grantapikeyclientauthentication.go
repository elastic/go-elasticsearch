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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/grantapikeyclientauthenticationscheme"
)

type _grantApiKeyClientAuthentication struct {
	v *types.GrantApiKeyClientAuthentication
}

func NewGrantApiKeyClientAuthentication(scheme grantapikeyclientauthenticationscheme.GrantApiKeyClientAuthenticationScheme, value string) *_grantApiKeyClientAuthentication {

	tmp := &_grantApiKeyClientAuthentication{v: types.NewGrantApiKeyClientAuthentication()}

	tmp.Scheme(scheme)

	tmp.Value(value)

	return tmp

}

func (s *_grantApiKeyClientAuthentication) Scheme(scheme grantapikeyclientauthenticationscheme.GrantApiKeyClientAuthenticationScheme) *_grantApiKeyClientAuthentication {

	s.v.Scheme = scheme
	return s
}

func (s *_grantApiKeyClientAuthentication) Value(value string) *_grantApiKeyClientAuthentication {

	s.v.Value = value

	return s
}

func (s *_grantApiKeyClientAuthentication) GrantApiKeyClientAuthenticationCaster() *types.GrantApiKeyClientAuthentication {
	return s.v
}
