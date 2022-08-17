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

// StoredScript type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Scripting.ts#L35-L39
type StoredScript struct {
	Lang    scriptlanguage.ScriptLanguage `json:"lang"`
	Options map[string]string             `json:"options,omitempty"`
	Source  string                        `json:"source"`
}

// StoredScriptBuilder holds StoredScript struct and provides a builder API.
type StoredScriptBuilder struct {
	v *StoredScript
}

// NewStoredScript provides a builder for the StoredScript struct.
func NewStoredScriptBuilder() *StoredScriptBuilder {
	r := StoredScriptBuilder{
		&StoredScript{
			Options: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the StoredScript struct
func (rb *StoredScriptBuilder) Build() StoredScript {
	return *rb.v
}

func (rb *StoredScriptBuilder) Lang(lang scriptlanguage.ScriptLanguage) *StoredScriptBuilder {
	rb.v.Lang = lang
	return rb
}

func (rb *StoredScriptBuilder) Options(value map[string]string) *StoredScriptBuilder {
	rb.v.Options = value
	return rb
}

func (rb *StoredScriptBuilder) Source(source string) *StoredScriptBuilder {
	rb.v.Source = source
	return rb
}
