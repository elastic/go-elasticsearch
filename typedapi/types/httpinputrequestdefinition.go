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


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/httpinputmethod"
)

// HttpInputRequestDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Input.ts#L74-L88
type HttpInputRequestDefinition struct {
	Auth              *HttpInputAuthentication           `json:"auth,omitempty"`
	Body              *string                            `json:"body,omitempty"`
	ConnectionTimeout *Duration                          `json:"connection_timeout,omitempty"`
	Headers           map[string]string                  `json:"headers,omitempty"`
	Host              *Host                              `json:"host,omitempty"`
	Method            *httpinputmethod.HttpInputMethod   `json:"method,omitempty"`
	Params            map[string]string                  `json:"params,omitempty"`
	Path              *string                            `json:"path,omitempty"`
	Port              *uint                              `json:"port,omitempty"`
	Proxy             *HttpInputProxy                    `json:"proxy,omitempty"`
	ReadTimeout       *Duration                          `json:"read_timeout,omitempty"`
	Scheme            *connectionscheme.ConnectionScheme `json:"scheme,omitempty"`
	Url               *string                            `json:"url,omitempty"`
}

// HttpInputRequestDefinitionBuilder holds HttpInputRequestDefinition struct and provides a builder API.
type HttpInputRequestDefinitionBuilder struct {
	v *HttpInputRequestDefinition
}

// NewHttpInputRequestDefinition provides a builder for the HttpInputRequestDefinition struct.
func NewHttpInputRequestDefinitionBuilder() *HttpInputRequestDefinitionBuilder {
	r := HttpInputRequestDefinitionBuilder{
		&HttpInputRequestDefinition{
			Headers: make(map[string]string, 0),
			Params:  make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the HttpInputRequestDefinition struct
func (rb *HttpInputRequestDefinitionBuilder) Build() HttpInputRequestDefinition {
	return *rb.v
}

func (rb *HttpInputRequestDefinitionBuilder) Auth(auth *HttpInputAuthenticationBuilder) *HttpInputRequestDefinitionBuilder {
	v := auth.Build()
	rb.v.Auth = &v
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Body(body string) *HttpInputRequestDefinitionBuilder {
	rb.v.Body = &body
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) ConnectionTimeout(connectiontimeout *DurationBuilder) *HttpInputRequestDefinitionBuilder {
	v := connectiontimeout.Build()
	rb.v.ConnectionTimeout = &v
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Headers(value map[string]string) *HttpInputRequestDefinitionBuilder {
	rb.v.Headers = value
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Host(host Host) *HttpInputRequestDefinitionBuilder {
	rb.v.Host = &host
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Method(method httpinputmethod.HttpInputMethod) *HttpInputRequestDefinitionBuilder {
	rb.v.Method = &method
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Params(value map[string]string) *HttpInputRequestDefinitionBuilder {
	rb.v.Params = value
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Path(path string) *HttpInputRequestDefinitionBuilder {
	rb.v.Path = &path
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Port(port uint) *HttpInputRequestDefinitionBuilder {
	rb.v.Port = &port
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Proxy(proxy *HttpInputProxyBuilder) *HttpInputRequestDefinitionBuilder {
	v := proxy.Build()
	rb.v.Proxy = &v
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) ReadTimeout(readtimeout *DurationBuilder) *HttpInputRequestDefinitionBuilder {
	v := readtimeout.Build()
	rb.v.ReadTimeout = &v
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Scheme(scheme connectionscheme.ConnectionScheme) *HttpInputRequestDefinitionBuilder {
	rb.v.Scheme = &scheme
	return rb
}

func (rb *HttpInputRequestDefinitionBuilder) Url(url string) *HttpInputRequestDefinitionBuilder {
	rb.v.Url = &url
	return rb
}
