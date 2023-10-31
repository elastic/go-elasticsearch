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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// FieldTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/stats/types.ts#L136-L167
type FieldTypes struct {
	// Count The number of occurrences of the field type in selected nodes.
	Count int `json:"count"`
	// IndexCount The number of indices containing the field type in selected nodes.
	IndexCount int `json:"index_count"`
	// IndexedVectorCount For dense_vector field types, number of indexed vector types in selected
	// nodes.
	IndexedVectorCount *int64 `json:"indexed_vector_count,omitempty"`
	// IndexedVectorDimMax For dense_vector field types, the maximum dimension of all indexed vector
	// types in selected nodes.
	IndexedVectorDimMax *int64 `json:"indexed_vector_dim_max,omitempty"`
	// IndexedVectorDimMin For dense_vector field types, the minimum dimension of all indexed vector
	// types in selected nodes.
	IndexedVectorDimMin *int64 `json:"indexed_vector_dim_min,omitempty"`
	// Name The name for the field type in selected nodes.
	Name string `json:"name"`
	// ScriptCount The number of fields that declare a script.
	ScriptCount *int `json:"script_count,omitempty"`
}

func (s *FieldTypes) UnmarshalJSON(data []byte) error {

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

		case "count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "index_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.IndexCount = value
			case float64:
				f := int(v)
				s.IndexCount = f
			}

		case "indexed_vector_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndexedVectorCount = &value
			case float64:
				f := int64(v)
				s.IndexedVectorCount = &f
			}

		case "indexed_vector_dim_max":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndexedVectorDimMax = &value
			case float64:
				f := int64(v)
				s.IndexedVectorDimMax = &f
			}

		case "indexed_vector_dim_min":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndexedVectorDimMin = &value
			case float64:
				f := int64(v)
				s.IndexedVectorDimMin = &f
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "script_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ScriptCount = &value
			case float64:
				f := int(v)
				s.ScriptCount = &f
			}

		}
	}
	return nil
}

// NewFieldTypes returns a FieldTypes.
func NewFieldTypes() *FieldTypes {
	r := &FieldTypes{}

	return r
}
