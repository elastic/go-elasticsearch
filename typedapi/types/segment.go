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

// Segment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/segments/types.ts#L28-L38
type Segment struct {
	Attributes  map[string]string `json:"attributes"`
	Committed   bool              `json:"committed"`
	Compound    bool              `json:"compound"`
	DeletedDocs int64             `json:"deleted_docs"`
	Generation  int               `json:"generation"`
	NumDocs     int64             `json:"num_docs"`
	Search      bool              `json:"search"`
	SizeInBytes Float64           `json:"size_in_bytes"`
	Version     string            `json:"version"`
}

func (s *Segment) UnmarshalJSON(data []byte) error {

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

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return fmt.Errorf("%s | %w", "Attributes", err)
			}

		case "committed":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Committed", err)
				}
				s.Committed = value
			case bool:
				s.Committed = v
			}

		case "compound":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Compound", err)
				}
				s.Compound = value
			case bool:
				s.Compound = v
			}

		case "deleted_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DeletedDocs", err)
				}
				s.DeletedDocs = value
			case float64:
				f := int64(v)
				s.DeletedDocs = f
			}

		case "generation":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Generation", err)
				}
				s.Generation = value
			case float64:
				f := int(v)
				s.Generation = f
			}

		case "num_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumDocs", err)
				}
				s.NumDocs = value
			case float64:
				f := int64(v)
				s.NumDocs = f
			}

		case "search":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Search", err)
				}
				s.Search = value
			case bool:
				s.Search = v
			}

		case "size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SizeInBytes", err)
				}
				f := Float64(value)
				s.SizeInBytes = f
			case float64:
				f := Float64(v)
				s.SizeInBytes = f
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewSegment returns a Segment.
func NewSegment() *Segment {
	r := &Segment{
		Attributes: make(map[string]string),
	}

	return r
}
