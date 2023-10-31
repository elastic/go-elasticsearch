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

// GeoHashGridAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/bucket.ts#L405-L430
type GeoHashGridAggregation struct {
	// Bounds The bounding box to filter the points in each bucket.
	Bounds GeoBounds `json:"bounds,omitempty"`
	// Field Field containing indexed `geo_point` or `geo_shape` values.
	// If the field contains an array, `geohash_grid` aggregates all array values.
	Field *string  `json:"field,omitempty"`
	Meta  Metadata `json:"meta,omitempty"`
	Name  *string  `json:"name,omitempty"`
	// Precision The string length of the geohashes used to define cells/buckets in the
	// results.
	Precision GeoHashPrecision `json:"precision,omitempty"`
	// ShardSize Allows for more accurate counting of the top cells returned in the final
	// result the aggregation.
	// Defaults to returning `max(10,(size x number-of-shards))` buckets from each
	// shard.
	ShardSize *int `json:"shard_size,omitempty"`
	// Size The maximum number of geohash buckets to return.
	Size *int `json:"size,omitempty"`
}

func (s *GeoHashGridAggregation) UnmarshalJSON(data []byte) error {

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

		case "bounds":
			if err := dec.Decode(&s.Bounds); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "precision":
			if err := dec.Decode(&s.Precision); err != nil {
				return err
			}

		case "shard_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
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

// NewGeoHashGridAggregation returns a GeoHashGridAggregation.
func NewGeoHashGridAggregation() *GeoHashGridAggregation {
	r := &GeoHashGridAggregation{}

	return r
}
