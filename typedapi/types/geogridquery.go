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

// GeoGridQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/geo.ts#L98-L103
type GeoGridQuery struct {
	AdditionalGeoGridQueryProperty map[string]json.RawMessage `json:"-"`
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost      *float32 `json:"boost,omitempty"`
	Geohash    *string  `json:"geohash,omitempty"`
	Geohex     *string  `json:"geohex,omitempty"`
	Geotile    *string  `json:"geotile,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
}

func (s *GeoGridQuery) UnmarshalJSON(data []byte) error {

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

		case "geohash":
			if err := dec.Decode(&s.Geohash); err != nil {
				return fmt.Errorf("%s | %w", "Geohash", err)
			}

		case "geohex":
			if err := dec.Decode(&s.Geohex); err != nil {
				return fmt.Errorf("%s | %w", "Geohex", err)
			}

		case "geotile":
			if err := dec.Decode(&s.Geotile); err != nil {
				return fmt.Errorf("%s | %w", "Geotile", err)
			}

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

		default:

			if key, ok := t.(string); ok {
				if s.AdditionalGeoGridQueryProperty == nil {
					s.AdditionalGeoGridQueryProperty = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "AdditionalGeoGridQueryProperty", err)
				}
				s.AdditionalGeoGridQueryProperty[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoGridQuery) MarshalJSON() ([]byte, error) {
	type opt GeoGridQuery
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
	for key, value := range s.AdditionalGeoGridQueryProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalGeoGridQueryProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewGeoGridQuery returns a GeoGridQuery.
func NewGeoGridQuery() *GeoGridQuery {
	r := &GeoGridQuery{
		AdditionalGeoGridQueryProperty: make(map[string]json.RawMessage),
	}

	return r
}
