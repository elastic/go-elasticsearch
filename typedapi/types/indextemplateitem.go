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

// IndexTemplateItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/get_index_template/IndicesGetIndexTemplateResponse.ts#L29-L32
type IndexTemplateItem struct {
	IndexTemplate IndexTemplate `json:"index_template"`
	Name          Name          `json:"name"`
}

// IndexTemplateItemBuilder holds IndexTemplateItem struct and provides a builder API.
type IndexTemplateItemBuilder struct {
	v *IndexTemplateItem
}

// NewIndexTemplateItem provides a builder for the IndexTemplateItem struct.
func NewIndexTemplateItemBuilder() *IndexTemplateItemBuilder {
	r := IndexTemplateItemBuilder{
		&IndexTemplateItem{},
	}

	return &r
}

// Build finalize the chain and returns the IndexTemplateItem struct
func (rb *IndexTemplateItemBuilder) Build() IndexTemplateItem {
	return *rb.v
}

func (rb *IndexTemplateItemBuilder) IndexTemplate(indextemplate *IndexTemplateBuilder) *IndexTemplateItemBuilder {
	v := indextemplate.Build()
	rb.v.IndexTemplate = v
	return rb
}

func (rb *IndexTemplateItemBuilder) Name(name Name) *IndexTemplateItemBuilder {
	rb.v.Name = name
	return rb
}
