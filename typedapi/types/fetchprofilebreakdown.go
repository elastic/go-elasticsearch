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

// FetchProfileBreakdown type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/profile.ts#L239-L248
type FetchProfileBreakdown struct {
	LoadSource            *int `json:"load_source,omitempty"`
	LoadSourceCount       *int `json:"load_source_count,omitempty"`
	LoadStoredFields      *int `json:"load_stored_fields,omitempty"`
	LoadStoredFieldsCount *int `json:"load_stored_fields_count,omitempty"`
	NextReader            *int `json:"next_reader,omitempty"`
	NextReaderCount       *int `json:"next_reader_count,omitempty"`
	Process               *int `json:"process,omitempty"`
	ProcessCount          *int `json:"process_count,omitempty"`
}

func (s *FetchProfileBreakdown) UnmarshalJSON(data []byte) error {

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

		case "load_source":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LoadSource", err)
				}
				s.LoadSource = &value
			case float64:
				f := int(v)
				s.LoadSource = &f
			}

		case "load_source_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LoadSourceCount", err)
				}
				s.LoadSourceCount = &value
			case float64:
				f := int(v)
				s.LoadSourceCount = &f
			}

		case "load_stored_fields":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LoadStoredFields", err)
				}
				s.LoadStoredFields = &value
			case float64:
				f := int(v)
				s.LoadStoredFields = &f
			}

		case "load_stored_fields_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LoadStoredFieldsCount", err)
				}
				s.LoadStoredFieldsCount = &value
			case float64:
				f := int(v)
				s.LoadStoredFieldsCount = &f
			}

		case "next_reader":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NextReader", err)
				}
				s.NextReader = &value
			case float64:
				f := int(v)
				s.NextReader = &f
			}

		case "next_reader_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NextReaderCount", err)
				}
				s.NextReaderCount = &value
			case float64:
				f := int(v)
				s.NextReaderCount = &f
			}

		case "process":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Process", err)
				}
				s.Process = &value
			case float64:
				f := int(v)
				s.Process = &f
			}

		case "process_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ProcessCount", err)
				}
				s.ProcessCount = &value
			case float64:
				f := int(v)
				s.ProcessCount = &f
			}

		}
	}
	return nil
}

// NewFetchProfileBreakdown returns a FetchProfileBreakdown.
func NewFetchProfileBreakdown() *FetchProfileBreakdown {
	r := &FetchProfileBreakdown{}

	return r
}
