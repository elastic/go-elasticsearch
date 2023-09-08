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
// https://github.com/elastic/elasticsearch-specification/tree/b89646a75dd9e8001caf92d22bd8b3704c59dfdf

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// FileSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b89646a75dd9e8001caf92d22bd8b3704c59dfdf/specification/nodes/_types/Stats.ts#L286-L291
type FileSystem struct {
	Data      []DataPathStats  `json:"data,omitempty"`
	IoStats   *IoStats         `json:"io_stats,omitempty"`
	Timestamp *int64           `json:"timestamp,omitempty"`
	Total     *FileSystemTotal `json:"total,omitempty"`
}

func (s *FileSystem) UnmarshalJSON(data []byte) error {

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

		case "data":
			if err := dec.Decode(&s.Data); err != nil {
				return err
			}

		case "io_stats":
			if err := dec.Decode(&s.IoStats); err != nil {
				return err
			}

		case "timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Timestamp = &value
			case float64:
				f := int64(v)
				s.Timestamp = &f
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewFileSystem returns a FileSystem.
func NewFileSystem() *FileSystem {
	r := &FileSystem{}

	return r
}
