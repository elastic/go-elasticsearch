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

// RecoveryFiles type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/recovery/types.ts#L56-L62
type RecoveryFiles struct {
	Details   []FileDetails `json:"details,omitempty"`
	Percent   Percentage    `json:"percent"`
	Recovered int64         `json:"recovered"`
	Reused    int64         `json:"reused"`
	Total     int64         `json:"total"`
}

// RecoveryFilesBuilder holds RecoveryFiles struct and provides a builder API.
type RecoveryFilesBuilder struct {
	v *RecoveryFiles
}

// NewRecoveryFiles provides a builder for the RecoveryFiles struct.
func NewRecoveryFilesBuilder() *RecoveryFilesBuilder {
	r := RecoveryFilesBuilder{
		&RecoveryFiles{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryFiles struct
func (rb *RecoveryFilesBuilder) Build() RecoveryFiles {
	return *rb.v
}

func (rb *RecoveryFilesBuilder) Details(details []FileDetailsBuilder) *RecoveryFilesBuilder {
	tmp := make([]FileDetails, len(details))
	for _, value := range details {
		tmp = append(tmp, value.Build())
	}
	rb.v.Details = tmp
	return rb
}

func (rb *RecoveryFilesBuilder) Percent(percent *PercentageBuilder) *RecoveryFilesBuilder {
	v := percent.Build()
	rb.v.Percent = v
	return rb
}

func (rb *RecoveryFilesBuilder) Recovered(recovered int64) *RecoveryFilesBuilder {
	rb.v.Recovered = recovered
	return rb
}

func (rb *RecoveryFilesBuilder) Reused(reused int64) *RecoveryFilesBuilder {
	rb.v.Reused = reused
	return rb
}

func (rb *RecoveryFilesBuilder) Total(total int64) *RecoveryFilesBuilder {
	rb.v.Total = total
	return rb
}
