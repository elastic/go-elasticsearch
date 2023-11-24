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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// FiltersAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/aggregations/bucket.ts#L358-L378
type FiltersAggregation struct {
	// Filters Collection of queries from which to build buckets.
	Filters BucketsQuery `json:"filters,omitempty"`
	// Keyed By default, the named filters aggregation returns the buckets as an object.
	// Set to `false` to return the buckets as an array of objects.
	Keyed *bool    `json:"keyed,omitempty"`
	Meta  Metadata `json:"meta,omitempty"`
	Name  *string  `json:"name,omitempty"`
	// OtherBucket Set to `true` to add a bucket to the response which will contain all
	// documents that do not match any of the given filters.
	OtherBucket *bool `json:"other_bucket,omitempty"`
	// OtherBucketKey The key with which the other bucket is returned.
	OtherBucketKey *string `json:"other_bucket_key,omitempty"`
}

func (s *FiltersAggregation) UnmarshalJSON(data []byte) error {

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

		case "filters":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]Query, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Filters = o
			case '[':
				o := []Query{}
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Filters = o
			}

		case "keyed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Keyed = &value
			case bool:
				s.Keyed = &v
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

		case "other_bucket":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.OtherBucket = &value
			case bool:
				s.OtherBucket = &v
			}

		case "other_bucket_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.OtherBucketKey = &o

		}
	}
	return nil
}

// NewFiltersAggregation returns a FiltersAggregation.
func NewFiltersAggregation() *FiltersAggregation {
	r := &FiltersAggregation{}

	return r
}
