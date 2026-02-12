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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// NodeReloadResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/nodes/_types/NodeReloadResult.ts#L24-L43
type NodeReloadResult struct {
	// KeystoreDigest A SHA-256 hash of the keystore file contents.
	KeystoreDigest *string `json:"keystore_digest,omitempty"`
	// KeystoreLastModifiedTime The last modification time of the keystore file.
	KeystoreLastModifiedTime DateTime `json:"keystore_last_modified_time,omitempty"`
	// KeystorePath The path to the keystore file.
	KeystorePath    *string     `json:"keystore_path,omitempty"`
	Name            string      `json:"name"`
	ReloadException *ErrorCause `json:"reload_exception,omitempty"`
	// SecureSettingNames The names of the secure settings that were reloaded.
	SecureSettingNames []string `json:"secure_setting_names,omitempty"`
}

func (s *NodeReloadResult) UnmarshalJSON(data []byte) error {

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

		case "keystore_digest":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "KeystoreDigest", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.KeystoreDigest = &o

		case "keystore_last_modified_time":
			if err := dec.Decode(&s.KeystoreLastModifiedTime); err != nil {
				return fmt.Errorf("%s | %w", "KeystoreLastModifiedTime", err)
			}

		case "keystore_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "KeystorePath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.KeystorePath = &o

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "reload_exception":
			if err := dec.Decode(&s.ReloadException); err != nil {
				return fmt.Errorf("%s | %w", "ReloadException", err)
			}

		case "secure_setting_names":
			if err := dec.Decode(&s.SecureSettingNames); err != nil {
				return fmt.Errorf("%s | %w", "SecureSettingNames", err)
			}

		}
	}
	return nil
}

// NewNodeReloadResult returns a NodeReloadResult.
func NewNodeReloadResult() *NodeReloadResult {
	r := &NodeReloadResult{}

	return r
}
