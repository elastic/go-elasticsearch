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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

// GeoDistanceSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/sort.ts#L58-L70
type GeoDistanceSort struct {
	DistanceType    *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`
	GeoDistanceSort map[string][]GeoLocation         `json:"-"`
	IgnoreUnmapped  *bool                            `json:"ignore_unmapped,omitempty"`
	Mode            *sortmode.SortMode               `json:"mode,omitempty"`
	Nested          *NestedSortValue                 `json:"nested,omitempty"`
	Order           *sortorder.SortOrder             `json:"order,omitempty"`
	Unit            *distanceunit.DistanceUnit       `json:"unit,omitempty"`
}

func (s *GeoDistanceSort) UnmarshalJSON(data []byte) error {

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

		case "distance_type":
			if err := dec.Decode(&s.DistanceType); err != nil {
				return fmt.Errorf("%s | %w", "DistanceType", err)
			}

		case "ignore_unmapped":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreUnmapped", err)
				}
				s.IgnoreUnmapped = &value
			case bool:
				s.IgnoreUnmapped = &v
			}

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}

		case "nested":
			if err := dec.Decode(&s.Nested); err != nil {
				return fmt.Errorf("%s | %w", "Nested", err)
			}

		case "order":
			if err := dec.Decode(&s.Order); err != nil {
				return fmt.Errorf("%s | %w", "Order", err)
			}

		case "unit":
			if err := dec.Decode(&s.Unit); err != nil {
				return fmt.Errorf("%s | %w", "Unit", err)
			}

		default:

			rawMsg := make(map[string]json.RawMessage, 0)
			dec.Decode(&rawMsg)
			for key, value := range rawMsg {
				switch {
				case bytes.HasPrefix(value, []byte("\"")), bytes.HasPrefix(value, []byte("{")):
					o := new(GeoLocation)
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return fmt.Errorf("%s | %w", "GeoDistanceSort", err)
					}
					s.GeoDistanceSort[key] = append(s.GeoDistanceSort[key], o)
				default:
					o := []GeoLocation{}
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return fmt.Errorf("%s | %w", "GeoDistanceSort", err)
					}
					s.GeoDistanceSort[key] = o
				}
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoDistanceSort) MarshalJSON() ([]byte, error) {
	type opt GeoDistanceSort
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
	for key, value := range s.GeoDistanceSort {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "GeoDistanceSort")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewGeoDistanceSort returns a GeoDistanceSort.
func NewGeoDistanceSort() *GeoDistanceSort {
	r := &GeoDistanceSort{
		GeoDistanceSort: make(map[string][]GeoLocation),
	}

	return r
}

type GeoDistanceSortVariant interface {
	GeoDistanceSortCaster() *GeoDistanceSort
}

func (s *GeoDistanceSort) GeoDistanceSortCaster() *GeoDistanceSort {
	return s
}
