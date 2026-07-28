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

// TransportMessageSizeHistogramBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7fd0bd13eaf28bd179fc906f57da09e852eb818e/specification/nodes/_types/Stats.ts#L1431-L1456
type TransportMessageSizeHistogramBucket struct {
	// Count The number of messages with a size that falls within the bounds of this
	// bucket.
	Count int64 `json:"count"`
	// Ge The inclusive lower bound of the bucket. May be omitted on the first bucket
	// if this bucket has no lower bound.
	Ge ByteSize `json:"ge,omitempty"`
	// GeBytes The inclusive lower bound of the bucket in bytes. May be omitted on the first
	// bucket if this bucket has no lower bound.
	GeBytes *int64 `json:"ge_bytes,omitempty"`
	// Lt The exclusive upper bound of the bucket. May be omitted on the last bucket if
	// this bucket has no upper bound.
	Lt ByteSize `json:"lt,omitempty"`
	// LtBytes The exclusive upper bound of the bucket in bytes. May be omitted on the last
	// bucket if this bucket has no upper bound.
	LtBytes *int64 `json:"lt_bytes,omitempty"`
}

func (s *TransportMessageSizeHistogramBucket) UnmarshalJSON(data []byte) error {

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

		case "ge":
			if err := dec.Decode(&s.Ge); err != nil {
				return fmt.Errorf("%s | %w", "Ge", err)
			}

		case "ge_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "GeBytes", err)
				}
				s.GeBytes = &value
			case float64:
				f := int64(v)
				s.GeBytes = &f
			}

		case "lt":
			if err := dec.Decode(&s.Lt); err != nil {
				return fmt.Errorf("%s | %w", "Lt", err)
			}

		case "lt_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LtBytes", err)
				}
				s.LtBytes = &value
			case float64:
				f := int64(v)
				s.LtBytes = &f
			}

		}
	}
	return nil
}

// NewTransportMessageSizeHistogramBucket returns a TransportMessageSizeHistogramBucket.
func NewTransportMessageSizeHistogramBucket() *TransportMessageSizeHistogramBucket {
	r := &TransportMessageSizeHistogramBucket{}

	return r
}
