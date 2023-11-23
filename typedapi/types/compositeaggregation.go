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

// CompositeAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/aggregations/bucket.ts#L120-L136
type CompositeAggregation struct {
	// After When paginating, use the `after_key` value returned in the previous response
	// to retrieve the next page.
	After CompositeAggregateKey `json:"after,omitempty"`
	Meta  Metadata              `json:"meta,omitempty"`
	Name  *string               `json:"name,omitempty"`
	// Size The number of composite buckets that should be returned.
	Size *int `json:"size,omitempty"`
	// Sources The value sources used to build composite buckets.
	// Keys are returned in the order of the `sources` definition.
	Sources []map[string]CompositeAggregationSource `json:"sources,omitempty"`
}

func (s *CompositeAggregation) UnmarshalJSON(data []byte) error {

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

		case "after":
			if err := dec.Decode(&s.After); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "sources":
			if err := dec.Decode(&s.Sources); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCompositeAggregation returns a CompositeAggregation.
func NewCompositeAggregation() *CompositeAggregation {
	r := &CompositeAggregation{}

	return r
}
