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

// UsageStatsShards type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/indices/field_usage_stats/IndicesFieldUsageStatsResponse.ts#L42-L47
type UsageStatsShards struct {
	Routing                 ShardRouting       `json:"routing"`
	Stats                   IndicesShardsStats `json:"stats"`
	TrackingId              string             `json:"tracking_id"`
	TrackingStartedAtMillis int64              `json:"tracking_started_at_millis"`
}

func (s *UsageStatsShards) UnmarshalJSON(data []byte) error {

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

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "stats":
			if err := dec.Decode(&s.Stats); err != nil {
				return err
			}

		case "tracking_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TrackingId = o

		case "tracking_started_at_millis":
			if err := dec.Decode(&s.TrackingStartedAtMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewUsageStatsShards returns a UsageStatsShards.
func NewUsageStatsShards() *UsageStatsShards {
	r := &UsageStatsShards{}

	return r
}
