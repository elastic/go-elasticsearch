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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/imageurldetail"
)

type _imageUrl struct {
	v *types.ImageUrl
}

func NewImageUrl(url string) *_imageUrl {

	tmp := &_imageUrl{v: types.NewImageUrl()}

	tmp.Url(url)

	return tmp

}

func (s *_imageUrl) Detail(detail imageurldetail.ImageUrlDetail) *_imageUrl {

	s.v.Detail = &detail
	return s
}

func (s *_imageUrl) Url(url string) *_imageUrl {

	s.v.Url = url

	return s
}

func (s *_imageUrl) ImageUrlCaster() *types.ImageUrl {
	return s.v
}
