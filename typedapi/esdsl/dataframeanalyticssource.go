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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dataframeAnalyticsSource struct {
	v *types.DataframeAnalyticsSource
}

func NewDataframeAnalyticsSource() *_dataframeAnalyticsSource {

	return &_dataframeAnalyticsSource{v: types.NewDataframeAnalyticsSource()}

}

// Index or indices on which to perform the analysis. It can be a single index
// or index pattern as well as an array of indices or patterns. NOTE: If your
// source indices contain documents with the same IDs, only the document that is
// indexed last appears in the destination index.
func (s *_dataframeAnalyticsSource) Index(indices ...string) *_dataframeAnalyticsSource {

	s.v.Index = indices

	return s
}

// The Elasticsearch query domain-specific language (DSL). This value
// corresponds to the query object in an Elasticsearch search POST body. All the
// options that are supported by Elasticsearch can be used, as this object is
// passed verbatim to Elasticsearch. By default, this property has the following
// value: {"match_all": {}}.
func (s *_dataframeAnalyticsSource) Query(query types.QueryVariant) *_dataframeAnalyticsSource {

	s.v.Query = query.QueryCaster()

	return s
}

// Definitions of runtime fields that will become part of the mapping of the
// destination index.
func (s *_dataframeAnalyticsSource) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_dataframeAnalyticsSource {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

// Specify `includes` and/or `excludes patterns to select which fields will be
// present in the destination. Fields that are excluded cannot be included in
// the analysis.
func (s *_dataframeAnalyticsSource) Source_(source_ types.DataframeAnalysisAnalyzedFieldsVariant) *_dataframeAnalyticsSource {

	s.v.Source_ = source_.DataframeAnalysisAnalyzedFieldsCaster()

	return s
}

func (s *_dataframeAnalyticsSource) DataframeAnalyticsSourceCaster() *types.DataframeAnalyticsSource {
	return s.v
}
