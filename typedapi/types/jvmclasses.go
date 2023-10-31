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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// JvmClasses type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/_types/Stats.ts#L908-L921
type JvmClasses struct {
	// CurrentLoadedCount Number of classes currently loaded by JVM.
	CurrentLoadedCount *int64 `json:"current_loaded_count,omitempty"`
	// TotalLoadedCount Total number of classes loaded since the JVM started.
	TotalLoadedCount *int64 `json:"total_loaded_count,omitempty"`
	// TotalUnloadedCount Total number of classes unloaded since the JVM started.
	TotalUnloadedCount *int64 `json:"total_unloaded_count,omitempty"`
}

func (s *JvmClasses) UnmarshalJSON(data []byte) error {

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

		case "current_loaded_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CurrentLoadedCount = &value
			case float64:
				f := int64(v)
				s.CurrentLoadedCount = &f
			}

		case "total_loaded_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalLoadedCount = &value
			case float64:
				f := int64(v)
				s.TotalLoadedCount = &f
			}

		case "total_unloaded_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.TotalUnloadedCount = &value
			case float64:
				f := int64(v)
				s.TotalUnloadedCount = &f
			}

		}
	}
	return nil
}

// NewJvmClasses returns a JvmClasses.
func NewJvmClasses() *JvmClasses {
	r := &JvmClasses{}

	return r
}
