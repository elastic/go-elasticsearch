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

// RuntimeFieldsType type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L262-L277
type RuntimeFieldsType struct {
	CharsMax        int64    `json:"chars_max"`
	CharsTotal      int64    `json:"chars_total"`
	Count           int64    `json:"count"`
	DocMax          int64    `json:"doc_max"`
	DocTotal        int64    `json:"doc_total"`
	IndexCount      int64    `json:"index_count"`
	Lang            []string `json:"lang"`
	LinesMax        int64    `json:"lines_max"`
	LinesTotal      int64    `json:"lines_total"`
	Name            Field    `json:"name"`
	ScriptlessCount int64    `json:"scriptless_count"`
	ShadowedCount   int64    `json:"shadowed_count"`
	SourceMax       int64    `json:"source_max"`
	SourceTotal     int64    `json:"source_total"`
}

// RuntimeFieldsTypeBuilder holds RuntimeFieldsType struct and provides a builder API.
type RuntimeFieldsTypeBuilder struct {
	v *RuntimeFieldsType
}

// NewRuntimeFieldsType provides a builder for the RuntimeFieldsType struct.
func NewRuntimeFieldsTypeBuilder() *RuntimeFieldsTypeBuilder {
	r := RuntimeFieldsTypeBuilder{
		&RuntimeFieldsType{},
	}

	return &r
}

// Build finalize the chain and returns the RuntimeFieldsType struct
func (rb *RuntimeFieldsTypeBuilder) Build() RuntimeFieldsType {
	return *rb.v
}

func (rb *RuntimeFieldsTypeBuilder) CharsMax(charsmax int64) *RuntimeFieldsTypeBuilder {
	rb.v.CharsMax = charsmax
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) CharsTotal(charstotal int64) *RuntimeFieldsTypeBuilder {
	rb.v.CharsTotal = charstotal
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) Count(count int64) *RuntimeFieldsTypeBuilder {
	rb.v.Count = count
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) DocMax(docmax int64) *RuntimeFieldsTypeBuilder {
	rb.v.DocMax = docmax
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) DocTotal(doctotal int64) *RuntimeFieldsTypeBuilder {
	rb.v.DocTotal = doctotal
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) IndexCount(indexcount int64) *RuntimeFieldsTypeBuilder {
	rb.v.IndexCount = indexcount
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) Lang(lang ...string) *RuntimeFieldsTypeBuilder {
	rb.v.Lang = lang
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) LinesMax(linesmax int64) *RuntimeFieldsTypeBuilder {
	rb.v.LinesMax = linesmax
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) LinesTotal(linestotal int64) *RuntimeFieldsTypeBuilder {
	rb.v.LinesTotal = linestotal
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) Name(name Field) *RuntimeFieldsTypeBuilder {
	rb.v.Name = name
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) ScriptlessCount(scriptlesscount int64) *RuntimeFieldsTypeBuilder {
	rb.v.ScriptlessCount = scriptlesscount
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) ShadowedCount(shadowedcount int64) *RuntimeFieldsTypeBuilder {
	rb.v.ShadowedCount = shadowedcount
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) SourceMax(sourcemax int64) *RuntimeFieldsTypeBuilder {
	rb.v.SourceMax = sourcemax
	return rb
}

func (rb *RuntimeFieldsTypeBuilder) SourceTotal(sourcetotal int64) *RuntimeFieldsTypeBuilder {
	rb.v.SourceTotal = sourcetotal
	return rb
}
