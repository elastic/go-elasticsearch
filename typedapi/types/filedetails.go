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

// FileDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/recovery/types.ts#L50-L54
type FileDetails struct {
	Length    int64  `json:"length"`
	Name      string `json:"name"`
	Recovered int64  `json:"recovered"`
}

// FileDetailsBuilder holds FileDetails struct and provides a builder API.
type FileDetailsBuilder struct {
	v *FileDetails
}

// NewFileDetails provides a builder for the FileDetails struct.
func NewFileDetailsBuilder() *FileDetailsBuilder {
	r := FileDetailsBuilder{
		&FileDetails{},
	}

	return &r
}

// Build finalize the chain and returns the FileDetails struct
func (rb *FileDetailsBuilder) Build() FileDetails {
	return *rb.v
}

func (rb *FileDetailsBuilder) Length(length int64) *FileDetailsBuilder {
	rb.v.Length = length
	return rb
}

func (rb *FileDetailsBuilder) Name(name string) *FileDetailsBuilder {
	rb.v.Name = name
	return rb
}

func (rb *FileDetailsBuilder) Recovered(recovered int64) *FileDetailsBuilder {
	rb.v.Recovered = recovered
	return rb
}
