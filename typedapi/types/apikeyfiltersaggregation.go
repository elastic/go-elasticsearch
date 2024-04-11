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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ApiKeyFiltersAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/security/query_api_keys/types.ts#L208-L228
type ApiKeyFiltersAggregation struct {
	// Filters Collection of queries from which to build buckets.
	Filters BucketsApiKeyQueryContainer `json:"filters,omitempty"`
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

func (s *ApiKeyFiltersAggregation) UnmarshalJSON(data []byte) error {

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
				o := make(map[string]ApiKeyQueryContainer, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Filters", err)
				}
				s.Filters = o
			case '[':
				o := []ApiKeyQueryContainer{}
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Filters", err)
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
					return fmt.Errorf("%s | %w", "Keyed", err)
				}
				s.Keyed = &value
			case bool:
				s.Keyed = &v
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
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
					return fmt.Errorf("%s | %w", "OtherBucket", err)
				}
				s.OtherBucket = &value
			case bool:
				s.OtherBucket = &v
			}

		case "other_bucket_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "OtherBucketKey", err)
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

// NewApiKeyFiltersAggregation returns a ApiKeyFiltersAggregation.
func NewApiKeyFiltersAggregation() *ApiKeyFiltersAggregation {
	r := &ApiKeyFiltersAggregation{}

	return r
}
