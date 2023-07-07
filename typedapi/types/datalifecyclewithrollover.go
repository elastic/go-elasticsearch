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
// https://github.com/elastic/elasticsearch-specification/tree/76e25d34bff1060e300c95f4be468ef88e4f3465

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// DataLifecycleWithRollover type.
//
// https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/indices/_types/DataLifecycle.ts#L31-L48
type DataLifecycleWithRollover struct {
	// DataRetention If defined, every document added to this data stream will be stored at least
	// for this time frame.
	// Any time after this duration the document could be deleted.
	// When empty, every document in this data stream will be stored indefinitely.
	DataRetention Duration `json:"data_retention,omitempty"`
	// Rollover The conditions which will trigger the rollover of a backing index as
	// configured by the cluster setting `cluster.lifecycle.default.rollover`.
	// This property is an implementation detail and it will only be retrieved when
	// the query param `include_defaults` is set to true.
	// The contents of this field are subject to change.
	Rollover *DlmRolloverConditions `json:"rollover,omitempty"`
}

func (s *DataLifecycleWithRollover) UnmarshalJSON(data []byte) error {

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

		case "data_retention":
			if err := dec.Decode(&s.DataRetention); err != nil {
				return err
			}

		case "rollover":
			if err := dec.Decode(&s.Rollover); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataLifecycleWithRollover returns a DataLifecycleWithRollover.
func NewDataLifecycleWithRollover() *DataLifecycleWithRollover {
	r := &DataLifecycleWithRollover{}

	return r
}
