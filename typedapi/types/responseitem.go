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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/failurestorestatus"
)

// ResponseItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/bulk/types.ts#L37-L84
type ResponseItem struct {
	// Error Additional information about the failed operation.
	// The property is returned only for failed operations.
	Error         *ErrorCause                            `json:"error,omitempty"`
	FailureStore  *failurestorestatus.FailureStoreStatus `json:"failure_store,omitempty"`
	ForcedRefresh *bool                                  `json:"forced_refresh,omitempty"`
	Get           *InlineGetDictUserDefined              `json:"get,omitempty"`
	// Id_ The document ID associated with the operation.
	Id_ *string `json:"_id,omitempty"`
	// Index_ The name of the index associated with the operation.
	// If the operation targeted a data stream, this is the backing index into which
	// the document was written.
	Index_ string `json:"_index"`
	// PrimaryTerm_ The primary term assigned to the document for the operation.
	// This property is returned only for successful operations.
	PrimaryTerm_ *int64 `json:"_primary_term,omitempty"`
	// Result The result of the operation.
	// Successful values are `created`, `deleted`, and `updated`.
	Result *string `json:"result,omitempty"`
	// SeqNo_ The sequence number assigned to the document for the operation.
	// Sequence numbers are used to ensure an older version of a document doesn't
	// overwrite a newer version.
	SeqNo_ *int64 `json:"_seq_no,omitempty"`
	// Shards_ Shard information for the operation.
	Shards_ *ShardStatistics `json:"_shards,omitempty"`
	// Status The HTTP status code returned for the operation.
	Status int `json:"status"`
	// Version_ The document version associated with the operation.
	// The document version is incremented each time the document is updated.
	// This property is returned only for successful actions.
	Version_ *int64 `json:"_version,omitempty"`
}

func (s *ResponseItem) UnmarshalJSON(data []byte) error {

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

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}

		case "failure_store":
			if err := dec.Decode(&s.FailureStore); err != nil {
				return fmt.Errorf("%s | %w", "FailureStore", err)
			}

		case "forced_refresh":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ForcedRefresh", err)
				}
				s.ForcedRefresh = &value
			case bool:
				s.ForcedRefresh = &v
			}

		case "get":
			if err := dec.Decode(&s.Get); err != nil {
				return fmt.Errorf("%s | %w", "Get", err)
			}

		case "_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id_ = &o

		case "_index":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Index_ = o

		case "_primary_term":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryTerm_", err)
				}
				s.PrimaryTerm_ = &value
			case float64:
				f := int64(v)
				s.PrimaryTerm_ = &f
			}

		case "result":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Result", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Result = &o

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return fmt.Errorf("%s | %w", "SeqNo_", err)
			}

		case "_shards":
			if err := dec.Decode(&s.Shards_); err != nil {
				return fmt.Errorf("%s | %w", "Shards_", err)
			}

		case "status":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Status", err)
				}
				s.Status = value
			case float64:
				f := int(v)
				s.Status = f
			}

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return fmt.Errorf("%s | %w", "Version_", err)
			}

		}
	}
	return nil
}

// NewResponseItem returns a ResponseItem.
func NewResponseItem() *ResponseItem {
	r := &ResponseItem{}

	return r
}
