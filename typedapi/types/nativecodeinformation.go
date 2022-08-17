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

// NativeCodeInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/info/types.ts#L29-L32
type NativeCodeInformation struct {
	BuildHash string        `json:"build_hash"`
	Version   VersionString `json:"version"`
}

// NativeCodeInformationBuilder holds NativeCodeInformation struct and provides a builder API.
type NativeCodeInformationBuilder struct {
	v *NativeCodeInformation
}

// NewNativeCodeInformation provides a builder for the NativeCodeInformation struct.
func NewNativeCodeInformationBuilder() *NativeCodeInformationBuilder {
	r := NativeCodeInformationBuilder{
		&NativeCodeInformation{},
	}

	return &r
}

// Build finalize the chain and returns the NativeCodeInformation struct
func (rb *NativeCodeInformationBuilder) Build() NativeCodeInformation {
	return *rb.v
}

func (rb *NativeCodeInformationBuilder) BuildHash(buildhash string) *NativeCodeInformationBuilder {
	rb.v.BuildHash = buildhash
	return rb
}

func (rb *NativeCodeInformationBuilder) Version(version VersionString) *NativeCodeInformationBuilder {
	rb.v.Version = version
	return rb
}
