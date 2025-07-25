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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _tableValuesContainer struct {
	v *types.TableValuesContainer
}

func NewTableValuesContainer() *_tableValuesContainer {
	return &_tableValuesContainer{v: types.NewTableValuesContainer()}
}

// AdditionalTableValuesContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_tableValuesContainer) AdditionalTableValuesContainerProperty(key string, value json.RawMessage) *_tableValuesContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalTableValuesContainerProperty = tmp
	return s
}

func (s *_tableValuesContainer) Float64(float64s ...[]types.Float64) *_tableValuesContainer {

	for _, v := range float64s {

		s.v.Float64 = append(s.v.Float64, v)

	}
	return s
}

func (s *_tableValuesContainer) Int(ints ...[]int) *_tableValuesContainer {

	for _, v := range ints {

		s.v.Int = append(s.v.Int, v)

	}
	return s
}

func (s *_tableValuesContainer) Int64(int64s ...[]int64) *_tableValuesContainer {

	for _, v := range int64s {

		s.v.Int64 = append(s.v.Int64, v)

	}
	return s
}

func (s *_tableValuesContainer) Keyword(keywords ...[]string) *_tableValuesContainer {

	for _, v := range keywords {

		s.v.Keyword = append(s.v.Keyword, v)

	}
	return s
}

func (s *_tableValuesContainer) TableValuesContainerCaster() *types.TableValuesContainer {
	return s.v
}
