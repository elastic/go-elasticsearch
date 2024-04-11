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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DataStreamLifecycleRolloverConditions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/indices/_types/DataStreamLifecycle.ts#L57-L69
type DataStreamLifecycleRolloverConditions struct {
	MaxAge              *string  `json:"max_age,omitempty"`
	MaxDocs             *int64   `json:"max_docs,omitempty"`
	MaxPrimaryShardDocs *int64   `json:"max_primary_shard_docs,omitempty"`
	MaxPrimaryShardSize ByteSize `json:"max_primary_shard_size,omitempty"`
	MaxSize             ByteSize `json:"max_size,omitempty"`
	MinAge              Duration `json:"min_age,omitempty"`
	MinDocs             *int64   `json:"min_docs,omitempty"`
	MinPrimaryShardDocs *int64   `json:"min_primary_shard_docs,omitempty"`
	MinPrimaryShardSize ByteSize `json:"min_primary_shard_size,omitempty"`
	MinSize             ByteSize `json:"min_size,omitempty"`
}

func (s *DataStreamLifecycleRolloverConditions) UnmarshalJSON(data []byte) error {

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

		case "max_age":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxAge", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxAge = &o

		case "max_docs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxDocs", err)
				}
				s.MaxDocs = &value
			case float64:
				f := int64(v)
				s.MaxDocs = &f
			}

		case "max_primary_shard_docs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxPrimaryShardDocs", err)
				}
				s.MaxPrimaryShardDocs = &value
			case float64:
				f := int64(v)
				s.MaxPrimaryShardDocs = &f
			}

		case "max_primary_shard_size":
			if err := dec.Decode(&s.MaxPrimaryShardSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxPrimaryShardSize", err)
			}

		case "max_size":
			if err := dec.Decode(&s.MaxSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxSize", err)
			}

		case "min_age":
			if err := dec.Decode(&s.MinAge); err != nil {
				return fmt.Errorf("%s | %w", "MinAge", err)
			}

		case "min_docs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocs", err)
				}
				s.MinDocs = &value
			case float64:
				f := int64(v)
				s.MinDocs = &f
			}

		case "min_primary_shard_docs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinPrimaryShardDocs", err)
				}
				s.MinPrimaryShardDocs = &value
			case float64:
				f := int64(v)
				s.MinPrimaryShardDocs = &f
			}

		case "min_primary_shard_size":
			if err := dec.Decode(&s.MinPrimaryShardSize); err != nil {
				return fmt.Errorf("%s | %w", "MinPrimaryShardSize", err)
			}

		case "min_size":
			if err := dec.Decode(&s.MinSize); err != nil {
				return fmt.Errorf("%s | %w", "MinSize", err)
			}

		}
	}
	return nil
}

// NewDataStreamLifecycleRolloverConditions returns a DataStreamLifecycleRolloverConditions.
func NewDataStreamLifecycleRolloverConditions() *DataStreamLifecycleRolloverConditions {
	r := &DataStreamLifecycleRolloverConditions{}

	return r
}
