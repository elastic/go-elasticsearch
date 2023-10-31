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
)

// TimeSync type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/transform/_types/Transform.ts#L177-L189
type TimeSync struct {
	// Delay The time delay between the current time and the latest input data time.
	Delay Duration `json:"delay,omitempty"`
	// Field The date field that is used to identify new documents in the source. In
	// general, it’s a good idea to use a field
	// that contains the ingest timestamp. If you use a different field, you might
	// need to set the delay such that it
	// accounts for data transmission delays.
	Field string `json:"field"`
}

func (s *TimeSync) UnmarshalJSON(data []byte) error {

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

		case "delay":
			if err := dec.Decode(&s.Delay); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTimeSync returns a TimeSync.
func NewTimeSync() *TimeSync {
	r := &TimeSync{}

	return r
}
