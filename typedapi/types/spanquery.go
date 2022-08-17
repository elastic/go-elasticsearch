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

// SpanQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/span.ts#L79-L91
type SpanQuery struct {
	FieldMaskingSpan *SpanFieldMaskingQuery  `json:"field_masking_span,omitempty"`
	SpanContaining   *SpanContainingQuery    `json:"span_containing,omitempty"`
	SpanFirst        *SpanFirstQuery         `json:"span_first,omitempty"`
	SpanGap          *SpanGapQuery           `json:"span_gap,omitempty"`
	SpanMulti        *SpanMultiTermQuery     `json:"span_multi,omitempty"`
	SpanNear         *SpanNearQuery          `json:"span_near,omitempty"`
	SpanNot          *SpanNotQuery           `json:"span_not,omitempty"`
	SpanOr           *SpanOrQuery            `json:"span_or,omitempty"`
	SpanTerm         map[Field]SpanTermQuery `json:"span_term,omitempty"`
	SpanWithin       *SpanWithinQuery        `json:"span_within,omitempty"`
}

// SpanQueryBuilder holds SpanQuery struct and provides a builder API.
type SpanQueryBuilder struct {
	v *SpanQuery
}

// NewSpanQuery provides a builder for the SpanQuery struct.
func NewSpanQueryBuilder() *SpanQueryBuilder {
	r := SpanQueryBuilder{
		&SpanQuery{
			SpanTerm: make(map[Field]SpanTermQuery, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SpanQuery struct
func (rb *SpanQueryBuilder) Build() SpanQuery {
	return *rb.v
}

func (rb *SpanQueryBuilder) FieldMaskingSpan(fieldmaskingspan *SpanFieldMaskingQueryBuilder) *SpanQueryBuilder {
	v := fieldmaskingspan.Build()
	rb.v.FieldMaskingSpan = &v
	return rb
}

func (rb *SpanQueryBuilder) SpanContaining(spancontaining *SpanContainingQueryBuilder) *SpanQueryBuilder {
	v := spancontaining.Build()
	rb.v.SpanContaining = &v
	return rb
}

func (rb *SpanQueryBuilder) SpanFirst(spanfirst *SpanFirstQueryBuilder) *SpanQueryBuilder {
	v := spanfirst.Build()
	rb.v.SpanFirst = &v
	return rb
}

func (rb *SpanQueryBuilder) SpanGap(spangap *SpanGapQueryBuilder) *SpanQueryBuilder {
	v := spangap.Build()
	rb.v.SpanGap = &v
	return rb
}

func (rb *SpanQueryBuilder) SpanMulti(spanmulti *SpanMultiTermQueryBuilder) *SpanQueryBuilder {
	v := spanmulti.Build()
	rb.v.SpanMulti = &v
	return rb
}

func (rb *SpanQueryBuilder) SpanNear(spannear *SpanNearQueryBuilder) *SpanQueryBuilder {
	v := spannear.Build()
	rb.v.SpanNear = &v
	return rb
}

func (rb *SpanQueryBuilder) SpanNot(spannot *SpanNotQueryBuilder) *SpanQueryBuilder {
	v := spannot.Build()
	rb.v.SpanNot = &v
	return rb
}

func (rb *SpanQueryBuilder) SpanOr(spanor *SpanOrQueryBuilder) *SpanQueryBuilder {
	v := spanor.Build()
	rb.v.SpanOr = &v
	return rb
}

func (rb *SpanQueryBuilder) SpanTerm(values map[Field]*SpanTermQueryBuilder) *SpanQueryBuilder {
	tmp := make(map[Field]SpanTermQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.SpanTerm = tmp
	return rb
}

func (rb *SpanQueryBuilder) SpanWithin(spanwithin *SpanWithinQueryBuilder) *SpanQueryBuilder {
	v := spanwithin.Build()
	rb.v.SpanWithin = &v
	return rb
}
