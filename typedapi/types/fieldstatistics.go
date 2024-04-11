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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// FieldStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_global/termvectors/types.ts#L28-L32
type FieldStatistics struct {
	DocCount   int   `json:"doc_count"`
	SumDocFreq int64 `json:"sum_doc_freq"`
	SumTtf     int64 `json:"sum_ttf"`
}

func (s *FieldStatistics) UnmarshalJSON(data []byte) error {

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

		case "doc_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DocCount", err)
				}
				s.DocCount = value
			case float64:
				f := int(v)
				s.DocCount = f
			}

		case "sum_doc_freq":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SumDocFreq", err)
				}
				s.SumDocFreq = value
			case float64:
				f := int64(v)
				s.SumDocFreq = f
			}

		case "sum_ttf":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SumTtf", err)
				}
				s.SumTtf = value
			case float64:
				f := int64(v)
				s.SumTtf = f
			}

		}
	}
	return nil
}

// NewFieldStatistics returns a FieldStatistics.
func NewFieldStatistics() *FieldStatistics {
	r := &FieldStatistics{}

	return r
}
