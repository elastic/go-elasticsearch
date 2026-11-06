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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/documentextractioncontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/documentextractioncontenttype"
)

type _documentExtractionContentObjectItem struct {
	v *types.DocumentExtractionContentObjectItem
}

func NewDocumentExtractionContentObjectItem(type_ documentextractioncontenttype.DocumentExtractionContentType, value string) *_documentExtractionContentObjectItem {

	tmp := &_documentExtractionContentObjectItem{v: types.NewDocumentExtractionContentObjectItem()}

	tmp.Type(type_)

	tmp.Value(value)

	return tmp

}

func (s *_documentExtractionContentObjectItem) Format(format documentextractioncontentformat.DocumentExtractionContentFormat) *_documentExtractionContentObjectItem {

	s.v.Format = &format
	return s
}

func (s *_documentExtractionContentObjectItem) Type(type_ documentextractioncontenttype.DocumentExtractionContentType) *_documentExtractionContentObjectItem {

	s.v.Type = type_
	return s
}

func (s *_documentExtractionContentObjectItem) Value(value string) *_documentExtractionContentObjectItem {

	s.v.Value = value

	return s
}

func (s *_documentExtractionContentObjectItem) DocumentExtractionContentObjectItemCaster() *types.DocumentExtractionContentObjectItem {
	return s.v
}
