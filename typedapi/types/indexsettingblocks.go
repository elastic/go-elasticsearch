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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// IndexSettingBlocks type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L245-L251
type IndexSettingBlocks struct {
	Metadata            *bool  `json:"metadata,omitempty"`
	Read                *bool  `json:"read,omitempty"`
	ReadOnly            *bool  `json:"read_only,omitempty"`
	ReadOnlyAllowDelete *bool  `json:"read_only_allow_delete,omitempty"`
	Write               string `json:"write,omitempty"`
}

// IndexSettingBlocksBuilder holds IndexSettingBlocks struct and provides a builder API.
type IndexSettingBlocksBuilder struct {
	v *IndexSettingBlocks
}

// NewIndexSettingBlocks provides a builder for the IndexSettingBlocks struct.
func NewIndexSettingBlocksBuilder() *IndexSettingBlocksBuilder {
	r := IndexSettingBlocksBuilder{
		&IndexSettingBlocks{},
	}

	return &r
}

// Build finalize the chain and returns the IndexSettingBlocks struct
func (rb *IndexSettingBlocksBuilder) Build() IndexSettingBlocks {
	return *rb.v
}

func (rb *IndexSettingBlocksBuilder) Metadata(metadata bool) *IndexSettingBlocksBuilder {
	rb.v.Metadata = &metadata
	return rb
}

func (rb *IndexSettingBlocksBuilder) Read(read bool) *IndexSettingBlocksBuilder {
	rb.v.Read = &read
	return rb
}

func (rb *IndexSettingBlocksBuilder) ReadOnly(readonly bool) *IndexSettingBlocksBuilder {
	rb.v.ReadOnly = &readonly
	return rb
}

func (rb *IndexSettingBlocksBuilder) ReadOnlyAllowDelete(readonlyallowdelete bool) *IndexSettingBlocksBuilder {
	rb.v.ReadOnlyAllowDelete = &readonlyallowdelete
	return rb
}

func (rb *IndexSettingBlocksBuilder) Write(arg string) *IndexSettingBlocksBuilder {
	rb.v.Write = arg
	return rb
}
