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

import (
	"encoding/json"
)

// SortOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/sort.ts#L80-L89
type SortOptions struct {
	Doc_         *ScoreSort          `json:"_doc,omitempty"`
	GeoDistance_ *GeoDistanceSort    `json:"_geo_distance,omitempty"`
	Score_       *ScoreSort          `json:"_score,omitempty"`
	Script_      *ScriptSort         `json:"_script,omitempty"`
	SortOptions  map[Field]FieldSort `json:"-"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s SortOptions) MarshalJSON() ([]byte, error) {
	type opt SortOptions
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.SortOptions {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SortOptionsBuilder holds SortOptions struct and provides a builder API.
type SortOptionsBuilder struct {
	v *SortOptions
}

// NewSortOptions provides a builder for the SortOptions struct.
func NewSortOptionsBuilder() *SortOptionsBuilder {
	r := SortOptionsBuilder{
		&SortOptions{
			SortOptions: make(map[Field]FieldSort, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SortOptions struct
func (rb *SortOptionsBuilder) Build() SortOptions {
	return *rb.v
}

func (rb *SortOptionsBuilder) Doc_(doc_ *ScoreSortBuilder) *SortOptionsBuilder {
	v := doc_.Build()
	rb.v.Doc_ = &v
	return rb
}

func (rb *SortOptionsBuilder) GeoDistance_(geodistance_ *GeoDistanceSortBuilder) *SortOptionsBuilder {
	v := geodistance_.Build()
	rb.v.GeoDistance_ = &v
	return rb
}

func (rb *SortOptionsBuilder) Score_(score_ *ScoreSortBuilder) *SortOptionsBuilder {
	v := score_.Build()
	rb.v.Score_ = &v
	return rb
}

func (rb *SortOptionsBuilder) Script_(script_ *ScriptSortBuilder) *SortOptionsBuilder {
	v := script_.Build()
	rb.v.Script_ = &v
	return rb
}

func (rb *SortOptionsBuilder) SortOptions(values map[Field]*FieldSortBuilder) *SortOptionsBuilder {
	tmp := make(map[Field]FieldSort, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.SortOptions = tmp
	return rb
}
