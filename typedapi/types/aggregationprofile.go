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
// https://github.com/elastic/elasticsearch-specification/tree/6e0fb6b929f337b62bf0676bdf503e061121fad2

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// AggregationProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6e0fb6b929f337b62bf0676bdf503e061121fad2/specification/_global/search/_types/profile.ts#L77-L84
type AggregationProfile struct {
	Breakdown   AggregationBreakdown     `json:"breakdown"`
	Children    []AggregationProfile     `json:"children,omitempty"`
	Debug       *AggregationProfileDebug `json:"debug,omitempty"`
	Description string                   `json:"description"`
	TimeInNanos int64                    `json:"time_in_nanos"`
	Type        string                   `json:"type"`
}

func (s *AggregationProfile) UnmarshalJSON(data []byte) error {

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

		case "breakdown":
			if err := dec.Decode(&s.Breakdown); err != nil {
				return err
			}

		case "children":
			if err := dec.Decode(&s.Children); err != nil {
				return err
			}

		case "debug":
			if err := dec.Decode(&s.Debug); err != nil {
				return err
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = o

		case "time_in_nanos":
			if err := dec.Decode(&s.TimeInNanos); err != nil {
				return err
			}

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

		}
	}
	return nil
}

// NewAggregationProfile returns a AggregationProfile.
func NewAggregationProfile() *AggregationProfile {
	r := &AggregationProfile{}

	return r
}
