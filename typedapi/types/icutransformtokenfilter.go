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
// https://github.com/elastic/elasticsearch-specification/tree/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icutransformdirection"
)

// IcuTransformTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757/specification/_types/analysis/icu-plugin.ts#L24-L28
type IcuTransformTokenFilter struct {
	Dir     *icutransformdirection.IcuTransformDirection `json:"dir,omitempty"`
	Id      string                                       `json:"id"`
	Type    string                                       `json:"type,omitempty"`
	Version *string                                      `json:"version,omitempty"`
}

func (s *IcuTransformTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "dir":
			if err := dec.Decode(&s.Dir); err != nil {
				return fmt.Errorf("%s | %w", "Dir", err)
			}

		case "id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id = o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s IcuTransformTokenFilter) MarshalJSON() ([]byte, error) {
	type innerIcuTransformTokenFilter IcuTransformTokenFilter
	tmp := innerIcuTransformTokenFilter{
		Dir:     s.Dir,
		Id:      s.Id,
		Type:    s.Type,
		Version: s.Version,
	}

	tmp.Type = "icu_transform"

	return json.Marshal(tmp)
}

// NewIcuTransformTokenFilter returns a IcuTransformTokenFilter.
func NewIcuTransformTokenFilter() *IcuTransformTokenFilter {
	r := &IcuTransformTokenFilter{}

	return r
}
