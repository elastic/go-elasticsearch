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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _likeDocument struct {
	v *types.LikeDocument
}

func NewLikeDocument() *_likeDocument {

	return &_likeDocument{v: types.NewLikeDocument()}

}

func (s *_likeDocument) Doc(doc json.RawMessage) *_likeDocument {

	s.v.Doc = doc

	return s
}

func (s *_likeDocument) Fields(fields ...string) *_likeDocument {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

func (s *_likeDocument) Id_(id string) *_likeDocument {

	s.v.Id_ = &id

	return s
}

func (s *_likeDocument) Index_(indexname string) *_likeDocument {

	s.v.Index_ = &indexname

	return s
}

func (s *_likeDocument) PerFieldAnalyzer(perfieldanalyzer map[string]string) *_likeDocument {

	s.v.PerFieldAnalyzer = perfieldanalyzer
	return s
}

func (s *_likeDocument) AddPerFieldAnalyzer(key string, value string) *_likeDocument {

	var tmp map[string]string
	if s.v.PerFieldAnalyzer == nil {
		s.v.PerFieldAnalyzer = make(map[string]string)
	} else {
		tmp = s.v.PerFieldAnalyzer
	}

	tmp[key] = value

	s.v.PerFieldAnalyzer = tmp
	return s
}

func (s *_likeDocument) Routing(routing string) *_likeDocument {

	s.v.Routing = &routing

	return s
}

func (s *_likeDocument) Version(versionnumber int64) *_likeDocument {

	s.v.Version = &versionnumber

	return s
}

func (s *_likeDocument) VersionType(versiontype versiontype.VersionType) *_likeDocument {

	s.v.VersionType = &versiontype
	return s
}

func (s *_likeDocument) LikeDocumentCaster() *types.LikeDocument {
	return s.v
}
