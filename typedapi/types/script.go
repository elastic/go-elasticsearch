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

// Script holds the union for the following types:
//
//	InlineScript
//	StoredScriptId
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Scripting.ts#L56-L57
type Script interface{}

// ScriptBuilder holds Script struct and provides a builder API.
type ScriptBuilder struct {
	v Script
}

// NewScript provides a builder for the Script struct.
func NewScriptBuilder() *ScriptBuilder {
	return &ScriptBuilder{}
}

// Build finalize the chain and returns the Script struct
func (u *ScriptBuilder) Build() Script {
	return u.v
}

func (u *ScriptBuilder) InlineScript(inlinescript *InlineScriptBuilder) *ScriptBuilder {
	v := inlinescript.Build()
	u.v = &v
	return u
}

func (u *ScriptBuilder) StoredScriptId(storedscriptid *StoredScriptIdBuilder) *ScriptBuilder {
	v := storedscriptid.Build()
	u.v = &v
	return u
}
