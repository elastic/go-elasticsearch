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

// CompletionContext type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/search/_types/suggester.ts#L232-L261
type CompletionContext struct {
	// Boost The factor by which the score of the suggestion should be boosted.
	// The score is computed by multiplying the boost with the suggestion weight.
	Boost *Float64 `json:"boost,omitempty"`
	// Context The value of the category to filter/boost on.
	Context Context `json:"context"`
	// Neighbours An array of precision values at which neighboring geohashes should be taken
	// into account.
	// Precision value can be a distance value (`5m`, `10km`, etc.) or a raw geohash
	// precision (`1`..`12`).
	// Defaults to generating neighbors for index time precision level.
	Neighbours []GeoHashPrecision `json:"neighbours,omitempty"`
	// Precision The precision of the geohash to encode the query geo point.
	// Can be specified as a distance value (`5m`, `10km`, etc.), or as a raw
	// geohash precision (`1`..`12`).
	// Defaults to index time precision level.
	Precision GeoHashPrecision `json:"precision,omitempty"`
	// Prefix Whether the category value should be treated as a prefix or not.
	Prefix *bool `json:"prefix,omitempty"`
}

func (s *CompletionContext) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Context)
		return err
	}

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
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Boost = &f
			case float64:
				f := Float64(v)
				s.Boost = &f
			}

		case "context":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := new(GeoLocation)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Context = *o

			default:
				if err := localDec.Decode(&s.Context); err != nil {
					return err
				}
			}

		case "neighbours":
			if err := dec.Decode(&s.Neighbours); err != nil {
				return err
			}

		case "precision":
			if err := dec.Decode(&s.Precision); err != nil {
				return err
			}

		case "prefix":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Prefix = &value
			case bool:
				s.Prefix = &v
			}

		}
	}
	return nil
}

// NewCompletionContext returns a CompletionContext.
func NewCompletionContext() *CompletionContext {
	r := &CompletionContext{}

	return r
}
