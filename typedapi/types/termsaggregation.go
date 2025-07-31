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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

// TermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L963-L1031
type TermsAggregation struct {
	// CollectMode Determines how child aggregations should be calculated: breadth-first or
	// depth-first.
	CollectMode *termsaggregationcollectmode.TermsAggregationCollectMode `json:"collect_mode,omitempty"`
	// Exclude Values to exclude.
	// Accepts regular expressions and partitions.
	Exclude []string `json:"exclude,omitempty"`
	// ExecutionHint Determines whether the aggregation will use field values directly or global
	// ordinals.
	ExecutionHint *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	// Field The field from which to return terms.
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`
	// Include Values to include.
	// Accepts regular expressions and partitions.
	Include TermsInclude `json:"include,omitempty"`
	// MinDocCount Only return values that are found in more than `min_doc_count` hits.
	MinDocCount *int `json:"min_doc_count,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing       Missing                    `json:"missing,omitempty"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`
	// Order Specifies the sort order of the buckets.
	// Defaults to sorting by descending document count.
	Order  AggregateOrder `json:"order,omitempty"`
	Script *Script        `json:"script,omitempty"`
	// ShardMinDocCount Regulates the certainty a shard has if the term should actually be added to
	// the candidate list or not with respect to the `min_doc_count`.
	// Terms will only be considered if their local shard frequency within the set
	// is higher than the `shard_min_doc_count`.
	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`
	// ShardSize The number of candidate terms produced by each shard.
	// By default, `shard_size` will be automatically estimated based on the number
	// of shards and the `size` parameter.
	ShardSize *int `json:"shard_size,omitempty"`
	// ShowTermDocCountError Set to `true` to return the `doc_count_error_upper_bound`, which is an upper
	// bound to the error on the `doc_count` returned by each shard.
	ShowTermDocCountError *bool `json:"show_term_doc_count_error,omitempty"`
	// Size The number of buckets returned out of the overall terms list.
	Size *int `json:"size,omitempty"`
	// ValueType Coerced unmapped fields into the specified type.
	ValueType *string `json:"value_type,omitempty"`
}

func (s *TermsAggregation) UnmarshalJSON(data []byte) error {

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

		case "exclude":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Exclude", err)
				}

				s.Exclude = append(s.Exclude, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Exclude); err != nil {
					return fmt.Errorf("%s | %w", "Exclude", err)
				}
			}

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionHint", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "include":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Include", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		include_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Include", err)
				}

				switch t {

				case "num_partitions", "partition":
					o := NewTermsPartition()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Include", err)
					}
					s.Include = o
					break include_field

				}
			}
			if s.Include == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Include); err != nil {
					return fmt.Errorf("%s | %w", "Include", err)
				}
			}

		case "min_doc_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocCount", err)
				}
				s.MinDocCount = &value
			case float64:
				f := int(v)
				s.MinDocCount = &f
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "missing_bucket":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MissingBucket", err)
				}
				s.MissingBucket = &value
			case bool:
				s.MissingBucket = &v
			}

		case "missing_order":
			if err := dec.Decode(&s.MissingOrder); err != nil {
				return fmt.Errorf("%s | %w", "MissingOrder", err)
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

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
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

		case "value_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ValueType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ValueType = &o

		}
	}
	return nil
}

// NewTermsAggregation returns a TermsAggregation.
func NewTermsAggregation() *TermsAggregation {
	r := &TermsAggregation{}

	return r
}
