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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/storagetype"
)

// Storage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L547-L558
type Storage struct {
	// AllowMmap You can restrict the use of the mmapfs and the related hybridfs store type
	// via the setting node.store.allow_mmap.
	// This is a boolean setting indicating whether or not memory-mapping is
	// allowed. The default is to allow it. This
	// setting is useful, for example, if you are in an environment where you can
	// not control the ability to create a lot
	// of memory maps so you need disable the ability to use memory-mapping.
	AllowMmap *bool `json:"allow_mmap,omitempty"`
	// StatsRefreshInterval How often store statistics are refreshed
	StatsRefreshInterval Duration                `json:"stats_refresh_interval,omitempty"`
	Type                 storagetype.StorageType `json:"type"`
}

func (s *Storage) UnmarshalJSON(data []byte) error {

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

		case "allow_mmap":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowMmap", err)
				}
				s.AllowMmap = &value
			case bool:
				s.AllowMmap = &v
			}

		case "stats_refresh_interval":
			if err := dec.Decode(&s.StatsRefreshInterval); err != nil {
				return fmt.Errorf("%s | %w", "StatsRefreshInterval", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// NewStorage returns a Storage.
func NewStorage() *Storage {
	r := &Storage{}

	return r
}

type StorageVariant interface {
	StorageCaster() *Storage
}

func (s *Storage) StorageCaster() *Storage {
	return s
}
