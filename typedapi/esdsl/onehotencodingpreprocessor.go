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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _oneHotEncodingPreprocessor struct {
	v *types.OneHotEncodingPreprocessor
}

func NewOneHotEncodingPreprocessor(field string) *_oneHotEncodingPreprocessor {

	tmp := &_oneHotEncodingPreprocessor{v: types.NewOneHotEncodingPreprocessor()}

	tmp.Field(field)

	return tmp

}

func (s *_oneHotEncodingPreprocessor) Field(field string) *_oneHotEncodingPreprocessor {

	s.v.Field = field

	return s
}

func (s *_oneHotEncodingPreprocessor) HotMap(hotmap map[string]string) *_oneHotEncodingPreprocessor {

	s.v.HotMap = hotmap
	return s
}

func (s *_oneHotEncodingPreprocessor) AddHotMap(key string, value string) *_oneHotEncodingPreprocessor {

	var tmp map[string]string
	if s.v.HotMap == nil {
		s.v.HotMap = make(map[string]string)
	} else {
		tmp = s.v.HotMap
	}

	tmp[key] = value

	s.v.HotMap = tmp
	return s
}

func (s *_oneHotEncodingPreprocessor) PreprocessorCaster() *types.Preprocessor {
	container := types.NewPreprocessor()

	container.OneHotEncoding = s.v

	return container
}

func (s *_oneHotEncodingPreprocessor) OneHotEncodingPreprocessorCaster() *types.OneHotEncodingPreprocessor {
	return s.v
}
