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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TransportHistogram type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/_types/Stats.ts#L1170-L1184
type TransportHistogram struct {
	// Count The number of times a transport thread took a period of time within the
	// bounds of this bucket to handle an inbound message.
	Count *int64 `json:"count,omitempty"`
	// GeMillis The inclusive lower bound of the bucket in milliseconds. May be omitted on
	// the first bucket if this bucket has no lower bound.
	GeMillis *int64 `json:"ge_millis,omitempty"`
	// LtMillis The exclusive upper bound of the bucket in milliseconds.
	// May be omitted on the last bucket if this bucket has no upper bound.
	LtMillis *int64 `json:"lt_millis,omitempty"`
}

func (s *TransportHistogram) UnmarshalJSON(data []byte) error {

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
				s.Count = &value
			case float64:
				f := int64(v)
				s.Count = &f
			}

		case "ge_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "GeMillis", err)
				}
				s.GeMillis = &value
			case float64:
				f := int64(v)
				s.GeMillis = &f
			}

		case "lt_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LtMillis", err)
				}
				s.LtMillis = &value
			case float64:
				f := int64(v)
				s.LtMillis = &f
			}

		}
	}
	return nil
}

// NewTransportHistogram returns a TransportHistogram.
func NewTransportHistogram() *TransportHistogram {
	r := &TransportHistogram{}

	return r
}
