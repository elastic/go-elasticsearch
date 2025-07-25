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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termvectoroption"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _dynamicProperty struct {
	v *types.DynamicProperty
}

func NewDynamicProperty() *_dynamicProperty {

	return &_dynamicProperty{v: types.NewDynamicProperty()}

}

func (s *_dynamicProperty) Analyzer(analyzer string) *_dynamicProperty {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_dynamicProperty) Boost(boost types.Float64) *_dynamicProperty {

	s.v.Boost = &boost

	return s
}

func (s *_dynamicProperty) Coerce(coerce bool) *_dynamicProperty {

	s.v.Coerce = &coerce

	return s
}

func (s *_dynamicProperty) CopyTo(fields ...string) *_dynamicProperty {

	s.v.CopyTo = fields

	return s
}

func (s *_dynamicProperty) DocValues(docvalues bool) *_dynamicProperty {

	s.v.DocValues = &docvalues

	return s
}

func (s *_dynamicProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_dynamicProperty {

	s.v.Dynamic = &dynamic
	return s
}

func (s *_dynamicProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_dynamicProperty {

	s.v.EagerGlobalOrdinals = &eagerglobalordinals

	return s
}

func (s *_dynamicProperty) Enabled(enabled bool) *_dynamicProperty {

	s.v.Enabled = &enabled

	return s
}

func (s *_dynamicProperty) Fields(fields map[string]types.Property) *_dynamicProperty {

	s.v.Fields = fields
	return s
}

func (s *_dynamicProperty) AddField(key string, value types.PropertyVariant) *_dynamicProperty {

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

func (s *_dynamicProperty) Format(format string) *_dynamicProperty {

	s.v.Format = &format

	return s
}

func (s *_dynamicProperty) IgnoreAbove(ignoreabove int) *_dynamicProperty {

	s.v.IgnoreAbove = &ignoreabove

	return s
}

func (s *_dynamicProperty) IgnoreMalformed(ignoremalformed bool) *_dynamicProperty {

	s.v.IgnoreMalformed = &ignoremalformed

	return s
}

func (s *_dynamicProperty) Index(index bool) *_dynamicProperty {

	s.v.Index = &index

	return s
}

func (s *_dynamicProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_dynamicProperty {

	s.v.IndexOptions = &indexoptions
	return s
}

func (s *_dynamicProperty) IndexPhrases(indexphrases bool) *_dynamicProperty {

	s.v.IndexPhrases = &indexphrases

	return s
}

func (s *_dynamicProperty) IndexPrefixes(indexprefixes types.TextIndexPrefixesVariant) *_dynamicProperty {

	s.v.IndexPrefixes = indexprefixes.TextIndexPrefixesCaster()

	return s
}

func (s *_dynamicProperty) Locale(locale string) *_dynamicProperty {

	s.v.Locale = &locale

	return s
}

func (s *_dynamicProperty) Meta(meta map[string]string) *_dynamicProperty {

	s.v.Meta = meta
	return s
}

func (s *_dynamicProperty) AddMeta(key string, value string) *_dynamicProperty {

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

func (s *_dynamicProperty) Norms(norms bool) *_dynamicProperty {

	s.v.Norms = &norms

	return s
}

func (s *_dynamicProperty) NullValue(fieldvalue types.FieldValueVariant) *_dynamicProperty {

	s.v.NullValue = *fieldvalue.FieldValueCaster()

	return s
}

func (s *_dynamicProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_dynamicProperty {

	s.v.OnScriptError = &onscripterror
	return s
}

func (s *_dynamicProperty) PositionIncrementGap(positionincrementgap int) *_dynamicProperty {

	s.v.PositionIncrementGap = &positionincrementgap

	return s
}

func (s *_dynamicProperty) PrecisionStep(precisionstep int) *_dynamicProperty {

	s.v.PrecisionStep = &precisionstep

	return s
}

func (s *_dynamicProperty) Properties(properties map[string]types.Property) *_dynamicProperty {

	s.v.Properties = properties
	return s
}

func (s *_dynamicProperty) AddProperty(key string, value types.PropertyVariant) *_dynamicProperty {

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

func (s *_dynamicProperty) Script(script types.ScriptVariant) *_dynamicProperty {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_dynamicProperty) SearchAnalyzer(searchanalyzer string) *_dynamicProperty {

	s.v.SearchAnalyzer = &searchanalyzer

	return s
}

func (s *_dynamicProperty) SearchQuoteAnalyzer(searchquoteanalyzer string) *_dynamicProperty {

	s.v.SearchQuoteAnalyzer = &searchquoteanalyzer

	return s
}

func (s *_dynamicProperty) Store(store bool) *_dynamicProperty {

	s.v.Store = &store

	return s
}

func (s *_dynamicProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_dynamicProperty {

	s.v.SyntheticSourceKeep = &syntheticsourcekeep
	return s
}

func (s *_dynamicProperty) TermVector(termvector termvectoroption.TermVectorOption) *_dynamicProperty {

	s.v.TermVector = &termvector
	return s
}

func (s *_dynamicProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_dynamicProperty {

	s.v.TimeSeriesMetric = &timeseriesmetric
	return s
}

func (s *_dynamicProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_dynamicProperty) DynamicPropertyCaster() *types.DynamicProperty {
	return s.v
}
