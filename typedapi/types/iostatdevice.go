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
	"strconv"
)

// IoStatDevice type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/_types/Stats.ts#L298-L305
type IoStatDevice struct {
	DeviceName      *string `json:"device_name,omitempty"`
	Operations      *int64  `json:"operations,omitempty"`
	ReadKilobytes   *int64  `json:"read_kilobytes,omitempty"`
	ReadOperations  *int64  `json:"read_operations,omitempty"`
	WriteKilobytes  *int64  `json:"write_kilobytes,omitempty"`
	WriteOperations *int64  `json:"write_operations,omitempty"`
}

func (s *IoStatDevice) UnmarshalJSON(data []byte) error {

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

		case "device_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DeviceName = &o

		case "operations":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Operations = &value
			case float64:
				f := int64(v)
				s.Operations = &f
			}

		case "read_kilobytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ReadKilobytes = &value
			case float64:
				f := int64(v)
				s.ReadKilobytes = &f
			}

		case "read_operations":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ReadOperations = &value
			case float64:
				f := int64(v)
				s.ReadOperations = &f
			}

		case "write_kilobytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.WriteKilobytes = &value
			case float64:
				f := int64(v)
				s.WriteKilobytes = &f
			}

		case "write_operations":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.WriteOperations = &value
			case float64:
				f := int64(v)
				s.WriteOperations = &f
			}

		}
	}
	return nil
}

// NewIoStatDevice returns a IoStatDevice.
func NewIoStatDevice() *IoStatDevice {
	r := &IoStatDevice{}

	return r
}
