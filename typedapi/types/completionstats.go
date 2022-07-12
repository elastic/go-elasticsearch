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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// CompletionStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/Stats.ts#L53-L57
type CompletionStats struct {
	Fields      map[Field]FieldSizeUsage `json:"fields,omitempty"`
	Size        *ByteSize                `json:"size,omitempty"`
	SizeInBytes int64                    `json:"size_in_bytes"`
}

// CompletionStatsBuilder holds CompletionStats struct and provides a builder API.
type CompletionStatsBuilder struct {
	v *CompletionStats
}

// NewCompletionStats provides a builder for the CompletionStats struct.
func NewCompletionStatsBuilder() *CompletionStatsBuilder {
	r := CompletionStatsBuilder{
		&CompletionStats{
			Fields: make(map[Field]FieldSizeUsage, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CompletionStats struct
func (rb *CompletionStatsBuilder) Build() CompletionStats {
	return *rb.v
}

func (rb *CompletionStatsBuilder) Fields(values map[Field]*FieldSizeUsageBuilder) *CompletionStatsBuilder {
	tmp := make(map[Field]FieldSizeUsage, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *CompletionStatsBuilder) Size(size *ByteSizeBuilder) *CompletionStatsBuilder {
	v := size.Build()
	rb.v.Size = &v
	return rb
}

func (rb *CompletionStatsBuilder) SizeInBytes(sizeinbytes int64) *CompletionStatsBuilder {
	rb.v.SizeInBytes = sizeinbytes
	return rb
}
