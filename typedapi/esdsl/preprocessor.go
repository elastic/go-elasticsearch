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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _preprocessor struct {
	v *types.Preprocessor
}

func NewPreprocessor() *_preprocessor {
	return &_preprocessor{v: types.NewPreprocessor()}
}

// AdditionalPreprocessorProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_preprocessor) AdditionalPreprocessorProperty(key string, value json.RawMessage) *_preprocessor {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalPreprocessorProperty = tmp
	return s
}

func (s *_preprocessor) FrequencyEncoding(frequencyencoding types.FrequencyEncodingPreprocessorVariant) *_preprocessor {

	s.v.FrequencyEncoding = frequencyencoding.FrequencyEncodingPreprocessorCaster()

	return s
}

func (s *_preprocessor) OneHotEncoding(onehotencoding types.OneHotEncodingPreprocessorVariant) *_preprocessor {

	s.v.OneHotEncoding = onehotencoding.OneHotEncodingPreprocessorCaster()

	return s
}

func (s *_preprocessor) TargetMeanEncoding(targetmeanencoding types.TargetMeanEncodingPreprocessorVariant) *_preprocessor {

	s.v.TargetMeanEncoding = targetmeanencoding.TargetMeanEncodingPreprocessorCaster()

	return s
}

func (s *_preprocessor) PreprocessorCaster() *types.Preprocessor {
	return s.v
}
