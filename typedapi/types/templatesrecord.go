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

// TemplatesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/templates/types.ts#L22-L48
type TemplatesRecord struct {
	// ComposedOf component templates comprising index template
	ComposedOf *string `json:"composed_of,omitempty"`
	// IndexPatterns template index patterns
	IndexPatterns *string `json:"index_patterns,omitempty"`
	// Name template name
	Name *Name `json:"name,omitempty"`
	// Order template application order/priority number
	Order *string `json:"order,omitempty"`
	// Version version
	Version VersionString `json:"version,omitempty"`
}

// TemplatesRecordBuilder holds TemplatesRecord struct and provides a builder API.
type TemplatesRecordBuilder struct {
	v *TemplatesRecord
}

// NewTemplatesRecord provides a builder for the TemplatesRecord struct.
func NewTemplatesRecordBuilder() *TemplatesRecordBuilder {
	r := TemplatesRecordBuilder{
		&TemplatesRecord{},
	}

	return &r
}

// Build finalize the chain and returns the TemplatesRecord struct
func (rb *TemplatesRecordBuilder) Build() TemplatesRecord {
	return *rb.v
}

// ComposedOf component templates comprising index template

func (rb *TemplatesRecordBuilder) ComposedOf(composedof string) *TemplatesRecordBuilder {
	rb.v.ComposedOf = &composedof
	return rb
}

// IndexPatterns template index patterns

func (rb *TemplatesRecordBuilder) IndexPatterns(indexpatterns string) *TemplatesRecordBuilder {
	rb.v.IndexPatterns = &indexpatterns
	return rb
}

// Name template name

func (rb *TemplatesRecordBuilder) Name(name Name) *TemplatesRecordBuilder {
	rb.v.Name = &name
	return rb
}

// Order template application order/priority number

func (rb *TemplatesRecordBuilder) Order(order string) *TemplatesRecordBuilder {
	rb.v.Order = &order
	return rb
}

// Version version

func (rb *TemplatesRecordBuilder) Version(version VersionString) *TemplatesRecordBuilder {
	rb.v.Version = version
	return rb
}
