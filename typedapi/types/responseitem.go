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
// https://github.com/elastic/elasticsearch-specification/tree/3b09f9d8e90178243f8a340a7bc324aab152c602

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ResponseItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3b09f9d8e90178243f8a340a7bc324aab152c602/specification/_global/bulk/types.ts#L37-L49
type ResponseItem struct {
	Error         *ErrorCause               `json:"error,omitempty"`
	ForcedRefresh *bool                     `json:"forced_refresh,omitempty"`
	Get           *InlineGetDictUserDefined `json:"get,omitempty"`
	Id_           string                    `json:"_id,omitempty"`
	Index_        string                    `json:"_index"`
	PrimaryTerm_  *int64                    `json:"_primary_term,omitempty"`
	Result        *string                   `json:"result,omitempty"`
	SeqNo_        *int64                    `json:"_seq_no,omitempty"`
	Shards_       *ShardStatistics          `json:"_shards,omitempty"`
	Status        int                       `json:"status"`
	Version_      *int64                    `json:"_version,omitempty"`
}

func (s *ResponseItem) UnmarshalJSON(data []byte) error {

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

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return err
			}

		case "forced_refresh":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ForcedRefresh = &value
			case bool:
				s.ForcedRefresh = &v
			}

		case "get":
			if err := dec.Decode(&s.Get); err != nil {
				return err
			}

		case "_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id_ = o

		case "_index":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Index_ = o

		case "_primary_term":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PrimaryTerm_ = &value
			case float64:
				f := int64(v)
				s.PrimaryTerm_ = &f
			}

		case "result":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Result = &o

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return err
			}

		case "_shards":
			if err := dec.Decode(&s.Shards_); err != nil {
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

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewResponseItem returns a ResponseItem.
func NewResponseItem() *ResponseItem {
	r := &ResponseItem{}

	return r
}
