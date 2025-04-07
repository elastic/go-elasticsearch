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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _keywordProperty struct {
	v *types.KeywordProperty
}

func NewKeywordProperty() *_keywordProperty {

	return &_keywordProperty{v: types.NewKeywordProperty()}

}

func (s *_keywordProperty) Boost(boost types.Float64) *_keywordProperty {

	s.v.Boost = &boost

	return s
}

func (s *_keywordProperty) CopyTo(fields ...string) *_keywordProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_keywordProperty) DocValues(docvalues bool) *_keywordProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_keywordProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_keywordProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_keywordProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_keywordProperty {

	s.v.EagerGlobalOrdinals = &eagerglobalordinals

	return s
}

func (s *_keywordProperty) Fields(fields map[string]types.Property) *_keywordProperty {

	s.v.Fields = fields
	return s
}

func (s *_keywordProperty) AddField(key string, value types.PropertyVariant) *_keywordProperty {

	var tmp map[string]types.Property
	if s.v.Fields == nil {
		s.v.Fields = make(map[string]types.Property)
	} else {
		tmp = s.v.Fields
	}

	tmp[key] = *value.PropertyCaster()

	s.v.Fields = tmp
	return s
}

func (s *_keywordProperty) IgnoreAbove(ignoreabove int) *_keywordProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_keywordProperty) Index(index bool) *_keywordProperty {

	s.v.Index = &index

	return s
}

func (s *_keywordProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_keywordProperty {

	s.v.IndexOptions = &indexoptions
	return s
}

func (s *_keywordProperty) Meta(meta map[string]string) *_keywordProperty {

	s.v.Meta = meta
	return s
}

func (s *_keywordProperty) AddMeta(key string, value string) *_keywordProperty {

	var tmp map[string]string
	if s.v.Meta == nil {
		s.v.Meta = make(map[string]string)
	} else {
		tmp = s.v.Meta
	}

	tmp[key] = value

	s.v.Meta = tmp
	return s
}

func (s *_keywordProperty) Normalizer(normalizer string) *_keywordProperty {

	s.v.Normalizer = &normalizer

	return s
}

func (s *_keywordProperty) Norms(norms bool) *_keywordProperty {

	s.v.Norms = &norms

	return s
}

func (s *_keywordProperty) NullValue(nullvalue string) *_keywordProperty {

	s.v.NullValue = &nullvalue

	return s
}

func (s *_keywordProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_keywordProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_keywordProperty) Properties(properties map[string]types.Property) *_keywordProperty {

	s.v.Properties = properties
	return s
}

func (s *_keywordProperty) AddProperty(key string, value types.PropertyVariant) *_keywordProperty {

	var tmp map[string]types.Property
	if s.v.Properties == nil {
		s.v.Properties = make(map[string]types.Property)
	} else {
		tmp = s.v.Properties
	}

	tmp[key] = *value.PropertyCaster()

	s.v.Properties = tmp
	return s
}

func (s *_keywordProperty) Script(script types.ScriptVariant) *_keywordProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_keywordProperty) Similarity(similarity string) *_keywordProperty {

	s.v.Similarity = &similarity

	return s
}

func (s *_keywordProperty) SplitQueriesOnWhitespace(splitqueriesonwhitespace bool) *_keywordProperty {

	s.v.SplitQueriesOnWhitespace = &splitqueriesonwhitespace

	return s
}

func (s *_keywordProperty) Store(store bool) *_keywordProperty {

	s.v.Store = &store

	return s
}

func (s *_keywordProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_keywordProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_keywordProperty) TimeSeriesDimension(timeseriesdimension bool) *_keywordProperty {

	s.v.TimeSeriesDimension = &timeseriesdimension

	return s
}

func (s *_keywordProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_keywordProperty) KeywordPropertyCaster() *types.KeywordProperty {
	return s.v
}
