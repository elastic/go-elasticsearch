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

// BulkIndexByScrollFailure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/Errors.ts#L58-L64
type BulkIndexByScrollFailure struct {
	Cause  ErrorCause `json:"cause"`
	Id     string     `json:"id"`
	Index  string     `json:"index"`
	Status int        `json:"status"`
	Type   string     `json:"type"`
}

func (s *BulkIndexByScrollFailure) UnmarshalJSON(data []byte) error {

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

		case "cause":
			if err := dec.Decode(&s.Cause); err != nil {
				return err
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "status":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Status = value
			case float64:
				f := int(v)
				s.Status = f
			}

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

		}
	}
	return nil
}

// NewBulkIndexByScrollFailure returns a BulkIndexByScrollFailure.
func NewBulkIndexByScrollFailure() *BulkIndexByScrollFailure {
	r := &BulkIndexByScrollFailure{}

	return r
}
