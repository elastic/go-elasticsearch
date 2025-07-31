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

package syncjobupdatestats

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package syncjobupdatestats
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/sync_job_update_stats/SyncJobUpdateStatsRequest.ts#L24-L78
type Request struct {

	// DeletedDocumentCount The number of documents the sync job deleted.
	DeletedDocumentCount int64 `json:"deleted_document_count"`
	// IndexedDocumentCount The number of documents the sync job indexed.
	IndexedDocumentCount int64 `json:"indexed_document_count"`
	// IndexedDocumentVolume The total size of the data (in MiB) the sync job indexed.
	IndexedDocumentVolume int64 `json:"indexed_document_volume"`
	// LastSeen The timestamp to use in the `last_seen` property for the connector sync job.
	LastSeen types.Duration `json:"last_seen,omitempty"`
	// Metadata The connector-specific metadata.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// TotalDocumentCount The total number of documents in the target index after the sync job
	// finished.
	TotalDocumentCount *int `json:"total_document_count,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Syncjobupdatestats request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "deleted_document_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DeletedDocumentCount", err)
				}
				s.DeletedDocumentCount = value
			case float64:
				f := int64(v)
				s.DeletedDocumentCount = f
			}

		case "indexed_document_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexedDocumentCount", err)
				}
				s.IndexedDocumentCount = value
			case float64:
				f := int64(v)
				s.IndexedDocumentCount = f
			}

		case "indexed_document_volume":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexedDocumentVolume", err)
				}
				s.IndexedDocumentVolume = value
			case float64:
				f := int64(v)
				s.IndexedDocumentVolume = f
			}

		case "last_seen":
			if err := dec.Decode(&s.LastSeen); err != nil {
				return fmt.Errorf("%s | %w", "LastSeen", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "total_document_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalDocumentCount", err)
				}
				s.TotalDocumentCount = &value
			case float64:
				f := int(v)
				s.TotalDocumentCount = &f
			}

		}
	}
	return nil
}
