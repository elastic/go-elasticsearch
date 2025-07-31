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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/stopwordlanguage"
)

// This is provide all the types that are part of the union.
type _stopWords struct {
	v types.StopWords
}

func NewStopWords() *_stopWords {
	return &_stopWords{v: nil}
}

func (u *_stopWords) StopWordLanguage(stopwordlanguage stopwordlanguage.StopWordLanguage) *_stopWords {

	u.v = &stopwordlanguage
	return u
}

func (u *_stopWords) Strings(strings ...string) *_stopWords {

	u.v = make([]string, len(strings))
	u.v = strings

	return u
}

func (u *_stopWords) StopWordsCaster() *types.StopWords {
	return &u.v
}
