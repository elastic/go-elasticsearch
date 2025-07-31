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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/remoteclusterprivilege"
)

// RemoteClusterPrivileges type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/_types/Privileges.ts#L278-L290
type RemoteClusterPrivileges struct {
	// Clusters A list of cluster aliases to which the permissions in this entry apply.
	Clusters []string `json:"clusters"`
	// Privileges The cluster level privileges that owners of the role have on the remote
	// cluster.
	Privileges []remoteclusterprivilege.RemoteClusterPrivilege `json:"privileges"`
}

func (s *RemoteClusterPrivileges) UnmarshalJSON(data []byte) error {

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

		case "clusters":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Clusters", err)
				}

				s.Clusters = append(s.Clusters, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Clusters); err != nil {
					return fmt.Errorf("%s | %w", "Clusters", err)
				}
			}

		case "privileges":
			if err := dec.Decode(&s.Privileges); err != nil {
				return fmt.Errorf("%s | %w", "Privileges", err)
			}

		}
	}
	return nil
}

// NewRemoteClusterPrivileges returns a RemoteClusterPrivileges.
func NewRemoteClusterPrivileges() *RemoteClusterPrivileges {
	r := &RemoteClusterPrivileges{}

	return r
}
