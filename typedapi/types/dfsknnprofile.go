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

// DfsKnnProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/profile.ts#L181-L186
type DfsKnnProfile struct {
	Collector             []KnnCollectorResult    `json:"collector"`
	Query                 []KnnQueryProfileResult `json:"query"`
	RewriteTime           int64                   `json:"rewrite_time"`
	VectorOperationsCount *int64                  `json:"vector_operations_count,omitempty"`
}

func (s *DfsKnnProfile) UnmarshalJSON(data []byte) error {

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

		case "collector":
			if err := dec.Decode(&s.Collector); err != nil {
				return fmt.Errorf("%s | %w", "Collector", err)
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		case "rewrite_time":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RewriteTime", err)
				}
				s.RewriteTime = value
			case float64:
				f := int64(v)
				s.RewriteTime = f
			}

		case "vector_operations_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "VectorOperationsCount", err)
				}
				s.VectorOperationsCount = &value
			case float64:
				f := int64(v)
				s.VectorOperationsCount = &f
			}

		}
	}
	return nil
}

// NewDfsKnnProfile returns a DfsKnnProfile.
func NewDfsKnnProfile() *DfsKnnProfile {
	r := &DfsKnnProfile{}

	return r
}
