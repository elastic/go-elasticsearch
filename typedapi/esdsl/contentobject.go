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
// https://github.com/elastic/elasticsearch-specification/tree/1cdb0310e2ed76dea2491f2551fca3f9143b6dc4

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _contentObject struct {
	v types.ContentObject
}

func NewContentObject() *_contentObject {
	return &_contentObject{v: nil}
}

func (u *_contentObject) TextContent(textcontent types.TextContentVariant) *_contentObject {

	u.v = textcontent.TextContentCaster()

	return u
}

// Interface implementation for TextContent in ContentObject union
func (u *_textContent) ContentObjectCaster() *types.ContentObject {
	t := types.ContentObject(u.v)
	return &t
}

func (u *_contentObject) ImageUrlContent(imageurlcontent types.ImageUrlContentVariant) *_contentObject {

	u.v = imageurlcontent.ImageUrlContentCaster()

	return u
}

// Interface implementation for ImageUrlContent in ContentObject union
func (u *_imageUrlContent) ContentObjectCaster() *types.ContentObject {
	t := types.ContentObject(u.v)
	return &t
}

func (u *_contentObject) FileContentPart(filecontentpart types.FileContentPartVariant) *_contentObject {

	u.v = filecontentpart.FileContentPartCaster()

	return u
}

// Interface implementation for FileContentPart in ContentObject union
func (u *_fileContentPart) ContentObjectCaster() *types.ContentObject {
	t := types.ContentObject(u.v)
	return &t
}

func (u *_contentObject) ContentObjectCaster() *types.ContentObject {
	return &u.v
}
