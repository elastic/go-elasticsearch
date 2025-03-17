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

type _synonymRule struct {
	v *types.SynonymRule
}

func NewSynonymRule() *_synonymRule {

	return &_synonymRule{v: types.NewSynonymRule()}

}

// The identifier for the synonym rule.
// If you do not specify a synonym rule ID when you create a rule, an identifier
// is created automatically by Elasticsearch.
func (s *_synonymRule) Id(id string) *_synonymRule {

	s.v.Id = &id

	return s
}

// The synonyms that conform the synonym rule in Solr format.
func (s *_synonymRule) Synonyms(synonymstring string) *_synonymRule {

	s.v.Synonyms = synonymstring

	return s
}

func (s *_synonymRule) SynonymRuleCaster() *types.SynonymRule {
	return s.v
}
