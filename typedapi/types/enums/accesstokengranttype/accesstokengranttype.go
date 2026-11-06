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

// Package accesstokengranttype
package accesstokengranttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/security/get_token/types.ts#L25-L60
type AccessTokenGrantType struct {
	Name string
}

var (

	// Password This grant type implements the Resource Owner Password Credentials Grant of
	// OAuth2. In this grant, a trusted client exchanges the end user's credentials
	// for an access token and (possibly) a refresh token. The request needs to be
	// made by an authenticated user but happens on behalf of another authenticated
	// user (the one whose credentials are passed as request parameters). This grant
	// type is not suitable or designed for the self-service user creation of
	// tokens.
	Password = AccessTokenGrantType{"password"}

	// Clientcredentials This grant type implements the Client Credentials Grant of OAuth2. It is
	// geared for machine to machine communication and is not suitable or designed
	// for the self-service user creation of tokens. It generates only access tokens
	// that cannot be refreshed. The premise is that the entity that uses
	// `client_credentials` has constant access to a set of (client, not end-user)
	// credentials and can authenticate itself at will.
	Clientcredentials = AccessTokenGrantType{"client_credentials"}

	// Kerberos This grant type is supported internally and implements SPNEGO based Kerberos
	// support. The `_kerberos` grant type may change from version to version.
	Kerberos = AccessTokenGrantType{"_kerberos"}

	// Refreshtoken This grant type implements the Refresh Token Grant of OAuth2. In this grant a
	// user exchanges a previously issued refresh token for a new access token and a
	// new refresh token.
	Refreshtoken = AccessTokenGrantType{"refresh_token"}

	// Usermanagedserviceaccount This grant type is supported internally and exchanges a user-managed service
	// account token for an access token that represents that service account. The
	// request must be made by an authenticated caller with the `manage_token`
	// cluster privilege, on behalf of the service account whose token is passed in
	// the `service_account_token` parameter. The service account token is
	// authenticated on every exchange, so a disabled account or an invalidated
	// token is rejected. Tokens that belong to built-in (`elastic/*`) service
	// accounts are rejected. It generates only access tokens that cannot be
	// refreshed. The `_user_managed_service_account` grant type may change from
	// version to version.
	Usermanagedserviceaccount = AccessTokenGrantType{"_user_managed_service_account"}
)

func (a AccessTokenGrantType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AccessTokenGrantType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "password":
		*a = Password
	case "client_credentials":
		*a = Clientcredentials
	case "_kerberos":
		*a = Kerberos
	case "refresh_token":
		*a = Refreshtoken
	case "_user_managed_service_account":
		*a = Usermanagedserviceaccount
	default:
		*a = AccessTokenGrantType{string(text)}
	}

	return nil
}

func (a AccessTokenGrantType) String() string {
	return a.Name
}
