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

// AdaptiveSelection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L403-L432
type AdaptiveSelection struct {
	// AvgQueueSize The exponentially weighted moving average queue size of search requests on
	// the keyed node.
	AvgQueueSize *int64 `json:"avg_queue_size,omitempty"`
	// AvgResponseTime The exponentially weighted moving average response time of search requests on
	// the keyed node.
	AvgResponseTime Duration `json:"avg_response_time,omitempty"`
	// AvgResponseTimeNs The exponentially weighted moving average response time, in nanoseconds, of
	// search requests on the keyed node.
	AvgResponseTimeNs *int64 `json:"avg_response_time_ns,omitempty"`
	// AvgServiceTime The exponentially weighted moving average service time of search requests on
	// the keyed node.
	AvgServiceTime Duration `json:"avg_service_time,omitempty"`
	// AvgServiceTimeNs The exponentially weighted moving average service time, in nanoseconds, of
	// search requests on the keyed node.
	AvgServiceTimeNs *int64 `json:"avg_service_time_ns,omitempty"`
	// OutgoingSearches The number of outstanding search requests to the keyed node from the node
	// these stats are for.
	OutgoingSearches *int64 `json:"outgoing_searches,omitempty"`
	// Rank The rank of this node; used for shard selection when routing search requests.
	Rank *string `json:"rank,omitempty"`
}

func (s *AdaptiveSelection) UnmarshalJSON(data []byte) error {

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

		case "avg_queue_size":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AvgQueueSize = &value
			case float64:
				f := int64(v)
				s.AvgQueueSize = &f
			}

		case "avg_response_time":
			if err := dec.Decode(&s.AvgResponseTime); err != nil {
				return err
			}

		case "avg_response_time_ns":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AvgResponseTimeNs = &value
			case float64:
				f := int64(v)
				s.AvgResponseTimeNs = &f
			}

		case "avg_service_time":
			if err := dec.Decode(&s.AvgServiceTime); err != nil {
				return err
			}

		case "avg_service_time_ns":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AvgServiceTimeNs = &value
			case float64:
				f := int64(v)
				s.AvgServiceTimeNs = &f
			}

		case "outgoing_searches":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.OutgoingSearches = &value
			case float64:
				f := int64(v)
				s.OutgoingSearches = &f
			}

		case "rank":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Rank = &o

		}
	}
	return nil
}

// NewAdaptiveSelection returns a AdaptiveSelection.
func NewAdaptiveSelection() *AdaptiveSelection {
	r := &AdaptiveSelection{}

	return r
}
