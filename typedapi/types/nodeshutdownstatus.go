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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shutdownstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shutdowntype"
)

// NodeShutdownStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/shutdown/get_node/ShutdownGetNodeResponse.ts#L29-L38
type NodeShutdownStatus struct {
	NodeId                string                        `json:"node_id"`
	PersistentTasks       PersistentTaskStatus          `json:"persistent_tasks"`
	Plugins               PluginsStatus                 `json:"plugins"`
	Reason                string                        `json:"reason"`
	ShardMigration        ShardMigrationStatus          `json:"shard_migration"`
	ShutdownStartedmillis int64                         `json:"shutdown_startedmillis"`
	Status                shutdownstatus.ShutdownStatus `json:"status"`
	Type                  shutdowntype.ShutdownType     `json:"type"`
}

func (s *NodeShutdownStatus) UnmarshalJSON(data []byte) error {

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

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return fmt.Errorf("%s | %w", "NodeId", err)
			}

		case "persistent_tasks":
			if err := dec.Decode(&s.PersistentTasks); err != nil {
				return fmt.Errorf("%s | %w", "PersistentTasks", err)
			}

		case "plugins":
			if err := dec.Decode(&s.Plugins); err != nil {
				return fmt.Errorf("%s | %w", "Plugins", err)
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = o

		case "shard_migration":
			if err := dec.Decode(&s.ShardMigration); err != nil {
				return fmt.Errorf("%s | %w", "ShardMigration", err)
			}

		case "shutdown_startedmillis":
			if err := dec.Decode(&s.ShutdownStartedmillis); err != nil {
				return fmt.Errorf("%s | %w", "ShutdownStartedmillis", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// NewNodeShutdownStatus returns a NodeShutdownStatus.
func NewNodeShutdownStatus() *NodeShutdownStatus {
	r := &NodeShutdownStatus{}

	return r
}
