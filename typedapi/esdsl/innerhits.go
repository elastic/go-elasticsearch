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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _innerHits struct {
	v *types.InnerHits
}

func NewInnerHits() *_innerHits {

	return &_innerHits{v: types.NewInnerHits()}

}

func (s *_innerHits) Collapse(collapse types.FieldCollapseVariant) *_innerHits {

	s.v.Collapse = collapse.FieldCollapseCaster()

	return s
}

func (s *_innerHits) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *_innerHits {

	for _, v := range docvaluefields {

		s.v.DocvalueFields = append(s.v.DocvalueFields, *v.FieldAndFormatCaster())

	}
	return s
}

func (s *_innerHits) Explain(explain bool) *_innerHits {

	s.v.Explain = &explain

	return s
}

func (s *_innerHits) Fields(fields ...string) *_innerHits {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

func (s *_innerHits) From(from int) *_innerHits {

	s.v.From = &from

	return s
}

func (s *_innerHits) Highlight(highlight types.HighlightVariant) *_innerHits {

	s.v.Highlight = highlight.HighlightCaster()

	return s
}

func (s *_innerHits) IgnoreUnmapped(ignoreunmapped bool) *_innerHits {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_innerHits) Name(name string) *_innerHits {

	s.v.Name = &name

	return s
}

func (s *_innerHits) ScriptFields(scriptfields map[string]types.ScriptField) *_innerHits {

	s.v.ScriptFields = scriptfields
	return s
}

func (s *_innerHits) AddScriptField(key string, value types.ScriptFieldVariant) *_innerHits {

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

func (s *_innerHits) SeqNoPrimaryTerm(seqnoprimaryterm bool) *_innerHits {

	s.v.SeqNoPrimaryTerm = &seqnoprimaryterm

	return s
}

func (s *_innerHits) Size(size int) *_innerHits {

	s.v.Size = &size

	return s
}

func (s *_innerHits) Sort(sorts ...types.SortCombinationsVariant) *_innerHits {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

func (s *_innerHits) Source_(sourceconfig types.SourceConfigVariant) *_innerHits {

	s.v.Source_ = *sourceconfig.SourceConfigCaster()

	return s
}

func (s *_innerHits) StoredFields(fields ...string) *_innerHits {

	s.v.StoredFields = fields

	return s
}

func (s *_innerHits) TrackScores(trackscores bool) *_innerHits {

	s.v.TrackScores = &trackscores

	return s
}

func (s *_innerHits) Version(version bool) *_innerHits {

	s.v.Version = &version

	return s
}

func (s *_innerHits) InnerHitsCaster() *types.InnerHits {
	return s.v
}
