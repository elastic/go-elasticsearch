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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _requestItem struct {
	v types.RequestItem
}

func NewRequestItem() *_requestItem {
	return &_requestItem{v: nil}
}

func (u *_requestItem) MultisearchHeader(multisearchheader types.MultisearchHeaderVariant) *_requestItem {

	u.v = multisearchheader.MultisearchHeaderCaster()

	return u
}

// Interface implementation for MultisearchHeader in RequestItem union
func (u *_multisearchHeader) RequestItemCaster() *types.RequestItem {
	t := types.RequestItem(u.v)
	return &t
}

func (u *_requestItem) TemplateConfig(templateconfig types.TemplateConfigVariant) *_requestItem {

	u.v = templateconfig.TemplateConfigCaster()

	return u
}

// Interface implementation for TemplateConfig in RequestItem union
func (u *_templateConfig) RequestItemCaster() *types.RequestItem {
	t := types.RequestItem(u.v)
	return &t
}

func (u *_requestItem) RequestItemCaster() *types.RequestItem {
	return &u.v
}
