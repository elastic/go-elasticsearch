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

// Scripting type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L383-L388
type Scripting struct {
	CacheEvictions            *int64    `json:"cache_evictions,omitempty"`
	CompilationLimitTriggered *int64    `json:"compilation_limit_triggered,omitempty"`
	Compilations              *int64    `json:"compilations,omitempty"`
	Contexts                  []Context `json:"contexts,omitempty"`
}

// ScriptingBuilder holds Scripting struct and provides a builder API.
type ScriptingBuilder struct {
	v *Scripting
}

// NewScripting provides a builder for the Scripting struct.
func NewScriptingBuilder() *ScriptingBuilder {
	r := ScriptingBuilder{
		&Scripting{},
	}

	return &r
}

// Build finalize the chain and returns the Scripting struct
func (rb *ScriptingBuilder) Build() Scripting {
	return *rb.v
}

func (rb *ScriptingBuilder) CacheEvictions(cacheevictions int64) *ScriptingBuilder {
	rb.v.CacheEvictions = &cacheevictions
	return rb
}

func (rb *ScriptingBuilder) CompilationLimitTriggered(compilationlimittriggered int64) *ScriptingBuilder {
	rb.v.CompilationLimitTriggered = &compilationlimittriggered
	return rb
}

func (rb *ScriptingBuilder) Compilations(compilations int64) *ScriptingBuilder {
	rb.v.Compilations = &compilations
	return rb
}

func (rb *ScriptingBuilder) Contexts(contexts []ContextBuilder) *ScriptingBuilder {
	tmp := make([]Context, len(contexts))
	for _, value := range contexts {
		tmp = append(tmp, value.Build())
	}
	rb.v.Contexts = tmp
	return rb
}
