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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// TermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/bucket.ts#L380-L397
type TermsAggregation struct {
	CollectMode           *termsaggregationcollectmode.TermsAggregationCollectMode     `json:"collect_mode,omitempty"`
	Exclude               []string                                                     `json:"exclude,omitempty"`
	ExecutionHint         *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field                 *string                                                      `json:"field,omitempty"`
	Format                *string                                                      `json:"format,omitempty"`
	Include               TermsInclude                                                 `json:"include,omitempty"`
	Meta                  map[string]json.RawMessage                                   `json:"meta,omitempty"`
	MinDocCount           *int                                                         `json:"min_doc_count,omitempty"`
	Missing               Missing                                                      `json:"missing,omitempty"`
	MissingBucket         *bool                                                        `json:"missing_bucket,omitempty"`
	MissingOrder          *missingorder.MissingOrder                                   `json:"missing_order,omitempty"`
	Name                  *string                                                      `json:"name,omitempty"`
	Order                 AggregateOrder                                               `json:"order,omitempty"`
	Script                Script                                                       `json:"script,omitempty"`
	ShardSize             *int                                                         `json:"shard_size,omitempty"`
	ShowTermDocCountError *bool                                                        `json:"show_term_doc_count_error,omitempty"`
	Size                  *int                                                         `json:"size,omitempty"`
	ValueType             *string                                                      `json:"value_type,omitempty"`
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
				return err
			}

		case "exclude":
			if err := dec.Decode(&s.Exclude); err != nil {
				return err
			}

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return err
			}

		case "include":
			if err := dec.Decode(&s.Include); err != nil {
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

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return err
			}

		case "missing_bucket":
			if err := dec.Decode(&s.MissingBucket); err != nil {
				return err
			}

		case "missing_order":
			if err := dec.Decode(&s.MissingOrder); err != nil {
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

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
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

		case "value_type":
			if err := dec.Decode(&s.ValueType); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTermsAggregation returns a TermsAggregation.
func NewTermsAggregation() *TermsAggregation {
	r := &TermsAggregation{}

	return r
}
