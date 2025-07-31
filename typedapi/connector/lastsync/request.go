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

package lastsync

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syncstatus"
)

// Request holds the request body struct for the package lastsync
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/last_sync/ConnectorUpdateLastSyncRequest.ts#L26-L66
type Request struct {
	LastAccessControlSyncError       *string                `json:"last_access_control_sync_error,omitempty"`
	LastAccessControlSyncScheduledAt types.DateTime         `json:"last_access_control_sync_scheduled_at,omitempty"`
	LastAccessControlSyncStatus      *syncstatus.SyncStatus `json:"last_access_control_sync_status,omitempty"`
	LastDeletedDocumentCount         *int64                 `json:"last_deleted_document_count,omitempty"`
	LastIncrementalSyncScheduledAt   types.DateTime         `json:"last_incremental_sync_scheduled_at,omitempty"`
	LastIndexedDocumentCount         *int64                 `json:"last_indexed_document_count,omitempty"`
	LastSeen                         types.DateTime         `json:"last_seen,omitempty"`
	LastSyncError                    *string                `json:"last_sync_error,omitempty"`
	LastSyncScheduledAt              types.DateTime         `json:"last_sync_scheduled_at,omitempty"`
	LastSyncStatus                   *syncstatus.SyncStatus `json:"last_sync_status,omitempty"`
	LastSynced                       types.DateTime         `json:"last_synced,omitempty"`
	SyncCursor                       json.RawMessage        `json:"sync_cursor,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Lastsync request: %w", err)
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

		case "last_access_control_sync_error":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "LastAccessControlSyncError", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LastAccessControlSyncError = &o

		case "last_access_control_sync_scheduled_at":
			if err := dec.Decode(&s.LastAccessControlSyncScheduledAt); err != nil {
				return fmt.Errorf("%s | %w", "LastAccessControlSyncScheduledAt", err)
			}

		case "last_access_control_sync_status":
			if err := dec.Decode(&s.LastAccessControlSyncStatus); err != nil {
				return fmt.Errorf("%s | %w", "LastAccessControlSyncStatus", err)
			}

		case "last_deleted_document_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LastDeletedDocumentCount", err)
				}
				s.LastDeletedDocumentCount = &value
			case float64:
				f := int64(v)
				s.LastDeletedDocumentCount = &f
			}

		case "last_incremental_sync_scheduled_at":
			if err := dec.Decode(&s.LastIncrementalSyncScheduledAt); err != nil {
				return fmt.Errorf("%s | %w", "LastIncrementalSyncScheduledAt", err)
			}

		case "last_indexed_document_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LastIndexedDocumentCount", err)
				}
				s.LastIndexedDocumentCount = &value
			case float64:
				f := int64(v)
				s.LastIndexedDocumentCount = &f
			}

		case "last_seen":
			if err := dec.Decode(&s.LastSeen); err != nil {
				return fmt.Errorf("%s | %w", "LastSeen", err)
			}

		case "last_sync_error":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "LastSyncError", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LastSyncError = &o

		case "last_sync_scheduled_at":
			if err := dec.Decode(&s.LastSyncScheduledAt); err != nil {
				return fmt.Errorf("%s | %w", "LastSyncScheduledAt", err)
			}

		case "last_sync_status":
			if err := dec.Decode(&s.LastSyncStatus); err != nil {
				return fmt.Errorf("%s | %w", "LastSyncStatus", err)
			}

		case "last_synced":
			if err := dec.Decode(&s.LastSynced); err != nil {
				return fmt.Errorf("%s | %w", "LastSynced", err)
			}

		case "sync_cursor":
			if err := dec.Decode(&s.SyncCursor); err != nil {
				return fmt.Errorf("%s | %w", "SyncCursor", err)
			}

		}
	}
	return nil
}
