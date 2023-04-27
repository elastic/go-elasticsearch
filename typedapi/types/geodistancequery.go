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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"

	"fmt"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// GeoDistanceQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/query_dsl/geo.ts#L48-L57
type GeoDistanceQuery struct {
	Boost            *float32                                 `json:"boost,omitempty"`
	Distance         *string                                  `json:"distance,omitempty"`
	DistanceType     *geodistancetype.GeoDistanceType         `json:"distance_type,omitempty"`
	GeoDistanceQuery map[string]GeoLocation                   `json:"GeoDistanceQuery,omitempty"`
	QueryName_       *string                                  `json:"_name,omitempty"`
	ValidationMethod *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

func (s *GeoDistanceQuery) UnmarshalJSON(data []byte) error {

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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "distance":
			if err := dec.Decode(&s.Distance); err != nil {
				return err
			}

		case "distance_type":
			if err := dec.Decode(&s.DistanceType); err != nil {
				return err
			}

		case "GeoDistanceQuery":
			if s.GeoDistanceQuery == nil {
				s.GeoDistanceQuery = make(map[string]GeoLocation, 0)
			}
			if err := dec.Decode(&s.GeoDistanceQuery); err != nil {
				return err
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.QueryName_ = &o

		case "validation_method":
			if err := dec.Decode(&s.ValidationMethod); err != nil {
				return err
			}

		default:

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoDistanceQuery) MarshalJSON() ([]byte, error) {
	type opt GeoDistanceQuery
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
	for key, value := range s.GeoDistanceQuery {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "GeoDistanceQuery")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewGeoDistanceQuery returns a GeoDistanceQuery.
func NewGeoDistanceQuery() *GeoDistanceQuery {
	r := &GeoDistanceQuery{
		GeoDistanceQuery: make(map[string]GeoLocation, 0),
	}

	return r
}
