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

// BuildInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/info/types.ts#L24-L27
type BuildInformation struct {
	Date DateTime `json:"date"`
	Hash string   `json:"hash"`
}

// BuildInformationBuilder holds BuildInformation struct and provides a builder API.
type BuildInformationBuilder struct {
	v *BuildInformation
}

// NewBuildInformation provides a builder for the BuildInformation struct.
func NewBuildInformationBuilder() *BuildInformationBuilder {
	r := BuildInformationBuilder{
		&BuildInformation{},
	}

	return &r
}

// Build finalize the chain and returns the BuildInformation struct
func (rb *BuildInformationBuilder) Build() BuildInformation {
	return *rb.v
}

func (rb *BuildInformationBuilder) Date(date *DateTimeBuilder) *BuildInformationBuilder {
	v := date.Build()
	rb.v.Date = v
	return rb
}

func (rb *BuildInformationBuilder) Hash(hash string) *BuildInformationBuilder {
	rb.v.Hash = hash
	return rb
}
