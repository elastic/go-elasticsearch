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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// IngestTotal type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/_types/Stats.ts#L398-L415
type IngestTotal struct {
	// Count Total number of documents ingested during the lifetime of this node.
	Count int64 `json:"count"`
	// Current Total number of documents currently being ingested.
	Current int64 `json:"current"`
	// Failed Total number of failed ingest operations during the lifetime of this node.
	Failed int64 `json:"failed"`
	// TimeInMillis Total time, in milliseconds, spent preprocessing ingest documents during the
	// lifetime of this node.
	TimeInMillis int64 `json:"time_in_millis"`
}

func (s *IngestTotal) UnmarshalJSON(data []byte) error {

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

		case "current":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Current", err)
				}
				s.Current = value
			case float64:
				f := int64(v)
				s.Current = f
			}

		case "failed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Failed", err)
				}
				s.Failed = value
			case float64:
				f := int64(v)
				s.Failed = f
			}

		case "time_in_millis":
			if err := dec.Decode(&s.TimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TimeInMillis", err)
			}

		}
	}
	return nil
}

// NewIngestTotal returns a IngestTotal.
func NewIngestTotal() *IngestTotal {
	r := &IngestTotal{}

	return r
}
