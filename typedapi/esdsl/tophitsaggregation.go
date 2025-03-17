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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _topHitsAggregation struct {
	v *types.TopHitsAggregation
}

// A metric aggregation that returns the top matching documents per bucket.
func NewTopHitsAggregation() *_topHitsAggregation {

	return &_topHitsAggregation{v: types.NewTopHitsAggregation()}

}

// Fields for which to return doc values.
func (s *_topHitsAggregation) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *_topHitsAggregation {

	for _, v := range docvaluefields {

		s.v.DocvalueFields = append(s.v.DocvalueFields, *v.FieldAndFormatCaster())

	}
	return s
}

// If `true`, returns detailed information about score computation as part of a
// hit.
func (s *_topHitsAggregation) Explain(explain bool) *_topHitsAggregation {

	s.v.Explain = &explain

	return s
}

// The field on which to run the aggregation.
func (s *_topHitsAggregation) Field(field string) *_topHitsAggregation {

	s.v.Field = &field

	return s
}

// Array of wildcard (*) patterns. The request returns values for field names
// matching these patterns in the hits.fields property of the response.
func (s *_topHitsAggregation) Fields(fields ...types.FieldAndFormatVariant) *_topHitsAggregation {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, *v.FieldAndFormatCaster())

	}
	return s
}

// Starting document offset.
func (s *_topHitsAggregation) From(from int) *_topHitsAggregation {

	s.v.From = &from

	return s
}

// Specifies the highlighter to use for retrieving highlighted snippets from one
// or more fields in the search results.
func (s *_topHitsAggregation) Highlight(highlight types.HighlightVariant) *_topHitsAggregation {

	s.v.Highlight = highlight.HighlightCaster()

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_topHitsAggregation) Missing(missing types.MissingVariant) *_topHitsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_topHitsAggregation) Script(script types.ScriptVariant) *_topHitsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

// Returns the result of one or more script evaluations for each hit.
func (s *_topHitsAggregation) ScriptFields(scriptfields map[string]types.ScriptField) *_topHitsAggregation {

	s.v.ScriptFields = scriptfields
	return s
}

func (s *_topHitsAggregation) AddScriptField(key string, value types.ScriptFieldVariant) *_topHitsAggregation {

	var tmp map[string]types.ScriptField
	if s.v.ScriptFields == nil {
		s.v.ScriptFields = make(map[string]types.ScriptField)
	} else {
		tmp = s.v.ScriptFields
	}

	tmp[key] = *value.ScriptFieldCaster()

	s.v.ScriptFields = tmp
	return s
}

// If `true`, returns sequence number and primary term of the last modification
// of each hit.
func (s *_topHitsAggregation) SeqNoPrimaryTerm(seqnoprimaryterm bool) *_topHitsAggregation {

	s.v.SeqNoPrimaryTerm = &seqnoprimaryterm

	return s
}

// The maximum number of top matching hits to return per bucket.
func (s *_topHitsAggregation) Size(size int) *_topHitsAggregation {

	s.v.Size = &size

	return s
}

// Sort order of the top matching hits.
// By default, the hits are sorted by the score of the main query.
func (s *_topHitsAggregation) Sort(sorts ...types.SortCombinationsVariant) *_topHitsAggregation {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

// Selects the fields of the source that are returned.
func (s *_topHitsAggregation) Source_(sourceconfig types.SourceConfigVariant) *_topHitsAggregation {

	s.v.Source_ = *sourceconfig.SourceConfigCaster()

	return s
}

// Returns values for the specified stored fields (fields that use the `store`
// mapping option).
func (s *_topHitsAggregation) StoredFields(fields ...string) *_topHitsAggregation {

	s.v.StoredFields = fields

	return s
}

// If `true`, calculates and returns document scores, even if the scores are not
// used for sorting.
func (s *_topHitsAggregation) TrackScores(trackscores bool) *_topHitsAggregation {

	s.v.TrackScores = &trackscores

	return s
}

// If `true`, returns document version as part of a hit.
func (s *_topHitsAggregation) Version(version bool) *_topHitsAggregation {

	s.v.Version = &version

	return s
}

func (s *_topHitsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.TopHits = s.v

	return container
}

func (s *_topHitsAggregation) TopHitsAggregationCaster() *types.TopHitsAggregation {
	return s.v
}
