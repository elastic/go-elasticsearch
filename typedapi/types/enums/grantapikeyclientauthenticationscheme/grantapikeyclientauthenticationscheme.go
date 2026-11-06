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

// Package grantapikeyclientauthenticationscheme
package grantapikeyclientauthenticationscheme

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/security/grant_api_key/types.ts#L60-L65
type GrantApiKeyClientAuthenticationScheme struct {
	Name string
}

var (

	// SharedSecret The client authenticates with the shared secret configured for the JWT realm.
	SharedSecret = GrantApiKeyClientAuthenticationScheme{"SharedSecret"}
)

func (g GrantApiKeyClientAuthenticationScheme) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GrantApiKeyClientAuthenticationScheme) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "sharedsecret":
		*g = SharedSecret
	default:
		*g = GrantApiKeyClientAuthenticationScheme{string(text)}
	}

	return nil
}

func (g GrantApiKeyClientAuthenticationScheme) String() string {
	return g.Name
}
