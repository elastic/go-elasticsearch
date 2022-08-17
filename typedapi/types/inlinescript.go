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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
)

// InlineScript type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Scripting.ts#L45-L50
type InlineScript struct {
	Lang    *scriptlanguage.ScriptLanguage `json:"lang,omitempty"`
	Options map[string]string              `json:"options,omitempty"`
	Params  map[string]interface{}         `json:"params,omitempty"`
	Source  string                         `json:"source"`
}

// InlineScriptBuilder holds InlineScript struct and provides a builder API.
type InlineScriptBuilder struct {
	v *InlineScript
}

// NewInlineScript provides a builder for the InlineScript struct.
func NewInlineScriptBuilder() *InlineScriptBuilder {
	r := InlineScriptBuilder{
		&InlineScript{
			Options: make(map[string]string, 0),
			Params:  make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InlineScript struct
func (rb *InlineScriptBuilder) Build() InlineScript {
	return *rb.v
}

func (rb *InlineScriptBuilder) Lang(lang scriptlanguage.ScriptLanguage) *InlineScriptBuilder {
	rb.v.Lang = &lang
	return rb
}

func (rb *InlineScriptBuilder) Options(value map[string]string) *InlineScriptBuilder {
	rb.v.Options = value
	return rb
}

func (rb *InlineScriptBuilder) Params(value map[string]interface{}) *InlineScriptBuilder {
	rb.v.Params = value
	return rb
}

func (rb *InlineScriptBuilder) Source(source string) *InlineScriptBuilder {
	rb.v.Source = source
	return rb
}
