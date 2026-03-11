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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexTemplate struct {
	v *types.IndexTemplate
}

func NewIndexTemplate() *_indexTemplate {

	return &_indexTemplate{v: types.NewIndexTemplate()}

}

func (s *_indexTemplate) AllowAutoCreate(allowautocreate bool) *_indexTemplate {

	s.v.AllowAutoCreate = &allowautocreate

	return s
}

func (s *_indexTemplate) ComposedOf(composedofs ...string) *_indexTemplate {

	for _, v := range composedofs {

		s.v.ComposedOf = append(s.v.ComposedOf, v)

	}
	return s
}

func (s *_indexTemplate) CreatedDate(datetime types.DateTimeVariant) *_indexTemplate {

	s.v.CreatedDate = *datetime.DateTimeCaster()

	return s
}

func (s *_indexTemplate) CreatedDateMillis(epochtimeunitmillis int64) *_indexTemplate {

	s.v.CreatedDateMillis = &epochtimeunitmillis

	return s
}

func (s *_indexTemplate) DataStream(datastream types.IndexTemplateDataStreamConfigurationVariant) *_indexTemplate {

	s.v.DataStream = datastream.IndexTemplateDataStreamConfigurationCaster()

	return s
}

func (s *_indexTemplate) Deprecated(deprecated bool) *_indexTemplate {

	s.v.Deprecated = &deprecated

	return s
}

func (s *_indexTemplate) IgnoreMissingComponentTemplates(names ...string) *_indexTemplate {

	s.v.IgnoreMissingComponentTemplates = names

	return s
}

func (s *_indexTemplate) IndexPatterns(names ...string) *_indexTemplate {

	s.v.IndexPatterns = names

	return s
}

func (s *_indexTemplate) Meta_(metadata types.MetadataVariant) *_indexTemplate {

	s.v.Meta_ = *metadata.MetadataCaster()

	return s
}

func (s *_indexTemplate) ModifiedDate(datetime types.DateTimeVariant) *_indexTemplate {

	s.v.ModifiedDate = *datetime.DateTimeCaster()

	return s
}

func (s *_indexTemplate) ModifiedDateMillis(epochtimeunitmillis int64) *_indexTemplate {

	s.v.ModifiedDateMillis = &epochtimeunitmillis

	return s
}

func (s *_indexTemplate) Priority(priority int64) *_indexTemplate {

	s.v.Priority = &priority

	return s
}

func (s *_indexTemplate) Template(template types.IndexTemplateSummaryVariant) *_indexTemplate {

	s.v.Template = template.IndexTemplateSummaryCaster()

	return s
}

func (s *_indexTemplate) Version(versionnumber int64) *_indexTemplate {

	s.v.Version = &versionnumber

	return s
}

func (s *_indexTemplate) IndexTemplateCaster() *types.IndexTemplate {
	return s.v
}
