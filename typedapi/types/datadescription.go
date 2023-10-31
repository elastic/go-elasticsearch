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

// DataDescription type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Job.ts#L374-L390
type DataDescription struct {
	FieldDelimiter *string `json:"field_delimiter,omitempty"`
	// Format Only JSON format is supported at this time.
	Format *string `json:"format,omitempty"`
	// TimeField The name of the field that contains the timestamp.
	TimeField *string `json:"time_field,omitempty"`
	// TimeFormat The time format, which can be `epoch`, `epoch_ms`, or a custom pattern. The
	// value `epoch` refers to UNIX or Epoch time (the number of seconds since 1 Jan
	// 1970). The value `epoch_ms` indicates that time is measured in milliseconds
	// since the epoch. The `epoch` and `epoch_ms` time formats accept either
	// integer or real values. Custom patterns must conform to the Java
	// DateTimeFormatter class. When you use date-time formatting patterns, it is
	// recommended that you provide the full date, time and time zone. For example:
	// `yyyy-MM-dd'T'HH:mm:ssX`. If the pattern that you specify is not sufficient
	// to produce a complete timestamp, job creation fails.
	TimeFormat *string `json:"time_format,omitempty"`
}

func (s *DataDescription) UnmarshalJSON(data []byte) error {

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

		case "field_delimiter":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FieldDelimiter = &o

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "time_field":
			if err := dec.Decode(&s.TimeField); err != nil {
				return err
			}

		case "time_format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TimeFormat = &o

		}
	}
	return nil
}

// NewDataDescription returns a DataDescription.
func NewDataDescription() *DataDescription {
	r := &DataDescription{}

	return r
}
