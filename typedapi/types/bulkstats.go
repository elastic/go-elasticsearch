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

// BulkStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/Stats.ts#L68-L78
type BulkStats struct {
	AvgSize           ByteSize `json:"avg_size,omitempty"`
	AvgSizeInBytes    int64    `json:"avg_size_in_bytes"`
	AvgTime           Duration `json:"avg_time,omitempty"`
	AvgTimeInMillis   int64    `json:"avg_time_in_millis"`
	TotalOperations   int64    `json:"total_operations"`
	TotalSize         ByteSize `json:"total_size,omitempty"`
	TotalSizeInBytes  int64    `json:"total_size_in_bytes"`
	TotalTime         Duration `json:"total_time,omitempty"`
	TotalTimeInMillis int64    `json:"total_time_in_millis"`
}

func (s *BulkStats) UnmarshalJSON(data []byte) error {

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

		case "avg_size":
			if err := dec.Decode(&s.AvgSize); err != nil {
				return err
			}

		case "avg_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AvgSizeInBytes = value
			case float64:
				f := int64(v)
				s.AvgSizeInBytes = f
			}

		case "avg_time":
			if err := dec.Decode(&s.AvgTime); err != nil {
				return err
			}

		case "avg_time_in_millis":
			if err := dec.Decode(&s.AvgTimeInMillis); err != nil {
				return err
			}

		case "total_operations":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalOperations = value
			case float64:
				f := int64(v)
				s.TotalOperations = f
			}

		case "total_size":
			if err := dec.Decode(&s.TotalSize); err != nil {
				return err
			}

		case "total_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalSizeInBytes = value
			case float64:
				f := int64(v)
				s.TotalSizeInBytes = f
			}

		case "total_time":
			if err := dec.Decode(&s.TotalTime); err != nil {
				return err
			}

		case "total_time_in_millis":
			if err := dec.Decode(&s.TotalTimeInMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewBulkStats returns a BulkStats.
func NewBulkStats() *BulkStats {
	r := &BulkStats{}

	return r
}
