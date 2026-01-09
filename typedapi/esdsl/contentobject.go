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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _contentObject struct {
	v *types.ContentObject
}

func NewContentObject(text string, type_ string) *_contentObject {

	tmp := &_contentObject{v: types.NewContentObject()}

	tmp.Text(text)

	tmp.Type(type_)

	return tmp

}

func (s *_contentObject) Text(text string) *_contentObject {

	s.v.Text = text

	return s
}

func (s *_contentObject) Type(type_ string) *_contentObject {

	s.v.Type = type_

	return s
}

func (s *_contentObject) ContentObjectCaster() *types.ContentObject {
	return s.v
}
