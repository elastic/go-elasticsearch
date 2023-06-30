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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// IntervalsFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/fulltext.ts#L74-L86
type IntervalsFilter struct {
	After          *Intervals `json:"after,omitempty"`
	Before         *Intervals `json:"before,omitempty"`
	ContainedBy    *Intervals `json:"contained_by,omitempty"`
	Containing     *Intervals `json:"containing,omitempty"`
	NotContainedBy *Intervals `json:"not_contained_by,omitempty"`
	NotContaining  *Intervals `json:"not_containing,omitempty"`
	NotOverlapping *Intervals `json:"not_overlapping,omitempty"`
	Overlapping    *Intervals `json:"overlapping,omitempty"`
	Script         Script     `json:"script,omitempty"`
}

func (s *IntervalsFilter) UnmarshalJSON(data []byte) error {

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

		case "before":
			if err := dec.Decode(&s.Before); err != nil {
				return err
			}

		case "contained_by":
			if err := dec.Decode(&s.ContainedBy); err != nil {
				return err
			}

		case "containing":
			if err := dec.Decode(&s.Containing); err != nil {
				return err
			}

		case "not_contained_by":
			if err := dec.Decode(&s.NotContainedBy); err != nil {
				return err
			}

		case "not_containing":
			if err := dec.Decode(&s.NotContaining); err != nil {
				return err
			}

		case "not_overlapping":
			if err := dec.Decode(&s.NotOverlapping); err != nil {
				return err
			}

		case "overlapping":
			if err := dec.Decode(&s.Overlapping); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIntervalsFilter returns a IntervalsFilter.
func NewIntervalsFilter() *IntervalsFilter {
	r := &IntervalsFilter{}

	return r
}
