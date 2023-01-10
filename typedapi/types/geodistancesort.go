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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// GeoDistanceSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/sort.ts#L58-L66
type GeoDistanceSort struct {
	DistanceType    *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`
	GeoDistanceSort map[string][]GeoLocation         `json:"-"`
	IgnoreUnmapped  *bool                            `json:"ignore_unmapped,omitempty"`
	Mode            *sortmode.SortMode               `json:"mode,omitempty"`
	Order           *sortorder.SortOrder             `json:"order,omitempty"`
	Unit            *distanceunit.DistanceUnit       `json:"unit,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoDistanceSort) MarshalJSON() ([]byte, error) {
	type opt GeoDistanceSort
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
	for key, value := range s.GeoDistanceSort {
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewGeoDistanceSort returns a GeoDistanceSort.
func NewGeoDistanceSort() *GeoDistanceSort {
	r := &GeoDistanceSort{
		GeoDistanceSort: make(map[string][]GeoLocation, 0),
	}

	return r
}
