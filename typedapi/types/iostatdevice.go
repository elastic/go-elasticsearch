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
	"strconv"
)

// IoStatDevice type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L730-L755
type IoStatDevice struct {
	// DeviceName The Linux device name.
	DeviceName *string `json:"device_name,omitempty"`
	// Operations The total number of read and write operations for the device completed since
	// starting Elasticsearch.
	Operations *int64 `json:"operations,omitempty"`
	// ReadKilobytes The total number of kilobytes read for the device since starting
	// Elasticsearch.
	ReadKilobytes *int64 `json:"read_kilobytes,omitempty"`
	// ReadOperations The total number of read operations for the device completed since starting
	// Elasticsearch.
	ReadOperations *int64 `json:"read_operations,omitempty"`
	// WriteKilobytes The total number of kilobytes written for the device since starting
	// Elasticsearch.
	WriteKilobytes *int64 `json:"write_kilobytes,omitempty"`
	// WriteOperations The total number of write operations for the device completed since starting
	// Elasticsearch.
	WriteOperations *int64 `json:"write_operations,omitempty"`
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
