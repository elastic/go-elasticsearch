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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// IntervalsFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/_types/query_dsl/fulltext.ts#L112-L152
type IntervalsFilter struct {
	// After Query used to return intervals that follow an interval from the `filter`
	// rule.
	After *Intervals `json:"after,omitempty"`
	// Before Query used to return intervals that occur before an interval from the
	// `filter` rule.
	Before *Intervals `json:"before,omitempty"`
	// ContainedBy Query used to return intervals contained by an interval from the `filter`
	// rule.
	ContainedBy *Intervals `json:"contained_by,omitempty"`
	// Containing Query used to return intervals that contain an interval from the `filter`
	// rule.
	Containing *Intervals `json:"containing,omitempty"`
	// NotContainedBy Query used to return intervals that are **not** contained by an interval from
	// the `filter` rule.
	NotContainedBy *Intervals `json:"not_contained_by,omitempty"`
	// NotContaining Query used to return intervals that do **not** contain an interval from the
	// `filter` rule.
	NotContaining *Intervals `json:"not_containing,omitempty"`
	// NotOverlapping Query used to return intervals that do **not** overlap with an interval from
	// the `filter` rule.
	NotOverlapping *Intervals `json:"not_overlapping,omitempty"`
	// Overlapping Query used to return intervals that overlap with an interval from the
	// `filter` rule.
	Overlapping *Intervals `json:"overlapping,omitempty"`
	// Script Script used to return matching documents.
	// This script must return a boolean value: `true` or `false`.
	Script Script `json:"script,omitempty"`
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
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return err
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}

				switch t {

				case "lang", "options", "source":
					o := NewInlineScript()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Script = o

				case "id":
					o := NewStoredScriptId()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Script = o

				}
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
