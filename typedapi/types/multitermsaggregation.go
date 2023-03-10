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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// MultiTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/bucket.ts#L265-L274
type MultiTermsAggregation struct {
	CollectMode           *termsaggregationcollectmode.TermsAggregationCollectMode `json:"collect_mode,omitempty"`
	Meta                  map[string]json.RawMessage                               `json:"meta,omitempty"`
	MinDocCount           *int64                                                   `json:"min_doc_count,omitempty"`
	Name                  *string                                                  `json:"name,omitempty"`
	Order                 AggregateOrder                                           `json:"order,omitempty"`
	ShardMinDocCount      *int64                                                   `json:"shard_min_doc_count,omitempty"`
	ShardSize             *int                                                     `json:"shard_size,omitempty"`
	ShowTermDocCountError *bool                                                    `json:"show_term_doc_count_error,omitempty"`
	Size                  *int                                                     `json:"size,omitempty"`
	Terms                 []MultiTermLookup                                        `json:"terms"`
}

func (s *MultiTermsAggregation) UnmarshalJSON(data []byte) error {
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

		case "collect_mode":
			if err := dec.Decode(&s.CollectMode); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "min_doc_count":
			if err := dec.Decode(&s.MinDocCount); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "order":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {

			case '{':
				o := make(map[string]sortorder.SortOrder, 0)
				localDec.Decode(&o)
				s.Order = o

			case '[':
				o := make([]map[string]sortorder.SortOrder, 0)
				localDec.Decode(&o)
				s.Order = o
			}

		case "shard_min_doc_count":
			if err := dec.Decode(&s.ShardMinDocCount); err != nil {
				return err
			}

		case "shard_size":
			if err := dec.Decode(&s.ShardSize); err != nil {
				return err
			}

		case "show_term_doc_count_error":
			if err := dec.Decode(&s.ShowTermDocCountError); err != nil {
				return err
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return err
			}

		case "terms":
			if err := dec.Decode(&s.Terms); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMultiTermsAggregation returns a MultiTermsAggregation.
func NewMultiTermsAggregation() *MultiTermsAggregation {
	r := &MultiTermsAggregation{}

	return r
}
