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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptsorttype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// ScriptSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/sort.ts#L67-L73
type ScriptSort struct {
	Mode   *sortmode.SortMode             `json:"mode,omitempty"`
	Nested *NestedSortValue               `json:"nested,omitempty"`
	Order  *sortorder.SortOrder           `json:"order,omitempty"`
	Script Script                         `json:"script"`
	Type   *scriptsorttype.ScriptSortType `json:"type,omitempty"`
}

// ScriptSortBuilder holds ScriptSort struct and provides a builder API.
type ScriptSortBuilder struct {
	v *ScriptSort
}

// NewScriptSort provides a builder for the ScriptSort struct.
func NewScriptSortBuilder() *ScriptSortBuilder {
	r := ScriptSortBuilder{
		&ScriptSort{},
	}

	return &r
}

// Build finalize the chain and returns the ScriptSort struct
func (rb *ScriptSortBuilder) Build() ScriptSort {
	return *rb.v
}

func (rb *ScriptSortBuilder) Mode(mode sortmode.SortMode) *ScriptSortBuilder {
	rb.v.Mode = &mode
	return rb
}

func (rb *ScriptSortBuilder) Nested(nested *NestedSortValueBuilder) *ScriptSortBuilder {
	v := nested.Build()
	rb.v.Nested = &v
	return rb
}

func (rb *ScriptSortBuilder) Order(order sortorder.SortOrder) *ScriptSortBuilder {
	rb.v.Order = &order
	return rb
}

func (rb *ScriptSortBuilder) Script(script *ScriptBuilder) *ScriptSortBuilder {
	v := script.Build()
	rb.v.Script = v
	return rb
}

func (rb *ScriptSortBuilder) Type_(type_ scriptsorttype.ScriptSortType) *ScriptSortBuilder {
	rb.v.Type = &type_
	return rb
}
