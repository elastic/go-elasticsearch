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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// SignificantTermsAggregateBaseSignificantLongTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/aggregations/Aggregate.ts#L580-L585
type SignificantTermsAggregateBaseSignificantLongTermsBucket struct {
	BgCount  *int64                            `json:"bg_count,omitempty"`
	Buckets  BucketsSignificantLongTermsBucket `json:"buckets"`
	DocCount *int64                            `json:"doc_count,omitempty"`
	Meta     Metadata                          `json:"meta,omitempty"`
}

func (s *SignificantTermsAggregateBaseSignificantLongTermsBucket) UnmarshalJSON(data []byte) error {

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

		case "bg_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.BgCount = &value
			case float64:
				f := int64(v)
				s.BgCount = &f
			}

		case "buckets":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]SignificantLongTermsBucket, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Buckets = o
			case '[':
				o := []SignificantLongTermsBucket{}
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Buckets = o
			}

		case "doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocCount = &value
			case float64:
				f := int64(v)
				s.DocCount = &f
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSignificantTermsAggregateBaseSignificantLongTermsBucket returns a SignificantTermsAggregateBaseSignificantLongTermsBucket.
func NewSignificantTermsAggregateBaseSignificantLongTermsBucket() *SignificantTermsAggregateBaseSignificantLongTermsBucket {
	r := &SignificantTermsAggregateBaseSignificantLongTermsBucket{}

	return r
}
