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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SparseVectorQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/SparseVectorQuery.ts#L26-L80
type SparseVectorQuery struct {
	AdditionalSparseVectorQueryProperty map[string]json.RawMessage `json:"-"`
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// Field The name of the field that contains the token-weight pairs to be searched
	// against.
	// This field must be a mapped sparse_vector field.
	Field string `json:"field"`
	// InferenceId The inference ID to use to convert the query text into token-weight pairs.
	// It must be the same inference ID that was used to create the tokens from the
	// input text.
	// Only one of inference_id and query_vector is allowed.
	// If inference_id is specified, query must also be specified.
	// Only one of inference_id or query_vector may be supplied in a request.
	InferenceId *string `json:"inference_id,omitempty"`
	// Prune Whether to perform pruning, omitting the non-significant tokens from the
	// query to improve query performance.
	// If prune is true but the pruning_config is not specified, pruning will occur
	// but default values will be used.
	// Default: false
	Prune *bool `json:"prune,omitempty"`
	// PruningConfig Optional pruning configuration.
	// If enabled, this will omit non-significant tokens from the query in order to
	// improve query performance.
	// This is only used if prune is set to true.
	// If prune is set to true but pruning_config is not specified, default values
	// will be used.
	PruningConfig *TokenPruningConfig `json:"pruning_config,omitempty"`
	// Query The query text you want to use for search.
	// If inference_id is specified, query must also be specified.
	Query      *string `json:"query,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`
	// QueryVector Dictionary of precomputed sparse vectors and their associated weights.
	// Only one of inference_id or query_vector may be supplied in a request.
	QueryVector map[string]float32 `json:"query_vector,omitempty"`
}

func (s *SparseVectorQuery) UnmarshalJSON(data []byte) error {

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

		case "boost":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Boost", err)
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "inference_id":
			if err := dec.Decode(&s.InferenceId); err != nil {
				return fmt.Errorf("%s | %w", "InferenceId", err)
			}

		case "prune":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Prune", err)
				}
				s.Prune = &value
			case bool:
				s.Prune = &v
			}

		case "pruning_config":
			if err := dec.Decode(&s.PruningConfig); err != nil {
				return fmt.Errorf("%s | %w", "PruningConfig", err)
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = &o

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueryName_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "query_vector":
			if s.QueryVector == nil {
				s.QueryVector = make(map[string]float32, 0)
			}
			if err := dec.Decode(&s.QueryVector); err != nil {
				return fmt.Errorf("%s | %w", "QueryVector", err)
			}

		default:

			if key, ok := t.(string); ok {
				if s.AdditionalSparseVectorQueryProperty == nil {
					s.AdditionalSparseVectorQueryProperty = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "AdditionalSparseVectorQueryProperty", err)
				}
				s.AdditionalSparseVectorQueryProperty[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s SparseVectorQuery) MarshalJSON() ([]byte, error) {
	type opt SparseVectorQuery
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalSparseVectorQueryProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalSparseVectorQueryProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewSparseVectorQuery returns a SparseVectorQuery.
func NewSparseVectorQuery() *SparseVectorQuery {
	r := &SparseVectorQuery{
		AdditionalSparseVectorQueryProperty: make(map[string]json.RawMessage),
		QueryVector:                         make(map[string]float32),
	}

	return r
}
