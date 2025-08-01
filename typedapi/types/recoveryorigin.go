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

// RecoveryOrigin type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/recovery/types.ts#L76-L89
type RecoveryOrigin struct {
	BootstrapNewHistoryUuid *bool   `json:"bootstrap_new_history_uuid,omitempty"`
	Host                    *string `json:"host,omitempty"`
	Hostname                *string `json:"hostname,omitempty"`
	Id                      *string `json:"id,omitempty"`
	Index                   *string `json:"index,omitempty"`
	Ip                      *string `json:"ip,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Repository              *string `json:"repository,omitempty"`
	RestoreUUID             *string `json:"restoreUUID,omitempty"`
	Snapshot                *string `json:"snapshot,omitempty"`
	TransportAddress        *string `json:"transport_address,omitempty"`
	Version                 *string `json:"version,omitempty"`
}

func (s *RecoveryOrigin) UnmarshalJSON(data []byte) error {

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

		case "bootstrap_new_history_uuid":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BootstrapNewHistoryUuid", err)
				}
				s.BootstrapNewHistoryUuid = &value
			case bool:
				s.BootstrapNewHistoryUuid = &v
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return fmt.Errorf("%s | %w", "Host", err)
			}

		case "hostname":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Hostname", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Hostname = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "ip":
			if err := dec.Decode(&s.Ip); err != nil {
				return fmt.Errorf("%s | %w", "Ip", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "repository":
			if err := dec.Decode(&s.Repository); err != nil {
				return fmt.Errorf("%s | %w", "Repository", err)
			}

		case "restoreUUID":
			if err := dec.Decode(&s.RestoreUUID); err != nil {
				return fmt.Errorf("%s | %w", "RestoreUUID", err)
			}

		case "snapshot":
			if err := dec.Decode(&s.Snapshot); err != nil {
				return fmt.Errorf("%s | %w", "Snapshot", err)
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return fmt.Errorf("%s | %w", "TransportAddress", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewRecoveryOrigin returns a RecoveryOrigin.
func NewRecoveryOrigin() *RecoveryOrigin {
	r := &RecoveryOrigin{}

	return r
}
