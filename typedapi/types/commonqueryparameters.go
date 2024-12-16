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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CommonQueryParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/_spec_utils/behaviors.ts#L50-L84
type CommonQueryParameters struct {
	// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
	// when they occur.
	ErrorTrace *bool `json:"error_trace,omitempty"`
	// FilterPath Comma-separated list of filters in dot notation which reduce the response
	// returned by Elasticsearch.
	FilterPath []string `json:"filter_path,omitempty"`
	// Human When set to `true` will return statistics in a format suitable for humans.
	// For example `"exists_time": "1h"` for humans and
	// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
	// readable values will be omitted. This makes sense for responses being
	// consumed
	// only by machines.
	Human *bool `json:"human,omitempty"`
	// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
	// this option for debugging only.
	Pretty *bool `json:"pretty,omitempty"`
}

func (s *CommonQueryParameters) UnmarshalJSON(data []byte) error {

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

		case "error_trace":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ErrorTrace", err)
				}
				s.ErrorTrace = &value
			case bool:
				s.ErrorTrace = &v
			}

		case "filter_path":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "FilterPath", err)
				}

				s.FilterPath = append(s.FilterPath, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.FilterPath); err != nil {
					return fmt.Errorf("%s | %w", "FilterPath", err)
				}
			}

		case "human":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Human", err)
				}
				s.Human = &value
			case bool:
				s.Human = &v
			}

		case "pretty":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Pretty", err)
				}
				s.Pretty = &value
			case bool:
				s.Pretty = &v
			}

		}
	}
	return nil
}

// NewCommonQueryParameters returns a CommonQueryParameters.
func NewCommonQueryParameters() *CommonQueryParameters {
	r := &CommonQueryParameters{}

	return r
}
