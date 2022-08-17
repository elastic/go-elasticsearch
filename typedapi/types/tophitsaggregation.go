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

// TopHitsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L171-L184
type TopHitsAggregation struct {
	DocvalueFields   *Fields                `json:"docvalue_fields,omitempty"`
	Explain          *bool                  `json:"explain,omitempty"`
	Field            *Field                 `json:"field,omitempty"`
	From             *int                   `json:"from,omitempty"`
	Highlight        *Highlight             `json:"highlight,omitempty"`
	Missing          *Missing               `json:"missing,omitempty"`
	Script           *Script                `json:"script,omitempty"`
	ScriptFields     map[string]ScriptField `json:"script_fields,omitempty"`
	SeqNoPrimaryTerm *bool                  `json:"seq_no_primary_term,omitempty"`
	Size             *int                   `json:"size,omitempty"`
	Sort             *Sort                  `json:"sort,omitempty"`
	Source_          *SourceConfig          `json:"_source,omitempty"`
	StoredFields     *Fields                `json:"stored_fields,omitempty"`
	TrackScores      *bool                  `json:"track_scores,omitempty"`
	Version          *bool                  `json:"version,omitempty"`
}

// TopHitsAggregationBuilder holds TopHitsAggregation struct and provides a builder API.
type TopHitsAggregationBuilder struct {
	v *TopHitsAggregation
}

// NewTopHitsAggregation provides a builder for the TopHitsAggregation struct.
func NewTopHitsAggregationBuilder() *TopHitsAggregationBuilder {
	r := TopHitsAggregationBuilder{
		&TopHitsAggregation{
			ScriptFields: make(map[string]ScriptField, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TopHitsAggregation struct
func (rb *TopHitsAggregationBuilder) Build() TopHitsAggregation {
	return *rb.v
}

func (rb *TopHitsAggregationBuilder) DocvalueFields(docvaluefields *FieldsBuilder) *TopHitsAggregationBuilder {
	v := docvaluefields.Build()
	rb.v.DocvalueFields = &v
	return rb
}

func (rb *TopHitsAggregationBuilder) Explain(explain bool) *TopHitsAggregationBuilder {
	rb.v.Explain = &explain
	return rb
}

func (rb *TopHitsAggregationBuilder) Field(field Field) *TopHitsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *TopHitsAggregationBuilder) From(from int) *TopHitsAggregationBuilder {
	rb.v.From = &from
	return rb
}

func (rb *TopHitsAggregationBuilder) Highlight(highlight *HighlightBuilder) *TopHitsAggregationBuilder {
	v := highlight.Build()
	rb.v.Highlight = &v
	return rb
}

func (rb *TopHitsAggregationBuilder) Missing(missing *MissingBuilder) *TopHitsAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *TopHitsAggregationBuilder) Script(script *ScriptBuilder) *TopHitsAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *TopHitsAggregationBuilder) ScriptFields(values map[string]*ScriptFieldBuilder) *TopHitsAggregationBuilder {
	tmp := make(map[string]ScriptField, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ScriptFields = tmp
	return rb
}

func (rb *TopHitsAggregationBuilder) SeqNoPrimaryTerm(seqnoprimaryterm bool) *TopHitsAggregationBuilder {
	rb.v.SeqNoPrimaryTerm = &seqnoprimaryterm
	return rb
}

func (rb *TopHitsAggregationBuilder) Size(size int) *TopHitsAggregationBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *TopHitsAggregationBuilder) Sort(sort *SortBuilder) *TopHitsAggregationBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}

func (rb *TopHitsAggregationBuilder) Source_(source_ *SourceConfigBuilder) *TopHitsAggregationBuilder {
	v := source_.Build()
	rb.v.Source_ = &v
	return rb
}

func (rb *TopHitsAggregationBuilder) StoredFields(storedfields *FieldsBuilder) *TopHitsAggregationBuilder {
	v := storedfields.Build()
	rb.v.StoredFields = &v
	return rb
}

func (rb *TopHitsAggregationBuilder) TrackScores(trackscores bool) *TopHitsAggregationBuilder {
	rb.v.TrackScores = &trackscores
	return rb
}

func (rb *TopHitsAggregationBuilder) Version(version bool) *TopHitsAggregationBuilder {
	rb.v.Version = &version
	return rb
}
