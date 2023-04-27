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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// DataTiers type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/xpack/usage/types.ts#L333-L340
type DataTiers struct {
	Available   bool                     `json:"available"`
	DataCold    DataTierPhaseStatistics  `json:"data_cold"`
	DataContent DataTierPhaseStatistics  `json:"data_content"`
	DataFrozen  *DataTierPhaseStatistics `json:"data_frozen,omitempty"`
	DataHot     DataTierPhaseStatistics  `json:"data_hot"`
	DataWarm    DataTierPhaseStatistics  `json:"data_warm"`
	Enabled     bool                     `json:"enabled"`
}

func (s *DataTiers) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "data_cold":
			if err := dec.Decode(&s.DataCold); err != nil {
				return err
			}

		case "data_content":
			if err := dec.Decode(&s.DataContent); err != nil {
				return err
			}

		case "data_frozen":
			if err := dec.Decode(&s.DataFrozen); err != nil {
				return err
			}

		case "data_hot":
			if err := dec.Decode(&s.DataHot); err != nil {
				return err
			}

		case "data_warm":
			if err := dec.Decode(&s.DataWarm); err != nil {
				return err
			}

		case "enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		}
	}
	return nil
}

// NewDataTiers returns a DataTiers.
func NewDataTiers() *DataTiers {
	r := &DataTiers{}

	return r
}
