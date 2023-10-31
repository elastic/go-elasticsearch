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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
)

// MultisearchHeader type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/msearch/types.ts#L53-L68
type MultisearchHeader struct {
	AllowNoIndices            *bool                           `json:"allow_no_indices,omitempty"`
	AllowPartialSearchResults *bool                           `json:"allow_partial_search_results,omitempty"`
	CcsMinimizeRoundtrips     *bool                           `json:"ccs_minimize_roundtrips,omitempty"`
	ExpandWildcards           []expandwildcard.ExpandWildcard `json:"expand_wildcards,omitempty"`
	IgnoreThrottled           *bool                           `json:"ignore_throttled,omitempty"`
	IgnoreUnavailable         *bool                           `json:"ignore_unavailable,omitempty"`
	Index                     []string                        `json:"index,omitempty"`
	Preference                *string                         `json:"preference,omitempty"`
	RequestCache              *bool                           `json:"request_cache,omitempty"`
	Routing                   *string                         `json:"routing,omitempty"`
	SearchType                *searchtype.SearchType          `json:"search_type,omitempty"`
}

func (s *MultisearchHeader) UnmarshalJSON(data []byte) error {

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

		case "allow_no_indices":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowNoIndices = &value
			case bool:
				s.AllowNoIndices = &v
			}

		case "allow_partial_search_results":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowPartialSearchResults = &value
			case bool:
				s.AllowPartialSearchResults = &v
			}

		case "ccs_minimize_roundtrips":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.CcsMinimizeRoundtrips = &value
			case bool:
				s.CcsMinimizeRoundtrips = &v
			}

		case "expand_wildcards":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := &expandwildcard.ExpandWildcard{}
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.ExpandWildcards = append(s.ExpandWildcards, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.ExpandWildcards); err != nil {
					return err
				}
			}

		case "ignore_throttled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreThrottled = &value
			case bool:
				s.IgnoreThrottled = &v
			}

		case "ignore_unavailable":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreUnavailable = &value
			case bool:
				s.IgnoreUnavailable = &v
			}

		case "index":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Index = append(s.Index, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Index); err != nil {
					return err
				}
			}

		case "preference":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Preference = &o

		case "request_cache":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.RequestCache = &value
			case bool:
				s.RequestCache = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "search_type":
			if err := dec.Decode(&s.SearchType); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMultisearchHeader returns a MultisearchHeader.
func NewMultisearchHeader() *MultisearchHeader {
	r := &MultisearchHeader{}

	return r
}
