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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// MgetOperation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/mget/types.ts#L32-L55
type MgetOperation struct {
	// Id_ The unique document ID.
	Id_ string `json:"_id"`
	// Index_ The index that contains the document.
	Index_ *string `json:"_index,omitempty"`
	// Routing The key for the primary shard the document resides on. Required if routing is
	// used during indexing.
	Routing *string `json:"routing,omitempty"`
	// Source_ If `false`, excludes all _source fields.
	Source_ SourceConfig `json:"_source,omitempty"`
	// StoredFields The stored fields you want to retrieve.
	StoredFields []string                 `json:"stored_fields,omitempty"`
	Version      *int64                   `json:"version,omitempty"`
	VersionType  *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *MgetOperation) UnmarshalJSON(data []byte) error {

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

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "_source":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		source__field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Source_", err)
				}

				switch t {

				case "exclude_vectors", "excludes", "includes":
					o := NewSourceFilter()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Source_", err)
					}
					s.Source_ = o
					break source__field

				}
			}
			if s.Source_ == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Source_); err != nil {
					return fmt.Errorf("%s | %w", "Source_", err)
				}
			}

		case "stored_fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "StoredFields", err)
				}

				s.StoredFields = append(s.StoredFields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.StoredFields); err != nil {
					return fmt.Errorf("%s | %w", "StoredFields", err)
				}
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return fmt.Errorf("%s | %w", "VersionType", err)
			}

		}
	}
	return nil
}

// NewMgetOperation returns a MgetOperation.
func NewMgetOperation() *MgetOperation {
	r := &MgetOperation{}

	return r
}
