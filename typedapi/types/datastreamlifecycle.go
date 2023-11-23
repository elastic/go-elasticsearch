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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// DataStreamLifecycle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/_types/DataStreamLifecycle.ts#L25-L31
type DataStreamLifecycle struct {
	DataRetention Duration                         `json:"data_retention,omitempty"`
	Downsampling  *DataStreamLifecycleDownsampling `json:"downsampling,omitempty"`
}

func (s *DataStreamLifecycle) UnmarshalJSON(data []byte) error {

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

		case "downsampling":
			if err := dec.Decode(&s.Downsampling); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataStreamLifecycle returns a DataStreamLifecycle.
func NewDataStreamLifecycle() *DataStreamLifecycle {
	r := &DataStreamLifecycle{}

	return r
}
