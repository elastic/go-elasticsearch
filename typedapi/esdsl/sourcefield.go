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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sourcefieldmode"
)

type _sourceField struct {
	v *types.SourceField
}

func NewSourceField() *_sourceField {

	return &_sourceField{v: types.NewSourceField()}

}

func (s *_sourceField) Compress(compress bool) *_sourceField {

	s.v.Compress = &compress

	return s
}

func (s *_sourceField) CompressThreshold(compressthreshold string) *_sourceField {

	s.v.CompressThreshold = &compressthreshold

	return s
}

func (s *_sourceField) Enabled(enabled bool) *_sourceField {

	s.v.Enabled = &enabled

	return s
}

func (s *_sourceField) Excludes(excludes ...string) *_sourceField {

	for _, v := range excludes {

		s.v.Excludes = append(s.v.Excludes, v)

	}
	return s
}

func (s *_sourceField) Includes(includes ...string) *_sourceField {

	for _, v := range includes {

		s.v.Includes = append(s.v.Includes, v)

	}
	return s
}

func (s *_sourceField) Mode(mode sourcefieldmode.SourceFieldMode) *_sourceField {

	s.v.Mode = &mode
	return s
}

func (s *_sourceField) SourceFieldCaster() *types.SourceField {
	return s.v
}
