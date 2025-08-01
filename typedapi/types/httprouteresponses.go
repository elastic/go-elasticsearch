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

// HttpRouteResponses type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/_types/Stats.ts#L703-L708
type HttpRouteResponses struct {
	Count                 int64               `json:"count"`
	HandlingTimeHistogram []TimeHttpHistogram `json:"handling_time_histogram"`
	SizeHistogram         []SizeHttpHistogram `json:"size_histogram"`
	TotalSizeInBytes      int64               `json:"total_size_in_bytes"`
}

func (s *HttpRouteResponses) UnmarshalJSON(data []byte) error {

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

		case "handling_time_histogram":
			if err := dec.Decode(&s.HandlingTimeHistogram); err != nil {
				return fmt.Errorf("%s | %w", "HandlingTimeHistogram", err)
			}

		case "size_histogram":
			if err := dec.Decode(&s.SizeHistogram); err != nil {
				return fmt.Errorf("%s | %w", "SizeHistogram", err)
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

// NewHttpRouteResponses returns a HttpRouteResponses.
func NewHttpRouteResponses() *HttpRouteResponses {
	r := &HttpRouteResponses{}

	return r
}
