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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ReplicationAccess type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/security/_types/Privileges.ts#L444-L454
type ReplicationAccess struct {
	// AllowRestrictedIndices This needs to be set to true if the patterns in the names field should cover
	// system indices.
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`
	// Names A list of indices (or index name patterns) to which the permissions in this
	// entry apply.
	Names []string `json:"names"`
}

func (s *ReplicationAccess) UnmarshalJSON(data []byte) error {

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

		case "allow_restricted_indices":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowRestrictedIndices", err)
				}
				s.AllowRestrictedIndices = &value
			case bool:
				s.AllowRestrictedIndices = &v
			}

		case "names":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Names", err)
				}

				s.Names = append(s.Names, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Names); err != nil {
					return fmt.Errorf("%s | %w", "Names", err)
				}
			}

		}
	}
	return nil
}

// NewReplicationAccess returns a ReplicationAccess.
func NewReplicationAccess() *ReplicationAccess {
	r := &ReplicationAccess{}

	return r
}

type ReplicationAccessVariant interface {
	ReplicationAccessCaster() *ReplicationAccess
}

func (s *ReplicationAccess) ReplicationAccessCaster() *ReplicationAccess {
	return s
}
