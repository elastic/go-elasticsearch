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

// VertexDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/graph/_types/Vertex.ts#L30-L59
type VertexDefinition struct {
	// Exclude Prevents the specified terms from being included in the results.
	Exclude []string `json:"exclude,omitempty"`
	// Field Identifies a field in the documents of interest.
	Field string `json:"field"`
	// Include Identifies the terms of interest that form the starting points from which you
	// want to spider out.
	Include []VertexInclude `json:"include,omitempty"`
	// MinDocCount Specifies how many documents must contain a pair of terms before it is
	// considered to be a useful connection.
	// This setting acts as a certainty threshold.
	MinDocCount *int64 `json:"min_doc_count,omitempty"`
	// ShardMinDocCount Controls how many documents on a particular shard have to contain a pair of
	// terms before the connection is returned for global consideration.
	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`
	// Size Specifies the maximum number of vertex terms returned for each field.
	Size *int `json:"size,omitempty"`
}

func (s *VertexDefinition) UnmarshalJSON(data []byte) error {

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

		case "exclude":
			if err := dec.Decode(&s.Exclude); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "include":
			if err := dec.Decode(&s.Include); err != nil {
				return err
			}

		case "min_doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MinDocCount = &value
			case float64:
				f := int64(v)
				s.MinDocCount = &f
			}

		case "shard_min_doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ShardMinDocCount = &value
			case float64:
				f := int64(v)
				s.ShardMinDocCount = &f
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		}
	}
	return nil
}

// NewVertexDefinition returns a VertexDefinition.
func NewVertexDefinition() *VertexDefinition {
	r := &VertexDefinition{}

	return r
}
