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

// ByteSize holds the union for the following types:
//
//	int64
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/common.ts#L81-L82
type ByteSize interface{}

// ByteSizeBuilder holds ByteSize struct and provides a builder API.
type ByteSizeBuilder struct {
	v ByteSize
}

// NewByteSize provides a builder for the ByteSize struct.
func NewByteSizeBuilder() *ByteSizeBuilder {
	return &ByteSizeBuilder{}
}

// Build finalize the chain and returns the ByteSize struct
func (u *ByteSizeBuilder) Build() ByteSize {
	return u.v
}

func (u *ByteSizeBuilder) Int64(int64 int64) *ByteSizeBuilder {
	u.v = &int64
	return u
}

func (u *ByteSizeBuilder) String(string string) *ByteSizeBuilder {
	u.v = &string
	return u
}
