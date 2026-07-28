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
// https://github.com/elastic/elasticsearch-specification/tree/7fd0bd13eaf28bd179fc906f57da09e852eb818e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TransportActionMessageStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7fd0bd13eaf28bd179fc906f57da09e852eb818e/specification/nodes/_types/Stats.ts#L1412-L1429
type TransportActionMessageStats struct {
	// Count The number of messages of this kind that the node has handled for this
	// action.
	Count int64 `json:"count"`
	// Histogram The distribution of the sizes of the messages of this kind that the node has
	// handled for this action, represented as a histogram.
	Histogram []TransportMessageSizeHistogramBucket `json:"histogram"`
	// TotalSize The cumulative size of the messages of this kind that the node has handled
	// for this action.
	TotalSize ByteSize `json:"total_size,omitempty"`
	// TotalSizeInBytes The cumulative size, in bytes, of the messages of this kind that the node has
	// handled for this action.
	TotalSizeInBytes int64 `json:"total_size_in_bytes"`
}

func (s *TransportActionMessageStats) UnmarshalJSON(data []byte) error {

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

		case "count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "histogram":
			if err := dec.Decode(&s.Histogram); err != nil {
				return fmt.Errorf("%s | %w", "Histogram", err)
			}

		case "total_size":
			if err := dec.Decode(&s.TotalSize); err != nil {
				return fmt.Errorf("%s | %w", "TotalSize", err)
			}

		case "total_size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalSizeInBytes", err)
				}
				s.TotalSizeInBytes = value
			case float64:
				f := int64(v)
				s.TotalSizeInBytes = f
			}

		}
	}
	return nil
}

// NewTransportActionMessageStats returns a TransportActionMessageStats.
func NewTransportActionMessageStats() *TransportActionMessageStats {
	r := &TransportActionMessageStats{}

	return r
}
