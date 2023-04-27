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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// CleanupRepositoryResults type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/snapshot/cleanup_repository/SnapshotCleanupRepositoryResponse.ts#L29-L34
type CleanupRepositoryResults struct {
	// DeletedBlobs Number of binary large objects (blobs) removed during cleanup.
	DeletedBlobs int64 `json:"deleted_blobs"`
	// DeletedBytes Number of bytes freed by cleanup operations.
	DeletedBytes int64 `json:"deleted_bytes"`
}

func (s *CleanupRepositoryResults) UnmarshalJSON(data []byte) error {

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

		case "deleted_blobs":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DeletedBlobs = value
			case float64:
				f := int64(v)
				s.DeletedBlobs = f
			}

		case "deleted_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DeletedBytes = value
			case float64:
				f := int64(v)
				s.DeletedBytes = f
			}

		}
	}
	return nil
}

// NewCleanupRepositoryResults returns a CleanupRepositoryResults.
func NewCleanupRepositoryResults() *CleanupRepositoryResults {
	r := &CleanupRepositoryResults{}

	return r
}
