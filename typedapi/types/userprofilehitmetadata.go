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

// UserProfileHitMetadata type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/security/_types/UserProfile.ts#L27-L30
type UserProfileHitMetadata struct {
	PrimaryTerm_ int64 `json:"_primary_term"`
	SeqNo_       int64 `json:"_seq_no"`
}

func (s *UserProfileHitMetadata) UnmarshalJSON(data []byte) error {

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

		case "_primary_term":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryTerm_", err)
				}
				s.PrimaryTerm_ = value
			case float64:
				f := int64(v)
				s.PrimaryTerm_ = f
			}

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return fmt.Errorf("%s | %w", "SeqNo_", err)
			}

		}
	}
	return nil
}

// NewUserProfileHitMetadata returns a UserProfileHitMetadata.
func NewUserProfileHitMetadata() *UserProfileHitMetadata {
	r := &UserProfileHitMetadata{}

	return r
}
