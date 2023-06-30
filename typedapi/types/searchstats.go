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

// SearchStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/Stats.ts#L185-L204
type SearchStats struct {
	FetchCurrent        int64                  `json:"fetch_current"`
	FetchTime           Duration               `json:"fetch_time,omitempty"`
	FetchTimeInMillis   int64                  `json:"fetch_time_in_millis"`
	FetchTotal          int64                  `json:"fetch_total"`
	Groups              map[string]SearchStats `json:"groups,omitempty"`
	OpenContexts        *int64                 `json:"open_contexts,omitempty"`
	QueryCurrent        int64                  `json:"query_current"`
	QueryTime           Duration               `json:"query_time,omitempty"`
	QueryTimeInMillis   int64                  `json:"query_time_in_millis"`
	QueryTotal          int64                  `json:"query_total"`
	ScrollCurrent       int64                  `json:"scroll_current"`
	ScrollTime          Duration               `json:"scroll_time,omitempty"`
	ScrollTimeInMillis  int64                  `json:"scroll_time_in_millis"`
	ScrollTotal         int64                  `json:"scroll_total"`
	SuggestCurrent      int64                  `json:"suggest_current"`
	SuggestTime         Duration               `json:"suggest_time,omitempty"`
	SuggestTimeInMillis int64                  `json:"suggest_time_in_millis"`
	SuggestTotal        int64                  `json:"suggest_total"`
}

func (s *SearchStats) UnmarshalJSON(data []byte) error {

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

		case "fetch_current":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FetchCurrent = value
			case float64:
				f := int64(v)
				s.FetchCurrent = f
			}

		case "fetch_time":
			if err := dec.Decode(&s.FetchTime); err != nil {
				return err
			}

		case "fetch_time_in_millis":
			if err := dec.Decode(&s.FetchTimeInMillis); err != nil {
				return err
			}

		case "fetch_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FetchTotal = value
			case float64:
				f := int64(v)
				s.FetchTotal = f
			}

		case "groups":
			if s.Groups == nil {
				s.Groups = make(map[string]SearchStats, 0)
			}
			if err := dec.Decode(&s.Groups); err != nil {
				return err
			}

		case "open_contexts":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.OpenContexts = &value
			case float64:
				f := int64(v)
				s.OpenContexts = &f
			}

		case "query_current":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.QueryCurrent = value
			case float64:
				f := int64(v)
				s.QueryCurrent = f
			}

		case "query_time":
			if err := dec.Decode(&s.QueryTime); err != nil {
				return err
			}

		case "query_time_in_millis":
			if err := dec.Decode(&s.QueryTimeInMillis); err != nil {
				return err
			}

		case "query_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.QueryTotal = value
			case float64:
				f := int64(v)
				s.QueryTotal = f
			}

		case "scroll_current":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ScrollCurrent = value
			case float64:
				f := int64(v)
				s.ScrollCurrent = f
			}

		case "scroll_time":
			if err := dec.Decode(&s.ScrollTime); err != nil {
				return err
			}

		case "scroll_time_in_millis":
			if err := dec.Decode(&s.ScrollTimeInMillis); err != nil {
				return err
			}

		case "scroll_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ScrollTotal = value
			case float64:
				f := int64(v)
				s.ScrollTotal = f
			}

		case "suggest_current":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SuggestCurrent = value
			case float64:
				f := int64(v)
				s.SuggestCurrent = f
			}

		case "suggest_time":
			if err := dec.Decode(&s.SuggestTime); err != nil {
				return err
			}

		case "suggest_time_in_millis":
			if err := dec.Decode(&s.SuggestTimeInMillis); err != nil {
				return err
			}

		case "suggest_total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SuggestTotal = value
			case float64:
				f := int64(v)
				s.SuggestTotal = f
			}

		}
	}
	return nil
}

// NewSearchStats returns a SearchStats.
func NewSearchStats() *SearchStats {
	r := &SearchStats{
		Groups: make(map[string]SearchStats, 0),
	}

	return r
}
