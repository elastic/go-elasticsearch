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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _reasoningDetail struct {
	v types.ReasoningDetail
}

func NewReasoningDetail() *_reasoningDetail {
	return &_reasoningDetail{v: nil}
}

func (u *_reasoningDetail) EncryptedReasoningDetail(encryptedreasoningdetail types.EncryptedReasoningDetailVariant) *_reasoningDetail {

	u.v = encryptedreasoningdetail.EncryptedReasoningDetailCaster()

	return u
}

// Interface implementation for EncryptedReasoningDetail in ReasoningDetail union
func (u *_encryptedReasoningDetail) ReasoningDetailCaster() *types.ReasoningDetail {
	t := types.ReasoningDetail(u.v)
	return &t
}

func (u *_reasoningDetail) SummaryReasoningDetail(summaryreasoningdetail types.SummaryReasoningDetailVariant) *_reasoningDetail {

	u.v = summaryreasoningdetail.SummaryReasoningDetailCaster()

	return u
}

// Interface implementation for SummaryReasoningDetail in ReasoningDetail union
func (u *_summaryReasoningDetail) ReasoningDetailCaster() *types.ReasoningDetail {
	t := types.ReasoningDetail(u.v)
	return &t
}

func (u *_reasoningDetail) TextReasoningDetail(textreasoningdetail types.TextReasoningDetailVariant) *_reasoningDetail {

	u.v = textreasoningdetail.TextReasoningDetailCaster()

	return u
}

// Interface implementation for TextReasoningDetail in ReasoningDetail union
func (u *_textReasoningDetail) ReasoningDetailCaster() *types.ReasoningDetail {
	t := types.ReasoningDetail(u.v)
	return &t
}

func (u *_reasoningDetail) ReasoningDetailCaster() *types.ReasoningDetail {
	return &u.v
}
