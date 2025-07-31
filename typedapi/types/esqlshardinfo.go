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

// EsqlShardInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/esql/_types/EsqlResult.ts#L90-L96
type EsqlShardInfo struct {
	Failed     *int               `json:"failed,omitempty"`
	Failures   []EsqlShardFailure `json:"failures,omitempty"`
	Skipped    *int               `json:"skipped,omitempty"`
	Successful *int               `json:"successful,omitempty"`
	Total      int                `json:"total"`
}

func (s *EsqlShardInfo) UnmarshalJSON(data []byte) error {

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

		case "failed":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Failed", err)
				}
				s.Failed = &value
			case float64:
				f := int(v)
				s.Failed = &f
			}

		case "failures":
			if err := dec.Decode(&s.Failures); err != nil {
				return fmt.Errorf("%s | %w", "Failures", err)
			}

		case "skipped":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Skipped", err)
				}
				s.Skipped = &value
			case float64:
				f := int(v)
				s.Skipped = &f
			}

		case "successful":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Successful", err)
				}
				s.Successful = &value
			case float64:
				f := int(v)
				s.Successful = &f
			}

		case "total":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewEsqlShardInfo returns a EsqlShardInfo.
func NewEsqlShardInfo() *EsqlShardInfo {
	r := &EsqlShardInfo{}

	return r
}
