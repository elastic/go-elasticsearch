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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// TermsExclude type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/bucket.ts#L421-L422
type TermsExclude []string

// TermsExcludeBuilder holds TermsExclude struct and provides a builder API.
type TermsExcludeBuilder struct {
	v TermsExclude
}

// NewTermsExclude provides a builder for the TermsExclude struct.
func NewTermsExcludeBuilder() *TermsExcludeBuilder {
	return &TermsExcludeBuilder{}
}

// Build finalize the chain and returns the TermsExclude struct
func (b *TermsExcludeBuilder) Build() TermsExclude {
	return b.v
}

func (b *TermsExcludeBuilder) TermsExclude(value TermsExclude) *TermsExcludeBuilder {
	b.v = value
	return b
}
