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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syncjobtriggermethod"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syncjobtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syncstatus"
)

// ConnectorSyncJob type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/_types/SyncJob.ts#L53-L72
type ConnectorSyncJob struct {
	CancelationRequestedAt DateTime                                  `json:"cancelation_requested_at,omitempty"`
	CanceledAt             DateTime                                  `json:"canceled_at,omitempty"`
	CompletedAt            DateTime                                  `json:"completed_at,omitempty"`
	Connector              SyncJobConnectorReference                 `json:"connector"`
	CreatedAt              DateTime                                  `json:"created_at"`
	DeletedDocumentCount   int64                                     `json:"deleted_document_count"`
	Error                  *string                                   `json:"error,omitempty"`
	Id                     string                                    `json:"id"`
	IndexedDocumentCount   int64                                     `json:"indexed_document_count"`
	IndexedDocumentVolume  int64                                     `json:"indexed_document_volume"`
	JobType                syncjobtype.SyncJobType                   `json:"job_type"`
	LastSeen               DateTime                                  `json:"last_seen,omitempty"`
	Metadata               map[string]json.RawMessage                `json:"metadata"`
	StartedAt              DateTime                                  `json:"started_at,omitempty"`
	Status                 syncstatus.SyncStatus                     `json:"status"`
	TotalDocumentCount     int64                                     `json:"total_document_count"`
	TriggerMethod          syncjobtriggermethod.SyncJobTriggerMethod `json:"trigger_method"`
	WorkerHostname         *string                                   `json:"worker_hostname,omitempty"`
}

func (s *ConnectorSyncJob) UnmarshalJSON(data []byte) error {

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

		case "cancelation_requested_at":
			if err := dec.Decode(&s.CancelationRequestedAt); err != nil {
				return fmt.Errorf("%s | %w", "CancelationRequestedAt", err)
			}

		case "canceled_at":
			if err := dec.Decode(&s.CanceledAt); err != nil {
				return fmt.Errorf("%s | %w", "CanceledAt", err)
			}

		case "completed_at":
			if err := dec.Decode(&s.CompletedAt); err != nil {
				return fmt.Errorf("%s | %w", "CompletedAt", err)
			}

		case "connector":
			if err := dec.Decode(&s.Connector); err != nil {
				return fmt.Errorf("%s | %w", "Connector", err)
			}

		case "created_at":
			if err := dec.Decode(&s.CreatedAt); err != nil {
				return fmt.Errorf("%s | %w", "CreatedAt", err)
			}

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
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

		case "job_type":
			if err := dec.Decode(&s.JobType); err != nil {
				return fmt.Errorf("%s | %w", "JobType", err)
			}

		case "last_seen":
			if err := dec.Decode(&s.LastSeen); err != nil {
				return fmt.Errorf("%s | %w", "LastSeen", err)
			}

		case "metadata":
			if s.Metadata == nil {
				s.Metadata = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "started_at":
			if err := dec.Decode(&s.StartedAt); err != nil {
				return fmt.Errorf("%s | %w", "StartedAt", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "total_document_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalDocumentCount", err)
				}
				s.TotalDocumentCount = value
			case float64:
				f := int64(v)
				s.TotalDocumentCount = f
			}

		case "trigger_method":
			if err := dec.Decode(&s.TriggerMethod); err != nil {
				return fmt.Errorf("%s | %w", "TriggerMethod", err)
			}

		case "worker_hostname":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "WorkerHostname", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.WorkerHostname = &o

		}
	}
	return nil
}

// NewConnectorSyncJob returns a ConnectorSyncJob.
func NewConnectorSyncJob() *ConnectorSyncJob {
	r := &ConnectorSyncJob{
		Metadata: make(map[string]json.RawMessage),
	}

	return r
}
