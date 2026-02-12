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
type _msearchRequestItem struct {
	v types.MsearchRequestItem
}

func NewMsearchRequestItem() *_msearchRequestItem {
	return &_msearchRequestItem{v: nil}
}

func (u *_msearchRequestItem) MultisearchHeader(multisearchheader types.MultisearchHeaderVariant) *_msearchRequestItem {

	u.v = multisearchheader.MultisearchHeaderCaster()

	return u
}

// Interface implementation for MultisearchHeader in MsearchRequestItem union
func (u *_multisearchHeader) MsearchRequestItemCaster() *types.MsearchRequestItem {
	t := types.MsearchRequestItem(u.v)
	return &t
}

func (u *_msearchRequestItem) SearchRequestBody(searchrequestbody types.SearchRequestBodyVariant) *_msearchRequestItem {

	u.v = searchrequestbody.SearchRequestBodyCaster()

	return u
}

// Interface implementation for SearchRequestBody in MsearchRequestItem union
func (u *_searchRequestBody) MsearchRequestItemCaster() *types.MsearchRequestItem {
	t := types.MsearchRequestItem(u.v)
	return &t
}

func (u *_msearchRequestItem) MsearchRequestItemCaster() *types.MsearchRequestItem {
	return &u.v
}
