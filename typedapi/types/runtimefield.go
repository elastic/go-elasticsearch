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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/runtimefieldtype"
)

// RuntimeField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/RuntimeFields.ts#L26-L30
type RuntimeField struct {
	Format *string                           `json:"format,omitempty"`
	Script *Script                           `json:"script,omitempty"`
	Type   runtimefieldtype.RuntimeFieldType `json:"type"`
}

// RuntimeFieldBuilder holds RuntimeField struct and provides a builder API.
type RuntimeFieldBuilder struct {
	v *RuntimeField
}

// NewRuntimeField provides a builder for the RuntimeField struct.
func NewRuntimeFieldBuilder() *RuntimeFieldBuilder {
	r := RuntimeFieldBuilder{
		&RuntimeField{},
	}

	return &r
}

// Build finalize the chain and returns the RuntimeField struct
func (rb *RuntimeFieldBuilder) Build() RuntimeField {
	return *rb.v
}

func (rb *RuntimeFieldBuilder) Format(format string) *RuntimeFieldBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *RuntimeFieldBuilder) Script(script *ScriptBuilder) *RuntimeFieldBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *RuntimeFieldBuilder) Type_(type_ runtimefieldtype.RuntimeFieldType) *RuntimeFieldBuilder {
	rb.v.Type = type_
	return rb
}
