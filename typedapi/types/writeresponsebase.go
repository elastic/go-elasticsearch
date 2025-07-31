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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)

// WriteResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Base.ts#L36-L67
type WriteResponseBase struct {
	ForcedRefresh *bool `json:"forced_refresh,omitempty"`
	// Id_ The unique identifier for the added document.
	Id_ string `json:"_id"`
	// Index_ The name of the index the document was added to.
	Index_ string `json:"_index"`
	// PrimaryTerm_ The primary term assigned to the document for the indexing operation.
	PrimaryTerm_ *int64 `json:"_primary_term,omitempty"`
	// Result The result of the indexing operation: `created` or `updated`.
	Result result.Result `json:"result"`
	// SeqNo_ The sequence number assigned to the document for the indexing operation.
	// Sequence numbers are used to ensure an older version of a document doesn't
	// overwrite a newer version.
	SeqNo_ *int64 `json:"_seq_no,omitempty"`
	// Shards_ Information about the replication process of the operation.
	Shards_ ShardStatistics `json:"_shards"`
	// Version_ The document version, which is incremented each time the document is updated.
	Version_ int64 `json:"_version"`
}

func (s *WriteResponseBase) UnmarshalJSON(data []byte) error {

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

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

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
			if err := dec.Decode(&s.Result); err != nil {
				return fmt.Errorf("%s | %w", "Result", err)
			}

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return fmt.Errorf("%s | %w", "SeqNo_", err)
			}

		case "_shards":
			if err := dec.Decode(&s.Shards_); err != nil {
				return fmt.Errorf("%s | %w", "Shards_", err)
			}

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return fmt.Errorf("%s | %w", "Version_", err)
			}

		}
	}
	return nil
}

// NewWriteResponseBase returns a WriteResponseBase.
func NewWriteResponseBase() *WriteResponseBase {
	r := &WriteResponseBase{}

	return r
}
