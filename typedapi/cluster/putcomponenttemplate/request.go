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


package putcomponenttemplate

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putcomponenttemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/put_component_template/ClusterPutComponentTemplateRequest.ts#L29-L54
type Request struct {
	Aliases map[string]types.AliasDefinition `json:"aliases,omitempty"`

	Mappings *types.TypeMapping `json:"mappings,omitempty"`

	Meta_ *types.Metadata `json:"_meta,omitempty"`

	Settings *types.IndexSettings `json:"settings,omitempty"`

	Template types.IndexState `json:"template"`

	Version *types.VersionNumber `json:"version,omitempty"`
}

// RequestBuilder is the builder API for the putcomponenttemplate.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Aliases: make(map[string]types.AliasDefinition, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putcomponenttemplate request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Aliases(values map[string]*types.AliasDefinitionBuilder) *RequestBuilder {
	tmp := make(map[string]types.AliasDefinition, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aliases = tmp
	return rb
}

func (rb *RequestBuilder) Mappings(mappings *types.TypeMappingBuilder) *RequestBuilder {
	v := mappings.Build()
	rb.v.Mappings = &v
	return rb
}

func (rb *RequestBuilder) Meta_(meta_ *types.MetadataBuilder) *RequestBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

func (rb *RequestBuilder) Settings(settings *types.IndexSettingsBuilder) *RequestBuilder {
	v := settings.Build()
	rb.v.Settings = &v
	return rb
}

func (rb *RequestBuilder) Template(template *types.IndexStateBuilder) *RequestBuilder {
	v := template.Build()
	rb.v.Template = v
	return rb
}

func (rb *RequestBuilder) Version(version types.VersionNumber) *RequestBuilder {
	rb.v.Version = &version
	return rb
}
