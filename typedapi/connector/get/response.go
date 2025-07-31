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

package get

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/connectorstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syncstatus"
)

// Response holds the response body struct for the package get
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/get/ConnectorGetResponse.ts#L22-L24
type Response struct {
	ApiKeyId                         *string                         `json:"api_key_id,omitempty"`
	ApiKeySecretId                   *string                         `json:"api_key_secret_id,omitempty"`
	Configuration                    types.ConnectorConfiguration    `json:"configuration"`
	CustomScheduling                 types.ConnectorCustomScheduling `json:"custom_scheduling"`
	Description                      *string                         `json:"description,omitempty"`
	Error                            *string                         `json:"error,omitempty"`
	Features                         *types.ConnectorFeatures        `json:"features,omitempty"`
	Filtering                        []types.FilteringConfig         `json:"filtering"`
	Id                               *string                         `json:"id,omitempty"`
	IndexName                        *string                         `json:"index_name,omitempty"`
	IsNative                         bool                            `json:"is_native"`
	Language                         *string                         `json:"language,omitempty"`
	LastAccessControlSyncError       *string                         `json:"last_access_control_sync_error,omitempty"`
	LastAccessControlSyncScheduledAt types.DateTime                  `json:"last_access_control_sync_scheduled_at,omitempty"`
	LastAccessControlSyncStatus      *syncstatus.SyncStatus          `json:"last_access_control_sync_status,omitempty"`
	LastDeletedDocumentCount         *int64                          `json:"last_deleted_document_count,omitempty"`
	LastIncrementalSyncScheduledAt   types.DateTime                  `json:"last_incremental_sync_scheduled_at,omitempty"`
	LastIndexedDocumentCount         *int64                          `json:"last_indexed_document_count,omitempty"`
	LastSeen                         types.DateTime                  `json:"last_seen,omitempty"`
	LastSyncError                    *string                         `json:"last_sync_error,omitempty"`
	LastSyncScheduledAt              types.DateTime                  `json:"last_sync_scheduled_at,omitempty"`
	LastSyncStatus                   *syncstatus.SyncStatus          `json:"last_sync_status,omitempty"`
	LastSynced                       types.DateTime                  `json:"last_synced,omitempty"`
	Name                             *string                         `json:"name,omitempty"`
	Pipeline                         *types.IngestPipelineParams     `json:"pipeline,omitempty"`
	Scheduling                       types.SchedulingConfiguration   `json:"scheduling"`
	ServiceType                      *string                         `json:"service_type,omitempty"`
	Status                           connectorstatus.ConnectorStatus `json:"status"`
	SyncCursor                       json.RawMessage                 `json:"sync_cursor,omitempty"`
	SyncNow                          bool                            `json:"sync_now"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "api_key_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ApiKeyId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ApiKeyId = &o

		case "api_key_secret_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ApiKeySecretId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ApiKeySecretId = &o

		case "configuration":
			if err := dec.Decode(&s.Configuration); err != nil {
				return fmt.Errorf("%s | %w", "Configuration", err)
			}

		case "custom_scheduling":
			if err := dec.Decode(&s.CustomScheduling); err != nil {
				return fmt.Errorf("%s | %w", "CustomScheduling", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "error":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Error", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Error = &o

		case "features":
			if err := dec.Decode(&s.Features); err != nil {
				return fmt.Errorf("%s | %w", "Features", err)
			}

		case "filtering":
			if err := dec.Decode(&s.Filtering); err != nil {
				return fmt.Errorf("%s | %w", "Filtering", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "index_name":
			if err := dec.Decode(&s.IndexName); err != nil {
				return fmt.Errorf("%s | %w", "IndexName", err)
			}

		case "is_native":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IsNative", err)
				}
				s.IsNative = value
			case bool:
				s.IsNative = v
			}

		case "language":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Language", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Language = &o

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

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "pipeline":
			if err := dec.Decode(&s.Pipeline); err != nil {
				return fmt.Errorf("%s | %w", "Pipeline", err)
			}

		case "scheduling":
			if err := dec.Decode(&s.Scheduling); err != nil {
				return fmt.Errorf("%s | %w", "Scheduling", err)
			}

		case "service_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ServiceType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ServiceType = &o

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "sync_cursor":
			if err := dec.Decode(&s.SyncCursor); err != nil {
				return fmt.Errorf("%s | %w", "SyncCursor", err)
			}

		case "sync_now":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SyncNow", err)
				}
				s.SyncNow = value
			case bool:
				s.SyncNow = v
			}

		}
	}
	return nil
}
