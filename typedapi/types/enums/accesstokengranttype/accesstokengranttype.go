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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Package accesstokengranttype
package accesstokengranttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/security/get_token/types.ts#L23-L28
type AccessTokenGrantType struct {
	Name string
}

var (
	Password = AccessTokenGrantType{"password"}

	Clientcredentials = AccessTokenGrantType{"client_credentials"}

	Kerberos = AccessTokenGrantType{"_kerberos"}

	Refreshtoken = AccessTokenGrantType{"refresh_token"}
)

func (a AccessTokenGrantType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AccessTokenGrantType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "password":
		*a = Password
	case "client_credentials":
		*a = Clientcredentials
	case "_kerberos":
		*a = Kerberos
	case "refresh_token":
		*a = Refreshtoken
	default:
		*a = AccessTokenGrantType{string(text)}
	}

	return nil
}

func (a AccessTokenGrantType) String() string {
	return a.Name
}
