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

// QueryWatch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/watcher/_types/Watch.ts#L58-L64
type QueryWatch struct {
	Id_          string       `json:"_id"`
	PrimaryTerm_ *int         `json:"_primary_term,omitempty"`
	SeqNo_       *int64       `json:"_seq_no,omitempty"`
	Status       *WatchStatus `json:"status,omitempty"`
	Watch        *Watch       `json:"watch,omitempty"`
}

func (s *QueryWatch) UnmarshalJSON(data []byte) error {

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

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return err
			}

		case "_primary_term":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PrimaryTerm_ = &value
			case float64:
				f := int(v)
				s.PrimaryTerm_ = &f
			}

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return err
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return err
			}

		case "watch":
			if err := dec.Decode(&s.Watch); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewQueryWatch returns a QueryWatch.
func NewQueryWatch() *QueryWatch {
	r := &QueryWatch{}

	return r
}
