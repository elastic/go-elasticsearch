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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// FieldTypesMappings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L109-L134
type FieldTypesMappings struct {
	// FieldTypes Contains statistics about field data types used in selected nodes.
	FieldTypes []FieldTypes `json:"field_types"`
	// RuntimeFieldTypes Contains statistics about runtime field data types used in selected nodes.
	RuntimeFieldTypes []ClusterRuntimeFieldTypes `json:"runtime_field_types,omitempty"`
	// TotalDeduplicatedFieldCount Total number of fields in all non-system indices, accounting for mapping
	// deduplication.
	TotalDeduplicatedFieldCount *int `json:"total_deduplicated_field_count,omitempty"`
	// TotalDeduplicatedMappingSize Total size of all mappings after deduplication and compression.
	TotalDeduplicatedMappingSize ByteSize `json:"total_deduplicated_mapping_size,omitempty"`
	// TotalDeduplicatedMappingSizeInBytes Total size of all mappings, in bytes, after deduplication and compression.
	TotalDeduplicatedMappingSizeInBytes *int64 `json:"total_deduplicated_mapping_size_in_bytes,omitempty"`
	// TotalFieldCount Total number of fields in all non-system indices.
	TotalFieldCount *int `json:"total_field_count,omitempty"`
}

func (s *FieldTypesMappings) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "field_types":
			if err := dec.Decode(&s.FieldTypes); err != nil {
				return err
			}

		case "runtime_field_types":
			if err := dec.Decode(&s.RuntimeFieldTypes); err != nil {
				return err
			}

		case "total_deduplicated_field_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TotalDeduplicatedFieldCount = &value
			case float64:
				f := int(v)
				s.TotalDeduplicatedFieldCount = &f
			}

		case "total_deduplicated_mapping_size":
			if err := dec.Decode(&s.TotalDeduplicatedMappingSize); err != nil {
				return err
			}

		case "total_deduplicated_mapping_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalDeduplicatedMappingSizeInBytes = &value
			case float64:
				f := int64(v)
				s.TotalDeduplicatedMappingSizeInBytes = &f
			}

		case "total_field_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TotalFieldCount = &value
			case float64:
				f := int(v)
				s.TotalFieldCount = &f
			}

		}
	}
	return nil
}

// NewFieldTypesMappings returns a FieldTypesMappings.
func NewFieldTypesMappings() *FieldTypesMappings {
	r := &FieldTypesMappings{}

	return r
}
