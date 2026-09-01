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
// https://github.com/elastic/elasticsearch-specification/tree/9665eef0d78c41f20c4c83e69b7c8155efd58f24

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ProjectEncryptionKeyDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9665eef0d78c41f20c4c83e69b7c8155efd58f24/specification/_global/health_report/types.ts#L226-L237
type ProjectEncryptionKeyDetails struct {
	ActiveKeyId      *string `json:"active_key_id,omitempty"`
	ActivePasswordId string  `json:"active_password_id"`
	// EncryptionRequired Whether callers must refuse to store secrets when the service is not ready.
	// If `false`, callers may fall back to storing secrets in plaintext (with a
	// warning).
	EncryptionRequired bool    `json:"encryption_required"`
	KeyCount           *int    `json:"key_count,omitempty"`
	MetadataPasswordId *string `json:"metadata_password_id,omitempty"`
	State              string  `json:"state"`
}

func (s *ProjectEncryptionKeyDetails) UnmarshalJSON(data []byte) error {

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

		case "active_key_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ActiveKeyId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ActiveKeyId = &o

		case "active_password_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ActivePasswordId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ActivePasswordId = o

		case "encryption_required":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "EncryptionRequired", err)
				}
				s.EncryptionRequired = value
			case bool:
				s.EncryptionRequired = v
			}

		case "key_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "KeyCount", err)
				}
				s.KeyCount = &value
			case float64:
				f := int(v)
				s.KeyCount = &f
			}

		case "metadata_password_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MetadataPasswordId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MetadataPasswordId = &o

		case "state":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "State", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.State = o

		}
	}
	return nil
}

// NewProjectEncryptionKeyDetails returns a ProjectEncryptionKeyDetails.
func NewProjectEncryptionKeyDetails() *ProjectEncryptionKeyDetails {
	r := &ProjectEncryptionKeyDetails{}

	return r
}
