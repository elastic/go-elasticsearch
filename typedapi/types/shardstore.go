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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shardstoreallocation"
)

// ShardStore type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/shard_stores/types.ts#L30-L34
type ShardStore struct {
	Allocation     shardstoreallocation.ShardStoreAllocation `json:"allocation"`
	AllocationId   *string                                   `json:"allocation_id,omitempty"`
	ShardStore     map[string]ShardStoreNode                 `json:"ShardStore,omitempty"`
	StoreException *ShardStoreException                      `json:"store_exception,omitempty"`
}

func (s *ShardStore) UnmarshalJSON(data []byte) error {

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

		case "allocation":
			if err := dec.Decode(&s.Allocation); err != nil {
				return err
			}

		case "allocation_id":
			if err := dec.Decode(&s.AllocationId); err != nil {
				return err
			}

		case "ShardStore":
			if s.ShardStore == nil {
				s.ShardStore = make(map[string]ShardStoreNode, 0)
			}
			if err := dec.Decode(&s.ShardStore); err != nil {
				return err
			}

		case "store_exception":
			if err := dec.Decode(&s.StoreException); err != nil {
				return err
			}

		default:

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ShardStore) MarshalJSON() ([]byte, error) {
	type opt ShardStore
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.ShardStore {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "ShardStore")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewShardStore returns a ShardStore.
func NewShardStore() *ShardStore {
	r := &ShardStore{
		ShardStore: make(map[string]ShardStoreNode, 0),
	}

	return r
}
