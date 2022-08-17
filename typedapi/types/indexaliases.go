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

// IndexAliases type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/get_alias/IndicesGetAliasResponse.ts#L36-L38
type IndexAliases struct {
	Aliases map[string]AliasDefinition `json:"aliases"`
}

// IndexAliasesBuilder holds IndexAliases struct and provides a builder API.
type IndexAliasesBuilder struct {
	v *IndexAliases
}

// NewIndexAliases provides a builder for the IndexAliases struct.
func NewIndexAliasesBuilder() *IndexAliasesBuilder {
	r := IndexAliasesBuilder{
		&IndexAliases{
			Aliases: make(map[string]AliasDefinition, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexAliases struct
func (rb *IndexAliasesBuilder) Build() IndexAliases {
	return *rb.v
}

func (rb *IndexAliasesBuilder) Aliases(values map[string]*AliasDefinitionBuilder) *IndexAliasesBuilder {
	tmp := make(map[string]AliasDefinition, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aliases = tmp
	return rb
}
