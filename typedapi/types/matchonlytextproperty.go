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

// MatchOnlyTextProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L197-L222
type MatchOnlyTextProperty struct {
	// CopyTo Allows you to copy the values of multiple fields into a group
	// field, which can then be queried as a single field.
	CopyTo *Fields `json:"copy_to,omitempty"`
	// Fields Multi-fields allow the same string value to be indexed in multiple ways for
	// different purposes, such as one
	// field for search and a multi-field for sorting and aggregations, or the same
	// string value analyzed by different analyzers.
	Fields map[PropertyName]Property `json:"fields,omitempty"`
	// Meta Metadata about the field.
	Meta map[string]string `json:"meta,omitempty"`
	Type string            `json:"type,omitempty"`
}

// MatchOnlyTextPropertyBuilder holds MatchOnlyTextProperty struct and provides a builder API.
type MatchOnlyTextPropertyBuilder struct {
	v *MatchOnlyTextProperty
}

// NewMatchOnlyTextProperty provides a builder for the MatchOnlyTextProperty struct.
func NewMatchOnlyTextPropertyBuilder() *MatchOnlyTextPropertyBuilder {
	r := MatchOnlyTextPropertyBuilder{
		&MatchOnlyTextProperty{
			Fields: make(map[PropertyName]Property, 0),
			Meta:   make(map[string]string, 0),
		},
	}

	r.v.Type = "match_only_text"

	return &r
}

// Build finalize the chain and returns the MatchOnlyTextProperty struct
func (rb *MatchOnlyTextPropertyBuilder) Build() MatchOnlyTextProperty {
	return *rb.v
}

// CopyTo Allows you to copy the values of multiple fields into a group
// field, which can then be queried as a single field.

func (rb *MatchOnlyTextPropertyBuilder) CopyTo(copyto *FieldsBuilder) *MatchOnlyTextPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

// Fields Multi-fields allow the same string value to be indexed in multiple ways for
// different purposes, such as one
// field for search and a multi-field for sorting and aggregations, or the same
// string value analyzed by different analyzers.

func (rb *MatchOnlyTextPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *MatchOnlyTextPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

// Meta Metadata about the field.

func (rb *MatchOnlyTextPropertyBuilder) Meta(value map[string]string) *MatchOnlyTextPropertyBuilder {
	rb.v.Meta = value
	return rb
}
