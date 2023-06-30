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

// CompletionContext type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/search/_types/suggester.ts#L159-L166
type CompletionContext struct {
	Boost      *Float64           `json:"boost,omitempty"`
	Context    Context            `json:"context"`
	Neighbours []GeoHashPrecision `json:"neighbours,omitempty"`
	Precision  GeoHashPrecision   `json:"precision,omitempty"`
	Prefix     *bool              `json:"prefix,omitempty"`
}

func (s *CompletionContext) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Context)
		return err
	}

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

		case "boost":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Boost = &f
			case float64:
				f := Float64(v)
				s.Boost = &f
			}

		case "context":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := new(GeoLocation)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Context = *o

			default:
				if err := localDec.Decode(&s.Context); err != nil {
					return err
				}
			}

		case "neighbours":
			if err := dec.Decode(&s.Neighbours); err != nil {
				return err
			}

		case "precision":
			if err := dec.Decode(&s.Precision); err != nil {
				return err
			}

		case "prefix":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Prefix = &value
			case bool:
				s.Prefix = &v
			}

		}
	}
	return nil
}

// NewCompletionContext returns a CompletionContext.
func NewCompletionContext() *CompletionContext {
	r := &CompletionContext{}

	return r
}
