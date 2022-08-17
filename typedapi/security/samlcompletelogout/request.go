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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package samlcompletelogout

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package samlcompletelogout
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/saml_complete_logout/Request.ts#L23-L40
type Request struct {

	// Content If the SAML IdP sends the logout response with the HTTP-Post binding, this
	// field must be set to the value of the SAMLResponse form parameter from the
	// logout response.
	Content *string `json:"content,omitempty"`

	// Ids A json array with all the valid SAML Request Ids that the caller of the API
	// has for the current user.
	Ids types.Ids `json:"ids"`

	// QueryString If the SAML IdP sends the logout response with the HTTP-Redirect binding,
	// this field must be set to the query string of the redirect URI.
	QueryString *string `json:"query_string,omitempty"`

	// Realm The name of the SAML realm in Elasticsearch for which the configuration is
	// used to verify the logout response.
	Realm string `json:"realm"`
}

// RequestBuilder is the builder API for the samlcompletelogout.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Samlcompletelogout request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Content(content string) *RequestBuilder {
	rb.v.Content = &content
	return rb
}

func (rb *RequestBuilder) Ids(ids *types.IdsBuilder) *RequestBuilder {
	v := ids.Build()
	rb.v.Ids = v
	return rb
}

func (rb *RequestBuilder) QueryString(querystring string) *RequestBuilder {
	rb.v.QueryString = &querystring
	return rb
}

func (rb *RequestBuilder) Realm(realm string) *RequestBuilder {
	rb.v.Realm = realm
	return rb
}
