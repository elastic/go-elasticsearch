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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rerankinputformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rerankinputtype"
)

type _rerankInputObject struct {
	v *types.RerankInputObject
}

func NewRerankInputObject(type_ rerankinputtype.RerankInputType, value string) *_rerankInputObject {

	tmp := &_rerankInputObject{v: types.NewRerankInputObject()}

	tmp.Type(type_)

	tmp.Value(value)

	return tmp

}

func (s *_rerankInputObject) Format(format rerankinputformat.RerankInputFormat) *_rerankInputObject {

	s.v.Format = &format
	return s
}

func (s *_rerankInputObject) Type(type_ rerankinputtype.RerankInputType) *_rerankInputObject {

	s.v.Type = type_
	return s
}

func (s *_rerankInputObject) Value(value string) *_rerankInputObject {

	s.v.Value = value

	return s
}

func (s *_rerankInputObject) RerankInputObjectCaster() *types.RerankInputObject {
	return s.v
}
