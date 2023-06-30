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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Vector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/xpack/usage/types.ts#L445-L449
type Vector struct {
	Available               bool `json:"available"`
	DenseVectorDimsAvgCount int  `json:"dense_vector_dims_avg_count"`
	DenseVectorFieldsCount  int  `json:"dense_vector_fields_count"`
	Enabled                 bool `json:"enabled"`
	SparseVectorFieldsCount *int `json:"sparse_vector_fields_count,omitempty"`
}

func (s *Vector) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "dense_vector_dims_avg_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DenseVectorDimsAvgCount = value
			case float64:
				f := int(v)
				s.DenseVectorDimsAvgCount = f
			}

		case "dense_vector_fields_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DenseVectorFieldsCount = value
			case float64:
				f := int(v)
				s.DenseVectorFieldsCount = f
			}

		case "enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "sparse_vector_fields_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SparseVectorFieldsCount = &value
			case float64:
				f := int(v)
				s.SparseVectorFieldsCount = &f
			}

		}
	}
	return nil
}

// NewVector returns a Vector.
func NewVector() *Vector {
	r := &Vector{}

	return r
}
