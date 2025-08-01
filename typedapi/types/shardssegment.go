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

// ShardsSegment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/segments/types.ts#L46-L51
type ShardsSegment struct {
	NumCommittedSegments int                 `json:"num_committed_segments"`
	NumSearchSegments    int                 `json:"num_search_segments"`
	Routing              ShardSegmentRouting `json:"routing"`
	Segments             map[string]Segment  `json:"segments"`
}

func (s *ShardsSegment) UnmarshalJSON(data []byte) error {

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

		case "num_committed_segments":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumCommittedSegments", err)
				}
				s.NumCommittedSegments = value
			case float64:
				f := int(v)
				s.NumCommittedSegments = f
			}

		case "num_search_segments":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumSearchSegments", err)
				}
				s.NumSearchSegments = value
			case float64:
				f := int(v)
				s.NumSearchSegments = f
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "segments":
			if s.Segments == nil {
				s.Segments = make(map[string]Segment, 0)
			}
			if err := dec.Decode(&s.Segments); err != nil {
				return fmt.Errorf("%s | %w", "Segments", err)
			}

		}
	}
	return nil
}

// NewShardsSegment returns a ShardsSegment.
func NewShardsSegment() *ShardsSegment {
	r := &ShardsSegment{
		Segments: make(map[string]Segment),
	}

	return r
}
