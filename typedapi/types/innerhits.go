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

// InnerHits type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/hits.ts#L102-L120
type InnerHits struct {
	Collapse         *FieldCollapse        `json:"collapse,omitempty"`
	DocvalueFields   []FieldAndFormat      `json:"docvalue_fields,omitempty"`
	Explain          *bool                 `json:"explain,omitempty"`
	Fields           *Fields               `json:"fields,omitempty"`
	From             *int                  `json:"from,omitempty"`
	Highlight        *Highlight            `json:"highlight,omitempty"`
	IgnoreUnmapped   *bool                 `json:"ignore_unmapped,omitempty"`
	Name             *Name                 `json:"name,omitempty"`
	ScriptFields     map[Field]ScriptField `json:"script_fields,omitempty"`
	SeqNoPrimaryTerm *bool                 `json:"seq_no_primary_term,omitempty"`
	Size             *int                  `json:"size,omitempty"`
	Sort             *Sort                 `json:"sort,omitempty"`
	Source_          *SourceConfig         `json:"_source,omitempty"`
	StoredField      *Fields               `json:"stored_field,omitempty"`
	TrackScores      *bool                 `json:"track_scores,omitempty"`
	Version          *bool                 `json:"version,omitempty"`
}

// InnerHitsBuilder holds InnerHits struct and provides a builder API.
type InnerHitsBuilder struct {
	v *InnerHits
}

// NewInnerHits provides a builder for the InnerHits struct.
func NewInnerHitsBuilder() *InnerHitsBuilder {
	r := InnerHitsBuilder{
		&InnerHits{
			ScriptFields: make(map[Field]ScriptField, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InnerHits struct
func (rb *InnerHitsBuilder) Build() InnerHits {
	return *rb.v
}

func (rb *InnerHitsBuilder) Collapse(collapse *FieldCollapseBuilder) *InnerHitsBuilder {
	v := collapse.Build()
	rb.v.Collapse = &v
	return rb
}

func (rb *InnerHitsBuilder) DocvalueFields(docvalue_fields []FieldAndFormatBuilder) *InnerHitsBuilder {
	tmp := make([]FieldAndFormat, len(docvalue_fields))
	for _, value := range docvalue_fields {
		tmp = append(tmp, value.Build())
	}
	rb.v.DocvalueFields = tmp
	return rb
}

func (rb *InnerHitsBuilder) Explain(explain bool) *InnerHitsBuilder {
	rb.v.Explain = &explain
	return rb
}

func (rb *InnerHitsBuilder) Fields(fields *FieldsBuilder) *InnerHitsBuilder {
	v := fields.Build()
	rb.v.Fields = &v
	return rb
}

func (rb *InnerHitsBuilder) From(from int) *InnerHitsBuilder {
	rb.v.From = &from
	return rb
}

func (rb *InnerHitsBuilder) Highlight(highlight *HighlightBuilder) *InnerHitsBuilder {
	v := highlight.Build()
	rb.v.Highlight = &v
	return rb
}

func (rb *InnerHitsBuilder) IgnoreUnmapped(ignoreunmapped bool) *InnerHitsBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

func (rb *InnerHitsBuilder) Name(name Name) *InnerHitsBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *InnerHitsBuilder) ScriptFields(values map[Field]*ScriptFieldBuilder) *InnerHitsBuilder {
	tmp := make(map[Field]ScriptField, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ScriptFields = tmp
	return rb
}

func (rb *InnerHitsBuilder) SeqNoPrimaryTerm(seqnoprimaryterm bool) *InnerHitsBuilder {
	rb.v.SeqNoPrimaryTerm = &seqnoprimaryterm
	return rb
}

func (rb *InnerHitsBuilder) Size(size int) *InnerHitsBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *InnerHitsBuilder) Sort(sort *SortBuilder) *InnerHitsBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}

func (rb *InnerHitsBuilder) Source_(source_ *SourceConfigBuilder) *InnerHitsBuilder {
	v := source_.Build()
	rb.v.Source_ = &v
	return rb
}

func (rb *InnerHitsBuilder) StoredField(storedfield *FieldsBuilder) *InnerHitsBuilder {
	v := storedfield.Build()
	rb.v.StoredField = &v
	return rb
}

func (rb *InnerHitsBuilder) TrackScores(trackscores bool) *InnerHitsBuilder {
	rb.v.TrackScores = &trackscores
	return rb
}

func (rb *InnerHitsBuilder) Version(version bool) *InnerHitsBuilder {
	rb.v.Version = &version
	return rb
}
