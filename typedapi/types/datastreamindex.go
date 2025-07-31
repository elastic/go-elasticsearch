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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/managedby"
)

// DataStreamIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/_types/DataStream.ts#L136-L157
type DataStreamIndex struct {
	// IlmPolicy Name of the current ILM lifecycle policy configured for this backing index.
	IlmPolicy *string `json:"ilm_policy,omitempty"`
	// IndexName Name of the backing index.
	IndexName string `json:"index_name"`
	// IndexUuid Universally unique identifier (UUID) for the index.
	IndexUuid string `json:"index_uuid"`
	// ManagedBy Name of the lifecycle system that's currently managing this backing index.
	ManagedBy *managedby.ManagedBy `json:"managed_by,omitempty"`
	// PreferIlm Indicates if ILM should take precedence over DSL in case both are configured
	// to manage this index.
	PreferIlm *bool `json:"prefer_ilm,omitempty"`
}

func (s *DataStreamIndex) UnmarshalJSON(data []byte) error {

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

		case "ilm_policy":
			if err := dec.Decode(&s.IlmPolicy); err != nil {
				return fmt.Errorf("%s | %w", "IlmPolicy", err)
			}

		case "index_name":
			if err := dec.Decode(&s.IndexName); err != nil {
				return fmt.Errorf("%s | %w", "IndexName", err)
			}

		case "index_uuid":
			if err := dec.Decode(&s.IndexUuid); err != nil {
				return fmt.Errorf("%s | %w", "IndexUuid", err)
			}

		case "managed_by":
			if err := dec.Decode(&s.ManagedBy); err != nil {
				return fmt.Errorf("%s | %w", "ManagedBy", err)
			}

		case "prefer_ilm":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PreferIlm", err)
				}
				s.PreferIlm = &value
			case bool:
				s.PreferIlm = &v
			}

		}
	}
	return nil
}

// NewDataStreamIndex returns a DataStreamIndex.
func NewDataStreamIndex() *DataStreamIndex {
	r := &DataStreamIndex{}

	return r
}
