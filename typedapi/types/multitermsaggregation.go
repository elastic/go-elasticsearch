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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
)

// MultiTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L601-L641
type MultiTermsAggregation struct {
	// CollectMode Specifies the strategy for data collection.
	CollectMode *termsaggregationcollectmode.TermsAggregationCollectMode `json:"collect_mode,omitempty"`
	// MinDocCount The minimum number of documents in a bucket for it to be returned.
	MinDocCount *int64 `json:"min_doc_count,omitempty"`
	// Order Specifies the sort order of the buckets.
	// Defaults to sorting by descending document count.
	Order AggregateOrder `json:"order,omitempty"`
	// ShardMinDocCount The minimum number of documents in a bucket on each shard for it to be
	// returned.
	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`
	// ShardSize The number of candidate terms produced by each shard.
	// By default, `shard_size` will be automatically estimated based on the number
	// of shards and the `size` parameter.
	ShardSize *int `json:"shard_size,omitempty"`
	// ShowTermDocCountError Calculates the doc count error on per term basis.
	ShowTermDocCountError *bool `json:"show_term_doc_count_error,omitempty"`
	// Size The number of term buckets should be returned out of the overall terms list.
	Size *int `json:"size,omitempty"`
	// Terms The field from which to generate sets of terms.
	Terms []MultiTermLookup `json:"terms"`
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
				return fmt.Errorf("%s | %w", "CollectMode", err)
			}

		case "min_doc_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocCount", err)
				}
				s.MinDocCount = &value
			case float64:
				f := int64(v)
				s.MinDocCount = &f
			}

		case "order":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]sortorder.SortOrder, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}
				s.Order = o
			case '[':
				o := make([]map[string]sortorder.SortOrder, 0)
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}
				s.Order = o
			}

		case "shard_min_doc_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardMinDocCount", err)
				}
				s.ShardMinDocCount = &value
			case float64:
				f := int64(v)
				s.ShardMinDocCount = &f
			}

		case "shard_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSize", err)
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
			}

		case "show_term_doc_count_error":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShowTermDocCountError", err)
				}
				s.ShowTermDocCountError = &value
			case bool:
				s.ShowTermDocCountError = &v
			}

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "terms":
			if err := dec.Decode(&s.Terms); err != nil {
				return fmt.Errorf("%s | %w", "Terms", err)
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
